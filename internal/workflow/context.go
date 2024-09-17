package workflow

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
)

type WorkflowContext struct {
	ID   string
	Data map[string]interface{}
	mu   sync.RWMutex
	nc   *nats.Conn
}

func NewWorkflowContext(id string, nc *nats.Conn) *WorkflowContext {
	return &WorkflowContext{
		ID:   id,
		Data: make(map[string]interface{}),
		nc:   nc,
	}
}

func (wc *WorkflowContext) Set(key string, value interface{}) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	wc.Data[key] = value
}

func (wc *WorkflowContext) Get(key string) (interface{}, bool) {
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	value, ok := wc.Data[key]
	return value, ok
}

func (wc *WorkflowContext) Publish() error {
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	data, err := json.Marshal(wc.Data)
	if err != nil {
		return err
	}
	return wc.nc.Publish(fmt.Sprintf("workflow.%s.context", wc.ID), data)
}

type ContextPublisher struct {
	nc *nats.Conn
}

func NewContextPublisher(nc *nats.Conn) *ContextPublisher {
	return &ContextPublisher{nc: nc}
}

func (cp *ContextPublisher) Start() error {
	_, err := cp.nc.Subscribe("workflow.*.context", func(msg *nats.Msg) {
		workflowID := strings.Split(msg.Subject, ".")[1]
		context, err := cp.getWorkflowContext(workflowID)
		if err != nil {
			log.Printf("Error getting workflow context: %v", err)
			return
		}
		response, _ := json.Marshal(context)
		msg.Respond(response)
	})
	return err
}

func (cp *ContextPublisher) getWorkflowContext(workflowID string) (map[string]interface{}, error) {
	// Implement logic to retrieve the workflow context
	// This could involve fetching from a database or in-memory store
	return map[string]interface{}{}, nil
}
