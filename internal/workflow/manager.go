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
