package database

import (
	"testing"

	"github.com/ubclaunchpad/pinpoint/core/model"

	"github.com/ubclaunchpad/pinpoint/utils"
)

// newTestDB is a utility function to create a connection to a local dynamodb
// instance
func newTestDB(t *testing.T) *Database {
	s, err := utils.AWSSession(utils.AWSConfig(true))
	if err != nil {
		t.Errorf("setup failed: %s", err.Error())
	}
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Errorf("setup failed: %s", err.Error())
	}

	// attempt to create database with connection
	db, err := New(s, l)
	if err != nil {
		t.Error(err)
	}
	return db
}

func TestNew(t *testing.T) {
	newTestDB(t)
}

func TestClubuser_Dao(t *testing.T) {
	db := newTestDB(t)
	u := &model.User{Email: "abc@def.com", Name: "Bob Ross", Salt: "qwer1234"}
	c := &model.Club{
		ID:          "1234",
		Name:        "Launchpad",
		Description: "1337 h4x0r",
		Periods:     []string{"2018-2019"},
	}
	cu := &model.Clubuser{
		ClubID:   "1234",
		Email:    "abc@def.com",
		UserName: "Bob Ross",
		Role:     "President",
	}
	db.AddNewUser(u)
	res1, _ := db.GetUser(u.Email)
	t.Log(*res1)
	db.AddNewClub(c, cu)
	res2, _ := db.GetClub(c.ID)
	t.Log(*res2)
	res3, _ := db.GetAllClubusers(c.ID)
	t.Log(*res3)
}
