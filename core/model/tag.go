package model

// Tag model
type Tag struct {
	Applicant_ID    string `json:"pk"`
	Period_Event_ID string `json:"sk"`
	Tag_Name        string `json:"tag_name"`
}
