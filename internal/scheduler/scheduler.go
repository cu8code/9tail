package scheduler

import (
	"encoding/json"
	"log"

	"github.com/cu8code/9tail/internal/workflow"
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

	workflowID, ok := wcb["id"].(string)
	if !ok {
		log.Printf("Invalid workflow ID in WCB")
		return
	}

	ctx := workflow.NewWorkflowContext(workflowID, s.nc)

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
			s.executeBlock(blockMap, ctx)
		}
	}
}

func (s *Scheduler) executeBlock(blockMap map[string]interface{}, ctx *workflow.WorkflowContext) {
	blockJSON, _ := json.Marshal(blockMap)
	err := s.nc.Publish("block."+blockMap["id"].(string)+".execute", blockJSON)
	if err != nil {
		log.Printf("Error publishing block execution: %v", err)
	}
}
