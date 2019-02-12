package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ubclaunchpad/pinpoint/protobuf/fakes"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"github.com/ubclaunchpad/pinpoint/utils"
	"google.golang.org/grpc"
)

func TestUserRouter_createUser(t *testing.T) {
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		u *request.CreateAccount
	}

	type errs struct {
		createUserFail error
	}

	tests := []struct {
		name     string
		args     args
		errs     errs
		wantCode int
	}{
		{"bad input", args{nil}, errs{nil}, http.StatusBadRequest},
		{"successfully create user", args{&request.CreateAccount{
			Name:     "Create",
			Email:    "user@test.com",
			Password: "password",
		}}, errs{nil}, http.StatusCreated},
		{"unsuccessfully create user", args{&request.CreateAccount{
			Name:     "s",
			Email:    "s",
			Password: "s",
		}}, errs{errors.New("Invalid arguments")}, http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &fakes.FakeCoreClient{}

			if tt.errs.createUserFail != nil {
				fake.CreateAccountStub = func(c context.Context, r *request.CreateAccount, opts ...grpc.CallOption) (*response.Message, error) {
					return nil, errors.New("invalid signup arguments")
				}
			}

			// create user router
			u := NewUserRouter(l, fake)

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
			req, err := http.NewRequest("POST", "/create", reader)
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

			if tt.wantCode == http.StatusCreated && fake.CreateAccountCallCount() < 1 {
				t.Error("uhh")
			}
		})
	}
}

func TestUserRouter_verify(t *testing.T) {
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		hash string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"no hash", args{""}, http.StatusBadRequest},
		{"ok hash", args{"tom"}, http.StatusAccepted},
		{"bad hash", args{"robert"}, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create user router
			fake := &fakes.FakeCoreClient{}
			u := NewUserRouter(l, fake)

			// set stub behaviour
			fake.VerifyStub = func(c context.Context, r *request.Verify, opts ...grpc.CallOption) (*response.Message, error) {
				if r.GetHash() == "tom" {
					return &response.Message{Message: "hello"}, nil
				}
				return nil, errors.New("unknown hash")
			}

			// create request
			req, err := http.NewRequest("GET", "/verify?hash="+tt.args.hash, nil)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}

			if tt.wantCode == http.StatusAccepted && fake.VerifyCallCount() < 1 {
				t.Error("expected call to core.Verify")
			}
		})
	}
}

func TestUserRouter_login(t *testing.T) {
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		email, password string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"no args", args{"", ""}, http.StatusBadRequest},
		{"regular user", args{"demo", "demopassword"}, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create user router
			fake := &fakes.FakeCoreClient{}
			u := NewUserRouter(l, fake)

			fake.LoginStub = func(c context.Context, r *request.Login, opts ...grpc.CallOption) (*response.Message, error) {
				if r.GetEmail() == "demo" && r.GetPassword() == "demopassword" {
					return &response.Message{Message: "user successfully logged in"}, nil
				}
				return nil, errors.New("user not authenticated")
			}

			// Create request
			req, err := http.NewRequest("POST", "/login?email="+tt.args.email+"&password="+tt.args.password, nil)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}

			if tt.wantCode == http.StatusOK && fake.LoginCallCount() < 1 {
				t.Error("expected call to core.Login")
			}
		})
	}
}
