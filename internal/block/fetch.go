package block

import (
	"io/ioutil"
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
	url, ok := input.(map[string]interface{})["url"].(string)
	if !ok {
		return nil, ErrInvalidInput
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := string(body)
	ctx.Set(f.ID, result)
	return result, nil
}
