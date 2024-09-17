package block

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cu8code/9tail/internal/workflow"
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

func (f *FetchBlock) Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error) {
	// Extract the URL from input
	url, ok := input.(map[string]interface{})["url"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert the body to a string and store it in the workflow context
	result := string(body)
	ctx.Set(f.ID, result)

	// Signal the completion of this block by publishing a message to NATS
	err = f.NatsClient.Publish("block."+f.ID+".completed", []byte("done"))
	if err != nil {
		log.Printf("Error publishing block completion message: %v", err)
	}

	return result, nil
}
