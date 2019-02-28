package schema

import "encoding/json"

// CreateEvent defines a request to create an event
type CreateEvent struct {
	Name        string       `json:"name"`
	EventID     string       `json:"event_id"`
	Description string       `json:"description"`
	Fields      []FieldProps `json:"fields"`
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

	// Properties is one of LongText or ShortText in protobuf
	Properties json.RawMessage `json:"properties"`
}
