package schema

import "encoding/json"

// CreateEvent defines a request to create an event
type CreateEvent struct {
	Name   string       `json:"name"`
	Fields []FieldProps `json:"fields"`
}

const (
	// FieldTypeLongText is a type of a field of an event
	FieldTypeLongText string = "long_text"
	// FieldTypeShortText is a type of a field of an event
	FieldTypeShortText string = "short_text"
)

// FieldProps defines a field of an event
type FieldProps struct {
	Type string `json:"type"`

	// Properties is one of protobuf ...
	Properties json.RawMessage `json:"properties"`
}
