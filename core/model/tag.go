package model

// Tag model
type Tag struct {
	ApplicantID   string `json:"pk"`
	PeriodEventID string `json:"sk"`
	TagName       string `json:"tag_name"`
}
