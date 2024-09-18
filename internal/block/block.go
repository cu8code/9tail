package block

import (
	"github.com/cu8code/9tail/internal/workflow"
)

// Block interface
type Block interface {
	GetID() string
	GetType() string
	Execute(input interface{}, ctx *workflow.WorkflowContext) (interface{}, error)
}

// Port types
type PortType string

const (
	StringType    PortType = "string"
	NumberType    PortType = "number"
	BooleanType   PortType = "boolean"
	NullType      PortType = "null"
	UndefinedType PortType = "undefined"
	ArrayType     PortType = "array"
	ObjectType    PortType = "object"
	MapType       PortType = "map"
	SetType       PortType = "set"
	ImageType     PortType = "image"
	VideoType     PortType = "video"
	AudioType     PortType = "audio"
	AnyType       PortType = "any"
	UnknownType   PortType = "unknown"
	VariantType   PortType = "variant"
	DataFrameType PortType = "data-frame"
)

// Port direction
type PortDirection string

const (
	InputDirection  PortDirection = "input"
	OutputDirection PortDirection = "output"
)

// Parameter structure
type Parameter struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	CanBe string `json:"can-be,omitempty"`
	From  struct {
		Source string `json:"source"`
	} `json:"from"`
}
