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

	workflowID, ok := blockConfig["workflow_id"].(string)
	if !ok {
		log.Printf("Missing workflow ID in block config")
		return
	}

	var js, jserr = e.nc.JetStream()
	if jserr != nil {
		log.Printf("Error connecting to JetStream: %v", jserr)
		return
	}
	var store, store_err = js.KeyValue("workflow_contexts")
	if store_err != nil {
		log.Printf("Error getting Key-Value store: %v", store_err)
		return
	}

	ctx := workflow.NewWorkflowContext(workflowID, e.nc, store)
	ero := e.loadWorkflowContext(ctx)
	if ero != nil {
		log.Printf("Error loading workflow context: %v", ero)
		return
	}

	block := e.createBlock(blockConfig)
	result, err := block.Execute(blockConfig["input"], ctx)
	if err != nil {
		errJSON, _ := json.Marshal(err.Error())
		e.nc.Publish("block."+block.GetID()+".error", errJSON)
		return
	}

	ctx.Set(block.GetID(), result)
	ctx.Publish()

	resultJSON, _ := json.Marshal(result)
	e.nc.Publish("block."+block.GetID()+".output", resultJSON)
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
