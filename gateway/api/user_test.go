package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"github.com/ubclaunchpad/pinpoint/protobuf/fakes"
	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestUserRouter_createUser(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		u *schema.CreateUser
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"bad input", args{nil}, http.StatusBadRequest},
		{"successfully create user", args{&schema.CreateUser{
			Name:      "Create",
			Email:     "user@test.com",
			Password:  "password",
			CPassword: "cpassword",
			ESub:      true,
		}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &fakes.FakeCoreClient{}

			// set up mock CreateAccount
			fake.CreateAccount(nil, nil, nil)

			// create user router
			u := newUserRouter(l, fake)

			// create request
			var b []byte
			var err error
			if tt.args.u != nil {
				if b, err = json.Marshal(tt.args.u); err != nil {
					t.Error(err)
					return
				}
			}
			reader := bytes.NewReader(b)
			req, err := http.NewRequest("POST", "/create_user", reader)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()

			// Serve request
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}
		})
	}
}
