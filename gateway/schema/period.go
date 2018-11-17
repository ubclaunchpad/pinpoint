package schema

import "time"

// CreatePeriod defines a request to create an period
type CreatePeriod struct {
	Name       string
	StartYear  int
	StartMonth time.Month
	EndYear    int
	EndMonth   time.Month
}
