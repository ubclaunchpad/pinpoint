package database

import (
	"reflect"
	"testing"

	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

var club = &models.Club{
	ClubID:      "1234",
	Description: "1337 h4x0r",
}
var user = &models.ClubUser{
	ClubID: "1234",
	Email:  "abc@def.com",
	Name:   "Bob Ross",
	Role:   "Artist",
}

func TestDatabase_AddNewEvent_GetEvent(t *testing.T) {
	type args struct {
		clubID string
		event  *models.EventProps
	}
	type errs struct {
		addEvent  bool
		getEvent  bool
		getEvents bool
	}
	tests := []struct {
		name      string
		args      args
		err       errs
		wantEvent bool
	}{
		{"invalid event", args{
			"1234",
			&models.EventProps{},
		}, errs{true, true, true}, false},
		{"invalid club id", args{
			"",
			&models.EventProps{
				Period:  "Winter 2019",
				EventID: "001",
				Name:    "Recruiting",
			},
		}, errs{true, true, true}, false},
		{"valid", args{
			"1234",
			&models.EventProps{
				Period:  "Winter 2019",
				EventID: "001",
				Name:    "Recruiting",
			},
		}, errs{false, false, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := newTestDB(t)
			defer db.DeleteClub(club.ClubID)
			db.AddNewClub(club, user)
			defer db.DeleteEvent(tt.args.clubID, tt.args.event.Period, tt.args.event.EventID)
			if err := db.AddNewEvent(tt.args.clubID, tt.args.event); (err != nil) != tt.err.addEvent {
				t.Errorf("Database.AddNewEvent() error = %v, wantErr %v", err, tt.err.addEvent)
			}

			event, err := db.GetEvent(tt.args.clubID, tt.args.event.Period, tt.args.event.EventID)
			if (err != nil) != tt.err.getEvent {
				t.Errorf("Database.GetEvent() error = %v, wantErr %v", err, tt.err.getEvent)
				return
			}
			if tt.wantEvent {
				if !tt.err.getEvent && !reflect.DeepEqual(tt.args.event, event) {
					t.Errorf("expected: %+v, actual: %+v", tt.args.event, event)
					return
				}
			} else {
				if event != nil {
					t.Errorf("Didn't expect event, got: %+v", event)
				}
			}

			events, err := db.GetEvents(tt.args.clubID, tt.args.event.Period)
			if (err != nil) != tt.err.getEvents {
				t.Errorf("Database.GetEvents() error = %v, wantErr %v", err, tt.err.getEvents)
				return
			}
			if tt.wantEvent {
				expected := []*models.EventProps{tt.args.event}
				if !tt.err.getEvent && !reflect.DeepEqual(expected, events) {
					t.Errorf("expected: %+v, actual: %+v", expected, events)
					return
				}
			} else {
				if len(events) > 0 {
					t.Errorf("Didn't expect events, got: %+v", events)
				}
			}
		})
	}
}

func TestDatabase_Applicant(t *testing.T) {
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
		name          string
		args          args
		err           errs
		wantApplicant bool
	}{
		{"invalid applicant", args{
			"",
			&models.Applicant{},
		}, errs{true, true, true}, false},
		{"invalid club id", args{
			"",
			&models.Applicant{
				Period: "Winter Semester",
				Email:  user.Email,
				Name:   user.Name,
			},
		}, errs{true, true, true}, false},
		{"valid", args{
			"1234",
			&models.Applicant{
				Period: "Winter Semester",
				Email:  user.Email,
				Name:   user.Name,
			},
		}, errs{false, false, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := newTestDB(t)
			defer db.DeleteClub(club.ClubID)
			db.AddNewClub(club, user)
			defer db.DeleteApplicant(tt.args.clubID, tt.args.applicant.Period, tt.args.applicant.Email)
			if err := db.AddNewApplicant(tt.args.clubID, tt.args.applicant); (err != nil) != tt.err.addApplicant {
				t.Errorf("Database.AddNewClub() error = %v, wantErr %v", err, tt.err.addApplicant)
			}

			app, err := db.GetApplicant(tt.args.clubID, tt.args.applicant.Period, tt.args.applicant.Email)
			if (err != nil) != tt.err.getApplicant {
				t.Errorf("Database.GetClub() error = %v, wantErr %v", err, tt.err.getApplicant)
			}
			if tt.wantApplicant {
				if !reflect.DeepEqual(tt.args.applicant, app) {
					t.Errorf("Failed to get expect applicant, expected: %+v, actual: %+v", tt.args.applicant, app)
					return
				}
			}

			apps, err := db.GetApplicants(tt.args.clubID, tt.args.applicant.Period)
			if (err != nil) != tt.err.getApplicants {
				t.Errorf("Database.GetApplicants() error = %v, wantErr %v", err, tt.err.getApplicants)
			}
			if tt.wantApplicant {
				if !reflect.DeepEqual([]*models.Applicant{tt.args.applicant}, apps) {
					t.Errorf("Failed to get expect applicants, expected: %+v, actual: %+v", []*models.Applicant{tt.args.applicant}, apps)
					return
				}
			} else {
				if len(apps) > 0 {
					t.Errorf("Didn't expect tags, got: %+v", apps)
				}
			}

		})
	}
}

func TestDatabase_Application(t *testing.T) {
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
		name            string
		args            args
		err             errs
		wantApplication bool
	}{
		{"invalid applicant", args{
			"1234",
			&models.Application{},
		}, errs{true, true, true}, false},
		{"invalid club id", args{
			"",
			&models.Application{
				Period:  "Winter 2019",
				EventID: "001",
				Email:   "abc@def.com",
				Name:    "Recruiting",
				Entries: map[string]*models.FieldEntry{},
			},
		}, errs{true, true, true}, false},
		{"valid", args{
			"1234",
			&models.Application{
				Period:  "Winter 2019",
				EventID: "001",
				Email:   "abc@def.com",
				Name:    "Recruiting",
				Entries: map[string]*models.FieldEntry{},
			},
		}, errs{false, false, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := newTestDB(t)
			defer db.DeleteClub(club.ClubID)
			db.AddNewClub(club, user)
			defer db.DeleteApplication(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID, tt.args.application.Email)
			if err := db.AddNewApplication(tt.args.clubID, tt.args.application); (err != nil) != tt.err.addApplication {
				t.Errorf("Database.AddNewApplication() error = %v, wantErr %v", err, tt.err.addApplication)
			}

			app, err := db.GetApplication(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID, tt.args.application.Email)
			if (err != nil) != tt.err.getApplication {
				t.Errorf("Database.GetApplication() error = %v, wantErr %v", err, tt.err.getApplication)
			}
			var checkApp func(*models.Application, *models.Application) bool
			checkApp = func(actual *models.Application, expected *models.Application) bool {
				if actual.Period != app.Period || actual.EventID != app.EventID || actual.Email != app.Email {
					return false
				}
				return true
			}
			if tt.wantApplication {
				if !checkApp(tt.args.application, app) {
					t.Errorf("Failed to get expected application, expected: %+v, actual: %+v", *tt.args.application, *app)
					return
				}
			}

			apps, err := db.GetApplications(tt.args.clubID, tt.args.application.Period, tt.args.application.EventID)
			if (err != nil) != tt.err.getApplications {
				t.Errorf("Database.GetApplications() error = %v, wantErr %v", err, tt.err.getApplications)
			}
			if tt.wantApplication {
				expected := []*models.Application{tt.args.application}
				for i := 0; i < len(apps); i++ {
					if !checkApp(expected[i], apps[i]) {
						t.Errorf("Failed to get expected applications, expected: %+v, actual: %+v", expected, apps)
						return
					}
				}
			} else {
				if len(apps) > 0 {
					t.Errorf("Didn't expect applications, got: %+v", apps)
				}
			}
		})
	}
}

func TestDatabase_AddTag(t *testing.T) {
	type args struct {
		clubID string
		tag    *models.Tag
	}
	type errs struct {
		addTag  bool
		getTags bool
	}
	tests := []struct {
		name    string
		args    args
		err     errs
		wantTag bool
	}{
		{"invalid tag", args{
			"1234",
			&models.Tag{},
		}, errs{true, true}, false},
		{"invalid club id", args{
			"",
			&models.Tag{
				Period:  "Winter 2019",
				TagName: "Designer",
			},
		}, errs{true, true}, false},
		{"valid everything", args{
			"1234",
			&models.Tag{
				Period:  "Winter 2019",
				TagName: "Designer",
			},
		}, errs{false, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := newTestDB(t)
			defer db.DeleteClub(club.ClubID)
			db.AddNewClub(club, user)
			defer db.DeleteTag(tt.args.clubID, tt.args.tag.Period, tt.args.tag.TagName)
			if err := db.AddTag(tt.args.clubID, tt.args.tag); (err != nil) != tt.err.addTag {
				t.Errorf("Database.AddTag() error = %v, wantErr %v", err, tt.err.addTag)
			}

			tags, err := db.GetTags(tt.args.clubID, tt.args.tag.Period)
			if (err != nil) != tt.err.getTags {
				t.Errorf("Database.GetTags() error = %v, wantErr %v", err, tt.err.getTags)
				return
			}

			if tt.wantTag {
				expected := []*models.Tag{tt.args.tag}
				if !reflect.DeepEqual(expected, tags) {
					t.Errorf("Failed to get expect tags, expected: %+v, actual: %+v", expected, tags)
					return
				}
			} else {
				if len(tags) > 0 {
					t.Errorf("Didn't expect tags, got: %+v", tags)
				}
			}
		})
	}
}
