package schema

// CreatePeriod : defines a request to create an period
type CreatePeriod struct {
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}
