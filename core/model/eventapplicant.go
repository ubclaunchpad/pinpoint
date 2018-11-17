package model

// Event represents data for an event
type Event struct {
	Period      string
	EventID     string
	Name        string
	Description string
}

// Applicant represents all the applicant information
type Applicant struct {
	Period  string
	EventID string
	Email   string
	Name    string
	Info    map[string]interface{}
}
