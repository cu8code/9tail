Here's a brief explanation of the directory structure:

- cmd/server: Contains the main application entry point.
internal: Houses the core logic of the system, divided into workflow, block, and scheduler packages.
- pkg: Contains packages that could potentially be used by external projects, like the NATS client wrapper.
- api: For API-related code, such as HTTP handlers.
- config: For configuration-related code.
- scripts: For any build or deployment scripts.

// File: cmd/server/main.go

package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/yourusername/nats-workflow/internal/block"
	"github.com/yourusername/nats-workflow/internal/scheduler"
	"github.com/yourusername/nats-workflow/internal/workflow"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	wm := workflow.NewManager(nc)
	s := scheduler.NewScheduler(nc)
	be := block.NewExecutor(nc)

	// Start the components
	go wm.Start()
	go s.Start()
	go be.Start()

	log.Println("NATS-based Workflow System started")

	// Keep the main goroutine alive
	select {}
}

// File: internal/workflow/manager.go

package workflow

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Manager struct {
	nc *nats.Conn
}

func NewManager(nc *nats.Conn) *Manager {
	return &Manager{nc: nc}
}

func (m *Manager) Start() {
	sub, err := m.nc.Subscribe("workflow.*.start", m.handleWorkflowStart)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	select {} // Keep the goroutine alive
}

func (m *Manager) handleWorkflowStart(msg *nats.Msg) {
	var config map[string]interface{}
	err := json.Unmarshal(msg.Data, &config)
	if err != nil {
		log.Printf("Error unmarshaling workflow config: %v", err)
		return
	}

	workflowID := getWorkflowIDFromSubject(msg.Subject)
	wcb := m.createWCB(workflowID, config)

	wcbJSON, err := json.Marshal(wcb)
	if err != nil {
		log.Printf("Error marshaling WCB: %v", err)
		return
	}

	err = m.nc.Publish("workflow."+workflowID+".initiated", wcbJSON)
	if err != nil {
		log.Printf("Error publishing workflow initiation: %v", err)
	}
}

func (m *Manager) createWCB(workflowID string, config map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id":     workflowID,
		"blocks": config["blocks"],
		// Add other necessary fields
	}
}

func getWorkflowIDFromSubject(subject string) string {
	// Extract workflow ID from subject (e.g., "workflow.123.start" -> "123")
	// Implement this function based on your subject naming convention
	return "123" // Placeholder
}

// File: internal/scheduler/scheduler.go

package scheduler

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Scheduler struct {
	nc *nats.Conn
}

func NewScheduler(nc *nats.Conn) *Scheduler {
	return &Scheduler{nc: nc}
}

func (s *Scheduler) Start() {
	sub, err := s.nc.Subscribe("workflow.*.initiated", s.handleWorkflowInitiation)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	select {} // Keep the goroutine alive
}

func (s *Scheduler) handleWorkflowInitiation(msg *nats.Msg) {
	var wcb map[string]interface{}
	err := json.Unmarshal(msg.Data, &wcb)
	if err != nil {
		log.Printf("Error unmarshaling WCB: %v", err)
		return
	}

	blocks, ok := wcb["blocks"].([]interface{})
	if !ok {
		log.Printf("Invalid blocks data in WCB")
		return
	}

	for _, block := range blocks {
		blockMap, ok := block.(map[string]interface{})
		if !ok {
			continue
		}
		if dependencies, ok := blockMap["dependencies"].([]interface{}); !ok || len(dependencies) == 0 {
			blockJSON, _ := json.Marshal(blockMap)
			err = s.nc.Publish("block."+blockMap["id"].(string)+".execute", blockJSON)
			if err != nil {
				log.Printf("Error publishing block execution: %v", err)
			}
		}
	}
}

// File: internal/block/executor.go

package block

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Executor struct {
	nc *nats.Conn
}

func NewExecutor(nc *nats.Conn) *Executor {
	return &Executor{nc: nc}
}

func (e *Executor) Start() {
	sub, err := e.nc.Subscribe("block.*.execute", e.handleBlockExecution)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	select {} // Keep the goroutine alive
}

func (e *Executor) handleBlockExecution(msg *nats.Msg) {
	var blockConfig map[string]interface{}
	err := json.Unmarshal(msg.Data, &blockConfig)
	if err != nil {
		log.Printf("Error unmarshaling block config: %v", err)
		return
	}

	block := e.createBlock(blockConfig)
	result, err := block.Execute(blockConfig["input"])
	if err != nil {
		errJSON, _ := json.Marshal(err.Error())
		e.nc.Publish("block."+block.GetID()+".error", errJSON)
		return
	}

	resultJSON, _ := json.Marshal(result)
	e.nc.Publish("block."+block.GetID()+".output", resultJSON)
}

func (e *Executor) createBlock(config map[string]interface{}) Block {
	blockType, _ := config["type"].(string)
	blockID, _ := config["id"].(string)

	switch blockType {
	case "fetch":
		return NewFetchBlock(blockID, e.nc)
	case "html-scraper":
		return NewHtmlScraperBlock(blockID, e.nc)
	case "hook":
		return NewHookBlock(blockID, e.nc)
	case "response":
		return NewResponseBlock(blockID, e.nc)
	default:
		log.Printf("Unknown block type: %s", blockType)
		return nil
	}
}

// File: internal/block/block.go

package block

import (
	"github.com/nats-io/nats.go"
)

type Block interface {
	GetID() string
	GetType() string
	Execute(input interface{}) (interface{}, error)
}

type BaseBlock struct {
	ID         string
	Type       string
	NatsClient *nats.Conn
}

func (b *BaseBlock) GetID() string {
	return b.ID
}

func (b *BaseBlock) GetType() string {
	return b.Type
}

// File: internal/block/fetch_block.go

package block

import (
	"io/ioutil"
	"net/http"

	"github.com/nats-io/nats.go"
)

type FetchBlock struct {
	BaseBlock
}

func NewFetchBlock(id string, nc *nats.Conn) *FetchBlock {
	return &FetchBlock{
		BaseBlock: BaseBlock{
			ID:         id,
			Type:       "fetch",
			NatsClient: nc,
		},
	}
}

func (f *FetchBlock) Execute(input interface{}) (interface{}, error) {
	url, ok := input.(map[string]interface{})["url"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return string(body), nil
}

// Implement other block types (HtmlScraperBlock, HookBlock, ResponseBlock) similarly