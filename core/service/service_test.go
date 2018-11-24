package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ubclaunchpad/pinpoint/core/database/mocks"
	"github.com/ubclaunchpad/pinpoint/core/model"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestNew(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	acfg, _ := utils.AWSSession(utils.AWSConfig(true))
	badcfg, _ := session.NewSession(aws.NewConfig())
	type args struct {
		awsConfig client.ConfigProvider
		opts      Opts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid aws config", args{badcfg, Opts{}}, true},
		{"invalid tls config", args{acfg, Opts{
			TLSOpts: TLSOpts{
				CertFile: "../../dev/certs/asdf.crt",
			},
		}}, true},
		{"valid no tls", args{acfg, Opts{}}, false},
		{"valid tls config", args{acfg, Opts{
			TLSOpts: TLSOpts{
				CertFile: "../../dev/certs/127.0.0.1.crt",
				KeyFile:  "../../dev/certs/127.0.0.1.key",
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.awsConfig, l, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_Run(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	aws, _ := utils.AWSSession(utils.AWSConfig(true))
	s, err := New(aws, l, Opts{})
	if err != nil {
		t.Error(err)
		return
	}

	// stub run
	go s.Run("", "")
	time.Sleep(time.Millisecond)
	s.Stop()
}

func TestService_GetStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		req *request.Status
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Status
		wantErr bool
	}{
		{
			"get callback",
			args{nil, &request.Status{}},
			&response.Status{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.GetStatus(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Handshake(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"background context", args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := utils.NewLogger(true)
			if err != nil {
				t.Error(err)
				return
			}
			s := &Service{l: l}
			if tt.args.ctx == nil {
				tt.args.ctx = context.Background()
			}
			_, err = s.Handshake(tt.args.ctx, &request.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Handshake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_CreateAccount(t *testing.T) {
	type args struct {
		req *request.CreateAccount
	}
	tests := []struct {
		name    string
		args    args
		DBErr   bool
		wantErr bool
	}{
		{
			"get password requirement error",
			args{&request.CreateAccount{
				Email:    "test@pinpoint.com",
				Name:     "test",
				Password: "1234"}},
			false, true,
		},
		{
			"get email requirement error",
			args{&request.CreateAccount{
				Email:    "test",
				Name:     "test",
				Password: "1234pass."}},
			false, true,
		},
		{
			"db failure",
			args{&request.CreateAccount{
				Email:    "test@gmail.com",
				Name:     "test",
				Password: "1234pass."}},
			true, true,
		},
		{
			"success",
			args{&request.CreateAccount{
				Email:    "test@gmail.com",
				Name:     "test",
				Password: "1234pass."}},
			false, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &mocks.FakeDBClient{}
			if tt.DBErr {
				fk.AddNewUserReturns(errors.New("oh no"))
			}
			s := &Service{db: fk}
			_, err := s.CreateAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && fk.AddNewUserCallCount() < 1 {
				t.Error("expected call to AddNewUser")
			}
		})
	}
}

func TestService_Verify(t *testing.T) {
	expectedHash := "NmSdjumzjHOF7IAnafAK74LAPug="
	type args struct {
		ctx context.Context
		req *request.Verify
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get success given expected hash",
			args{nil, &request.Verify{Hash: expectedHash}},
			false,
		},
		{
			"get error",
			args{nil, &request.Verify{Hash: "incorrect hash"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &mocks.FakeDBClient{}
			fk.GetEmailVerificationStub = func(hash string) (*model.EmailVerification, error) {
				if hash != expectedHash {
					return nil, errors.New("oh no")
				}
				return &model.EmailVerification{Hash: hash}, nil
			}
			s := &Service{db: fk}

			_, err := s.Verify(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	correctEmail := "demo@demo.com"
	correctPassword := "demoPassword123!"
	correctSalt := "$2a$10$T/26fFbPqC9GY/zsQgGuGO1djroBCIXbL1kRXQpDw.OlKPniDTQt2---"

	type args struct {
		ctx context.Context
		req *request.Login
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get success with correct email and password",
			args{nil, &request.Login{Email: correctEmail, Password: correctPassword}},
			false,
		},
		{
			"get unauthorized with wrong email and password",
			args{nil, &request.Login{Email: "random@email.com", Password: "supersecurepassword"}},
			false,
		},
		{
			"get error with empty fields",
			args{nil, &request.Login{Email: "", Password: ""}},
			false,
		},
	}
	fk := &mocks.FakeDBClient{}
	fk.GetUserStub = func(email string) (*model.User, error) {
		if email != correctEmail {
			return &model.User{Email: email, Name: "", Salt: "", Verified: false}, nil
		}
		return &model.User{Email: correctEmail, Name: "", Salt: correctSalt, Verified: true}, nil
	}
	s := &Service{db: fk}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
