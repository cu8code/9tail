package block

import (
	"encoding/json"
	"log"

	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

type ResponseBlock struct {
	BaseBlock
}

func NewResponseBlock(id string, nc *nats.Conn) *ResponseBlock {
	return &ResponseBlock{
		BaseBlock: BaseBlock{
			ID:         id,
			Type:       "response",
			NatsClient: nc,
		},
	}
}

func (r *ResponseBlock) Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error) {
	log.Println("Executing response block")
	// For a response block, we might just want to format the input
	// into a standardized response structure
	response := map[string]interface{}{
		"status": "success",
		"data":   input,
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return string(jsonResponse), nil
}
