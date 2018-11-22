package database

import (
	"reflect"
	"testing"

	"github.com/ubclaunchpad/pinpoint/core/model"
)

func TestDatabase_AddNewTag_GetTag(t *testing.T) {
	type args struct {
		c  *model.Club
		cu *model.ClubUser
		tg *model.Tag
	}
	type errs struct {
		addClub   bool
		getClub   bool
		addTag    bool
		getTag    bool
		deleteTag bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"valid", args{
			&model.Club{
				ID:          "1234",
				Name:        "Launchpad",
				Description: "1337 h4x0r",
			},
			&model.ClubUser{
				ClubID:   "1234",
				Email:    "abc@def.com",
				UserName: "Bob Ross",
				Role:     "President",
			},
			&model.Tag{
				ApplicantID:   "1234",
				PeriodEventID: "1234_1233",
				TagName:       "Sponsorship Team",
			},
		}, errs{false, false, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := NewTestDB()
			defer db.DeleteClub(tt.args.c.ID)
			if err := db.AddNewClub(tt.args.c, tt.args.cu); (err != nil) != tt.err.addClub {
				t.Errorf("Database.AddNewClub() error = %v, wantErr %v", err, tt.err.addClub)
			}

			club, err := db.GetClub(tt.args.c.ID)
			if (err != nil) != tt.err.getClub {
				t.Errorf("Database.GetClub() error = %v, wantErr %v", err, tt.err.getClub)
			}
			if !reflect.DeepEqual(tt.args.c, club) {
				t.Errorf("Failed to get expected club, expected: %+v, actual %+v", tt.args.c, club)
				return
			}

			if err := db.AddNewTag(tt.args.tg, tt.args.c); (err != nil) != tt.err.addTag {
				t.Errorf("Database.AddNewTag() error = %v, wantErr %v", err, tt.err.addTag)
			}

			tag, err := db.GetTag(tt.args.tg.ApplicantID, "1234", "1233", tt.args.c)
			if (err != nil) != tt.err.getTag {
				t.Errorf("Database.GetTag() error = %v, wantErr %v", err, tt.err.getTag)
			}
			if !reflect.DeepEqual(tt.args.tg, tag) {
				t.Errorf("Failed to get expected tag, expected: %+v, actual %+v", tt.args.tg, tag)
				return
			}

			if err := db.DeleteTag(tt.args.tg.ApplicantID, "1234", "1233", tt.args.c); (err != nil) != tt.err.deleteTag {
				t.Errorf("Database.DeleteTag() error = %v, wantErr %v", err, tt.err.deleteTag)
			}
		})
	}
}
