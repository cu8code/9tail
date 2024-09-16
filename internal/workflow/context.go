package workflow

import (
	"encoding/json"
	"fmt"
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
