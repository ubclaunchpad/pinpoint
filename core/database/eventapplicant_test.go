package database

import (
	"reflect"
	"testing"

	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

var club = &models.Club{
	ClubID:      "1234",
	Name:        "Launchpad",
	Description: "1337 h4x0r",
}
var user = &models.ClubUser{
	ClubID: "1234",
	Email:  "abc@def.com",
	Name:   "Bob Ross",
	Role:   "Artist",
}

func TestDatabase_AddNewEvent_GetEvent(t *testing.T) {
	db, _ := NewTestDB()
	db.AddNewClub(club, user)
	type args struct {
		clubID string
		period string
		event  *models.EventProps
	}
	type errs struct {
		addEvent bool
		getEvent bool
		// getEvents bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"invalid", args{
			"",
			"",
			&models.EventProps{},
		}, errs{true, true}},
		{"valid", args{
			"1234",
			"Winter 2019",
			&models.EventProps{
				Period:  "Winter 2019",
				EventID: "001",
				Name:    "Recruiting",
			},
		}, errs{false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer db.DeleteEvent(tt.args.clubID, tt.args.period, tt.args.event.EventID)

			if err := db.AddNewEvent(tt.args.clubID, tt.args.event); (err != nil) != tt.err.addEvent {
				t.Errorf("Database.AddNewEvent() error = %v, wantErr %v", err, tt.err.addEvent)
			}

			event, err := db.GetEvent(tt.args.clubID, tt.args.period, tt.args.event.EventID)
			if (err != nil) != tt.err.getEvent {
				t.Errorf("Database.GetEvent() error = %v, wantErr %v", err, tt.err.getEvent)
				return
			}
			if !tt.err.getEvent && !reflect.DeepEqual(tt.args.event, event) {
				t.Errorf("expected: %+v, actual %+v", tt.args.event, event)
				return
			}

			//// Not sure how to get events. How to relate a club to a period?
			// _, err = db.GetEvents(tt.args.clubID, tt.args.period)
			// if (err != nil) != tt.err.getEvents {
			// 	t.Errorf("Database.GetEvents() error = %v, wantErr %v", err, tt.err.getEvents)
			// 	return
			// }
		})
	}
}

func TestDatabase_Applicant(t *testing.T) {
	db, _ := NewTestDB()
	db.AddNewClub(club, user)
	type args struct {
		clubID    string
		applicant *models.Applicant
	}
	type errs struct {
		addApplicant  bool
		getApplicant  bool
		getApplicants bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"invalid", args{
			"",
			&models.Applicant{},
		}, errs{true, true, true}},
		{"valid", args{
			"1234",
			&models.Applicant{
				Period: "Winter Semester",
				Email:  user.Email,
				Name:   user.Name,
			},
		}, errs{false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer db.DeleteApplicant(tt.args.clubID, tt.args.applicant.Period, tt.args.applicant.Email)
			if err := db.AddNewApplicant(tt.args.clubID, tt.args.applicant); (err != nil) != tt.err.addApplicant {
				t.Errorf("Database.AddNewClub() error = %v, wantErr %v", err, tt.err.addApplicant)
			}

			_, err := db.GetApplicant(tt.args.clubID, tt.args.applicant.Period, tt.args.applicant.Email)
			if (err != nil) != tt.err.getApplicant {
				t.Errorf("Database.GetClub() error = %v, wantErr %v", err, tt.err.getApplicant)
			}

			_, err = db.GetApplicants(tt.args.clubID, tt.args.applicant.Period)
			if (err != nil) != tt.err.getApplicants {
				t.Errorf("Database.GetApplicants() error = %v, wantErr %v", err, tt.err.getApplicants)
			}
		})
	}
}

func TestDatabase_Application(t *testing.T) {
	db, _ := NewTestDB()
	db.AddNewClub(club, user)
	type args struct {
		clubID      string
		application *models.Application
	}
	type errs struct {
		addApplication  bool
		getApplication  bool
		getApplications bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"invalid", args{
			"",
			&models.Application{},
		}, errs{true, true, true}},
		{"valid", args{
			"1234",
			&models.Application{
				Period:  "Winter 2019",
				EventID: "001",
				Email:   "abc@def.com",
				Name:    "Recruiting",
				Entries: map[string]*models.FieldEntry{},
			},
		}, errs{false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer db.DeleteApplication(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID, tt.args.application.Email)
			if err := db.AddNewApplication(tt.args.clubID, tt.args.application); (err != nil) != tt.err.addApplication {
				t.Errorf("Database.AddNewClub() error = %v, wantErr %v", err, tt.err.addApplication)
			}

			_, err := db.GetApplication(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID, tt.args.application.Email)
			if (err != nil) != tt.err.getApplication {
				t.Errorf("Database.GetClub() error = %v, wantErr %v", err, tt.err.getApplication)
			}

			_, err = db.GetApplications(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID)
			if (err != nil) != tt.err.getApplications {
				t.Errorf("Database.GetApplications() error = %v, wantErr %v", err, tt.err.getApplications)
			}
		})
	}
}

func TestDatabase_AddTag(t *testing.T) {
	db, _ := NewTestDB()
	db.AddNewClub(club, user)

	type args struct {
		clubID string
		tag    *models.Tag
	}
	type errs struct {
		addTag  bool
		getTags bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"invalid", args{
			"",
			&models.Tag{},
		}, errs{true, true}},
		{"valid", args{
			"1234",
			&models.Tag{Period: "Winter 2019", TagName: "Designer"},
		}, errs{false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer db.DeleteTag(tt.args.clubID, tt.args.tag.Period, tt.args.tag.TagName)
			if err := db.AddTag(tt.args.clubID, tt.args.tag); (err != nil) != tt.err.addTag {
				t.Errorf("Database.AddTag() error = %v, wantErr %v", err, tt.err.addTag)
			}

			_, err := db.GetTags(tt.args.clubID, tt.args.tag.Period)
			if (err != nil) != tt.err.getTags {
				t.Errorf("Database.GetTags() error = %v, wantErr %v", err, tt.err.getTags)
			}
		})
	}
}
