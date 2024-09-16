package block

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

type HookBlock struct {
	BaseBlock
}

func NewHookBlock(id string, nc *nats.Conn) *HookBlock {
	return &HookBlock{
		BaseBlock: BaseBlock{
			ID:         id,
			Type:       "hook",
			NatsClient: nc,
		},
	}
}

func (h *HookBlock) Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error) {
	log.Println("Executing hook block")
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidInput
	}

	url, ok := inputMap["url"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	payload, ok := inputMap["payload"]
	if !ok {
		return nil, ErrInvalidInput
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
