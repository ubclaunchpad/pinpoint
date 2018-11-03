package api

import (
	"testing"

	"github.com/ubclaunchpad/pinpoint/gateway/schema"
)

func TestCreateUser(t *testing.T) {
	_ := &schema.User{
		Name:     "Create",
		Email:    "user@test.com",
		Password: "password",
	}
}
