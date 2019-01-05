package schema

import "encoding/json"

// CreateEvent defines a request to create an event
type CreateEvent struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

// Field defines a field of an event
type Field struct {
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}
