package database

import "github.com/ubclaunchpad/pinpoint/protobuf/models"

type eventItem struct {
	PeidPK      string          `json:"pk"`
	PeidSK      string          `json:"sk"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Fields      []*models.Field `json:"fields"`
}

type applicantItem struct {
	Period string   `json:"pk"`
	Email  string   `json:"sk"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
}

type applicationItem struct {
	Peid    string                        `json:"pk"`
	Email   string                        `json:"sk"`
	Name    string                        `json:"name"`
	Entries map[string]*models.FieldEntry `json:"entries"`
}

type tagItem struct {
	Period  string `json:"pk"`
	TagName string `json:"sk"`
}

func newDBEvent(event *models.Event) *eventItem {
	peid := prefixPeriodEventID(event.Period, event.EventID)
	return &eventItem{
		PeidPK:      peid,
		PeidSK:      peid,
		Name:        event.Name,
		Description: event.Description,
		Fields:      event.Fields,
	}
}

func newEvent(item *eventItem) *models.Event {
	p, e := getPeriodAndEventID(item.PeidPK)
	return &models.Event{
		Period:      p,
		EventID:     e,
		Name:        item.Name,
		Description: item.Description,
		Fields:      item.Fields,
	}
}

func newDBApplicant(applicant *models.Applicant) *applicantItem {
	p := prefixPeriodID(applicant.Period)
	return &applicantItem{
		Period: p,
		Email:  prefixApplicantID(applicant.Email),
		Name:   applicant.Name,
		Tags:   applicant.Tags,
	}
}

func newApplicant(item *applicantItem) *models.Applicant {
	return &models.Applicant{
		Period: removePrefix(item.Period),
		Email:  removePrefix(item.Email),
		Name:   item.Name,
		Tags:   item.Tags,
	}
}

func newDBTag(tag *models.Tag) *tagItem {
	return &tagItem{
		Period:  prefixPeriodID(tag.Period),
		TagName: prefixTag(tag.TagName),
	}
}

func newTag(item *tagItem) *models.Tag {
	return &models.Tag{
		Period:  removePrefix(item.Period),
		TagName: removePrefix(item.TagName),
	}
}

func newDBApplication(application *models.Application) *applicationItem {
	peid := prefixPeriodEventID(application.Period, application.EventID)
	return &applicationItem{
		Peid:    peid,
		Email:   prefixApplicantID(application.Email),
		Name:    application.Name,
		Entries: application.Entries,
	}
}

func newApplication(item *applicationItem) *models.Application {
	p, e := getPeriodAndEventID(item.Peid)
	return &models.Application{
		Period:  p,
		EventID: e,
		Email:   removePrefix(item.Email),
		Name:    item.Name,
	}
}
