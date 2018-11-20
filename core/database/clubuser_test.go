package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/ubclaunchpad/pinpoint/core/model"
)

func TestDatabase_AddNewUser_GetUser(t *testing.T) {
	type args struct {
		u *model.User
		e *model.EmailVerification
	}
	type errs struct {
		addUser   bool
		getUser   bool
		getVerify bool
	}
	tests := []struct {
		name string
		args args
		err  errs
	}{
		{"invalid", args{
			&model.User{},
			&model.EmailVerification{
				Hash:   "asdf",
				Expiry: time.Now().Add(time.Hour),
			},
		}, errs{true, true, true}},
		{"valid", args{
			&model.User{
				Email: "abc@def.com",
				Name:  "Bob Ross",
				Salt:  "qwer1234",
			},
			&model.EmailVerification{
				Hash:   "asdf",
				Expiry: time.Now().Add(time.Hour),
			},
		}, errs{false, false, false}},
		{"expired", args{
			&model.User{
				Email: "abc@def.com",
				Name:  "Bob Ross",
				Salt:  "qwer1234",
			},
			&model.EmailVerification{
				Hash:   "asdf",
				Expiry: time.Now().Add(-time.Hour),
			},
		}, errs{false, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := NewTestDB()
			defer db.DeleteUser(tt.args.u.Email)
			if err := db.AddNewUser(tt.args.u, tt.args.e); (err != nil) != tt.err.addUser {
				t.Errorf("Database.AddNewUser() error = %v, wantErr %v", err, tt.err.addUser)
			}

			u, err := db.GetUser(tt.args.u.Email)
			if (err != nil) != tt.err.getUser {
				t.Errorf("Database.GetUser() error = %v, wantErr %v", err, tt.err.getUser)
				return
			}

			if !tt.err.getUser && !reflect.DeepEqual(tt.args.u, u) {
				t.Errorf("expected: %+v, actual %+v", tt.args.u, u)
				return
			}

			v, err := db.GetEmailVerification(tt.args.e.Hash)
			if (err != nil) != tt.err.getVerify {
				t.Errorf("Database.GetEmailVerification() error = %v, wantErr %v", err, tt.err.getVerify)
				return
			}

			if !tt.err.getVerify && (v.Hash != tt.args.e.Hash || v.Email != tt.args.e.Email) {
				t.Errorf("expected: %+v, actual %+v", tt.args.e, v)
				return
			}
		})
	}
}

func TestDatabase_AddNewClub_GetClub(t *testing.T) {
	type args struct {
		c  *model.Club
		cu *model.ClubUser
	}
	type errs struct {
		addClub      bool
		getClub      bool
		getClubUsers bool
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
		}, errs{false, false, false}},
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
				t.Errorf("Failed to get expect club, expected: %+v, actual %+v", tt.args.c, club)
				return
			}

			_, err = db.GetAllClubUsers(tt.args.c.ID)
			if (err != nil) != tt.err.getClubUsers {
				t.Errorf("Database.GetAllClubUsers() error = %v, wantErr %v", err, tt.err.getClubUsers)
			}
		})
	}
}
