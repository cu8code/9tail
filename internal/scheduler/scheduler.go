package scheduler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Scheduler struct {
	nc *nats.Conn
}

func NewScheduler(nc *nats.Conn) *Scheduler {
	return &Scheduler{nc: nc}
}

func (s *Scheduler) Start() {
	_, err := s.nc.Subscribe("workflow.*.initiated", s.handleWorkflowInitiation)
	if err != nil {
		log.Fatal(err)
	}
	// Keep the goroutine alive to listen for messages
	select {}
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

	// Iterate over blocks and execute them sequentially
	for _, block := range blocks {
		blockMap, ok := block.(map[string]interface{})
		if !ok {
			continue
		}

		// Execute block if there are no dependencies or if dependencies are already resolved
		if dependencies, ok := blockMap["dependencies"].([]interface{}); !ok || len(dependencies) == 0 {
			blockJSON, _ := json.Marshal(blockMap)

			// Publish the block for execution
			err = s.nc.Publish("block."+blockMap["id"].(string)+".execute", blockJSON)
			if err != nil {
				log.Printf("Error publishing block execution: %v", err)
				return
			}

			// Wait for the block to be completed before moving to the next one
			if err := s.waitForBlockCompletion(blockMap["id"].(string)); err != nil {
				log.Printf("Error waiting for block completion: %v", err)
				return
			}
		}
	}
}

// waitForBlockCompletion waits for a confirmation message that a block is completed
func (s *Scheduler) waitForBlockCompletion(blockID string) error {
	done := make(chan error)

	// Subscribe to the completion event for the current block
	sub, err := s.nc.Subscribe("block."+blockID+".completed", func(m *nats.Msg) {
		done <- nil // Signal that the block execution is complete
	})
	if err != nil {
		return err
	}

	// Timeout for block execution
	select {
	case <-done:
		log.Printf("Block %s execution completed", blockID)
		// Unsubscribe after the block is completed to prevent the issue of re-subscribing
		sub.Unsubscribe()
	case <-time.After(30 * time.Second): // Timeout of 30 seconds
		sub.Unsubscribe() // Unsubscribe in case of timeout
		return fmt.Errorf("block %s execution timeout", blockID)
	}

	return nil
}
