package schema

// CreateEvent : defines a request to create an event
type CreateEvent struct {
	Name   string
	Fields []byte
}
