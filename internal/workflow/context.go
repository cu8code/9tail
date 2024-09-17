package workflow

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
)

type WorkflowContext struct {
	ID string
	nc *nats.Conn
	kv nats.KeyValue
}

// NewWorkflowContext initializes a new WorkflowContext, passing JetStream KV store
func NewWorkflowContext(id string, nc *nats.Conn, kv nats.KeyValue) *WorkflowContext {
	return &WorkflowContext{
		ID: id,
		nc: nc,
		kv: kv,
	}
}

// Set updates the workflow context in JetStream directly
func (wc *WorkflowContext) Set(key string, value interface{}) error {
	// Load existing context data from JetStream
	entry, err := wc.kv.Get(wc.ID)
	if err != nil && err != nats.ErrKeyNotFound {
		return fmt.Errorf("failed to get context from KV store: %v", err)
	}

	data := make(map[string]interface{})
	if err == nil {
		// Unmarshal the existing context
		if err := json.Unmarshal(entry.Value(), &data); err != nil {
			return fmt.Errorf("failed to unmarshal context data: %v", err)
		}
	}

	// Set the new value in the context
	data[key] = value

	// Marshal the updated context and store it back in JetStream
	marshaledData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal context data: %v", err)
	}

	_, err = wc.kv.Put(wc.ID, marshaledData)
	return err
}

// Get retrieves the context value from JetStream for a given key
func (wc *WorkflowContext) Get(key string) (interface{}, bool, error) {
	// Load context data from JetStream
	entry, err := wc.kv.Get(wc.ID)
	if err != nil {
		if err == nats.ErrKeyNotFound {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed to get context from KV store: %v", err)
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(entry.Value(), &data); err != nil {
		return nil, false, fmt.Errorf("failed to unmarshal context data: %v", err)
	}

	// Retrieve the value for the key
	value, ok := data[key]
	return value, ok, nil
}

// Publish publishes the current context to the NATS topic
func (wc *WorkflowContext) Publish() error {
	// Load context data from JetStream
	entry, err := wc.kv.Get(wc.ID)
	if err != nil && err != nats.ErrKeyNotFound {
		return fmt.Errorf("failed to get context from KV store: %v", err)
	}

	var data map[string]interface{}
	if err == nil {
		if err := json.Unmarshal(entry.Value(), &data); err != nil {
			return fmt.Errorf("failed to unmarshal context data: %v", err)
		}
	}

	// Marshal and publish the data
	marshaledData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = wc.nc.Publish(fmt.Sprintf("workflow.%s.context", wc.ID), marshaledData)
	return err
}

// String returns the current context as a string by loading it from JetStream
func (wc *WorkflowContext) String() string {
	// Load context data from JetStream
	entry, err := wc.kv.Get(wc.ID)
	if err != nil {
		return fmt.Sprintf("Error retrieving context: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(entry.Value(), &data); err != nil {
		return fmt.Sprintf("Error unmarshaling context data: %v", err)
	}

	return fmt.Sprintf("%v", data)
}

type ContextPublisher struct {
	nc *nats.Conn
	js nats.JetStreamContext
	kv nats.KeyValue
}

func NewContextPublisher(nc *nats.Conn) (*ContextPublisher, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to create JetStream context: %v", err)
	}
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: "workflow_contexts",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create or get Key-Value store: %v", err)
	}
	return &ContextPublisher{
		nc: nc,
		js: js,
		kv: kv,
	}, nil
}

func (cp *ContextPublisher) Start() error {
	_, err := cp.nc.Subscribe("workflow.*.context", func(msg *nats.Msg) {
		workflowID := strings.Split(msg.Subject, ".")[1]
		context := NewWorkflowContext(workflowID, cp.nc, cp.kv)
		response, _ := json.Marshal(context)
		msg.Respond(response)
	})
	return err
}
