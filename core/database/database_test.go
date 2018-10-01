package database

import (
	"testing"

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
