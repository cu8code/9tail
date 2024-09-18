package block

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

type Executor struct {
	nc       *nats.Conn
	registry *BlockRegistry
}

func NewExecutor(nc *nats.Conn, registry *BlockRegistry) *Executor {
	return &Executor{nc: nc, registry: registry}
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

	workflowID, ok := blockConfig["workflow_id"].(string)
	if !ok {
		log.Printf("Missing workflow ID in block config")
		return
	}

	blockType, ok := blockConfig["type"].(string)
	if !ok {
		log.Printf("Missing block type in block config")
		return
	}

	subject, ok := e.registry.blocks[blockType]
	if !ok {
		log.Printf("Unknown block type: %s", blockType)
		return
	}

	js, err := e.nc.JetStream()
	if err != nil {
		log.Printf("Error connecting to JetStream: %v", err)
		return
	}

	store, err := js.KeyValue("workflow_contexts")
	if err != nil {
		log.Printf("Error getting Key-Value store: %v", err)
		return
	}

	ctx := workflow.NewWorkflowContext(workflowID, e.nc, store)
	err = e.loadWorkflowContext(ctx)
	if err != nil {
		log.Printf("Error loading workflow context: %v", err)
		return
	}

	// Execute the block by sending a message to its execution subject
	response, err := e.nc.Request(subject, msg.Data, 30*time.Second)
	if err != nil {
		log.Printf("Error executing block: %v", err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Data, &result)
	if err != nil {
		log.Printf("Error unmarshaling block result: %v", err)
		return
	}

	blockID, _ := blockConfig["id"].(string)
	ctx.Set(blockID, result)
	ctx.Publish()

	resultJSON, _ := json.Marshal(result)
	e.nc.Publish(fmt.Sprintf("block.%s.output", blockID), resultJSON)
}

func (e *Executor) loadWorkflowContext(ctx *workflow.WorkflowContext) error {
	msg, err := e.nc.Request(fmt.Sprintf("workflow.%s.context", ctx.ID), nil, 5*time.Second)
	if err != nil {
		log.Printf("Error loading workflow context: %v", err)
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("Error unmarshaling workflow context: %v", err)
		return err
	}

	for k, v := range data {
		ctx.Set(k, v)
	}
	return nil
}
