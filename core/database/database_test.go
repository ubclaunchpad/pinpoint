package database

import (
	"reflect"
	"testing"
	//"time"
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

func TestClubUser(t *testing.T) {
	db := newTestDB(t)
	u := &model.User{
		Email: "abc@def.com",
		Name:  "Bob Ross",
		Salt:  "qwer1234",
	}
	c := &model.Club{
		ID:          "1234",
		Name:        "Launchpad",
		Description: "1337 h4x0r",
	}
	cu := &model.ClubUser{
		ClubID:   "1234",
		Email:    "abc@def.com",
		UserName: "Bob Ross",
		Role:     "President",
	}
	cus := []*model.ClubUser{cu}

	err := db.AddNewUser(u)
	if err != nil {
		t.Errorf("Failed to add new user: %s", err.Error())
		t.FailNow()
	}

	userActual, err := db.GetUser(u.Email)
	if err != nil {
		t.Errorf("Failed to get user: %s", err.Error())
		t.FailNow()
	}
	if !reflect.DeepEqual(u, userActual) {
		t.Errorf("Failed to get expect user, expected: %+v, actual %+v", u, userActual)
		t.FailNow()
	}

	err = db.AddNewClub(c, cu)
	if err != nil {
		t.Errorf("Failed to add new user: %s", err.Error())
		t.FailNow()
	}

	clubActual, err := db.GetClub(c.ID)
	if err != nil {
		t.Errorf("Failed to add new user: %s", err.Error())
		t.FailNow()
	}
	if !reflect.DeepEqual(c, clubActual) {
		t.Errorf("Failed to get expect club, expected: %+v, actual %+v", c, clubActual)
		t.FailNow()
	}

	clubUsersActual, err := db.GetAllClubUsers(c.ID)
	if err != nil {
		t.Errorf("Failed to add new user: %s", err.Error())
		t.FailNow()
	}
	for i := range clubUsersActual {
		if !reflect.DeepEqual(*cus[i], *clubUsersActual[i]) {
			t.Errorf("Failed to get expect club, expected: %+v, actual %+v", *cus[i], *clubUsersActual[i])
			t.FailNow()
		}
	}

	err = db.DeleteClub(c.ID)
	if err != nil {
		t.Errorf("Failed to delete club %s", err.Error())
		t.FailNow()
	}

	err = db.DeleteUser(u.Email)
	if err != nil {
		t.Errorf("Failed to delete user %s", err.Error())
		t.FailNow()
	}




}

	/* Tag Tests */
func TestTag(t *testing.T) {
	db := newTestDB(t)
	tag := &model.Tag{
		Applicant_ID: "1234",
		Period_Event_ID: "1234_1233",
		Tag_Name: "Sponsorship Team",
	}
	err := db.AddNewTag(tag)
	if err != nil {
		t.Errorf("Failed to add new tag: %s", err.Error())
		t.FailNow()
	}

	tagActual, err := db.GetTag(tag.Applicant_ID, "1234" , "1233")
	if err != nil {
		t.Errorf("Failed to get tag: %s", err.Error())
		t.FailNow()
	}
	if !reflect.DeepEqual(tag, tagActual) {
		t.Errorf("tag collected is not as expected, expected: %+v, actual %+v", tag, tagActual)
		t.FailNow()
	}


	tagNew, err := db.ChangeTagName(tag.Applicant_ID, "1234" , "1233",  "Marketing Team")
	if err != nil {
		t.Errorf("Failed to change tag %s", err.Error())
		t.Errorf("Failed to change tag %+v", tagNew)
		t.FailNow()
	}

	tag.Tag_Name = "Marketing Team"

	tagActual, err = db.GetTag(tag.Applicant_ID, "1234" , "1233")
	if err != nil {
		t.Errorf("Failed to get tag: %s", err.Error())
		t.FailNow()
	}
	if !reflect.DeepEqual(tagNew, tag) {
		t.Errorf("tag collected is not as expected, expected: %+v, actual %+v", tag, tagNew)
		t.FailNow()
	}

	err = db.DeleteTag(tag.Applicant_ID, "1234" , "1233")
	if err != nil {
		t.Errorf("Failed to delete tag %s", err.Error())
		t.FailNow()
	}
}
