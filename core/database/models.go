package database

import (
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

type userItem struct {
	EmailPK  string `json:"pk"`
	EmailSK  string `json:"sk"`
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	Verified bool   `json:"verified"`
}

type clubItem struct {
	ClubIDPK    string `json:"pk"`
	ClubIDSK    string `json:"sk"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type clubUserItem struct {
	ClubID string `json:"pk"`
	Email  string `json:"sk"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type emailVerificationItem struct {
	Email  string `json:"pk"`
	Hash   string `json:"sk"`
	Expiry int64  `json:"expiry"`
}

type eventItem struct {
	PeidPK      string                          `json:"pk"`
	PeidSK      string                          `json:"sk"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Fields      []*models.EventProps_FieldProps `json:"fields"`
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

func newDBUser(u *models.User) *userItem {
	e := prefixUserEmail(u.Email)
	return &userItem{
		EmailPK:  e,
		EmailSK:  e,
		Name:     u.Name,
		Hash:     u.Hash,
		Verified: u.Verified,
	}
}

func newUser(i *userItem) *models.User {
	e := removePrefix(i.EmailPK)
	return &models.User{
		Email:    e,
		Name:     i.Name,
		Hash:     i.Hash,
		Verified: i.Verified,
	}
}

func newDBClub(c *models.Club) *clubItem {
	id := prefixClubID(c.ClubID)
	return &clubItem{
		ClubIDPK:    id,
		ClubIDSK:    id,
		Name:        c.Name,
		Description: c.Description,
	}
}

func newClub(i *clubItem) *models.Club {
	id := removePrefix(i.ClubIDPK)
	return &models.Club{
		ClubID:      id,
		Name:        i.Name,
		Description: i.Description,
	}
}

func newDBClubUser(cu *models.ClubUser) *clubUserItem {
	id := prefixClubID(cu.ClubID)
	e := prefixUserEmail(cu.Email)
	return &clubUserItem{
		ClubID: id,
		Email:  e,
		Name:   cu.Name,
		Role:   cu.Role,
	}
}

func newClubUser(i *clubUserItem) *models.ClubUser {
	id := removePrefix(i.ClubID)
	e := removePrefix(i.Email)
	return &models.ClubUser{
		ClubID: id,
		Email:  e,
		Name:   i.Name,
		Role:   i.Role,
	}
}

func newDBEmailVerification(ev *models.EmailVerification) *emailVerificationItem {
	e := prefixUserEmail(ev.Email)
	h := prefixVerificationHash(ev.Hash)
	return &emailVerificationItem{
		Email:  e,
		Hash:   h,
		Expiry: ev.Expiry,
	}
}

func newEmailVerification(i *emailVerificationItem) *models.EmailVerification {
	e := removePrefix(i.Email)
	h := removePrefix(i.Hash)
	return &models.EmailVerification{
		Email:  e,
		Hash:   h,
		Expiry: i.Expiry,
	}
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
