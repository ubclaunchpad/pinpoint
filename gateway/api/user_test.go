package api

import (
	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := &schema.User{
		Name: "Create"
		Email: "user@test.com"
		Password: "password"
	}
}