package block

import (
	"github.com/cu8code/9tail/internal/workflow"
	"github.com/nats-io/nats.go"
)

type Block interface {
	GetID() string
	GetType() string
	Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error)
}

type BaseBlock struct {
	ID         string
	Type       string
	NatsClient *nats.Conn
}

func (b *BaseBlock) GetID() string {
	return b.ID
}

func (b *BaseBlock) GetType() string {
	return b.Type
}
