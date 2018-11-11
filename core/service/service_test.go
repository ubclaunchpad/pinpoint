package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ubclaunchpad/pinpoint/core/database/mocks"
	"github.com/ubclaunchpad/pinpoint/core/model"

	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestService_New(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	aws, _ := utils.AWSSession(utils.AWSConfig(true))
	if _, err = New(aws, l, Opts{}); err != nil {
		t.Error(err)
		return
	}

	// with TLS
	if _, err = New(aws, l, Opts{
		TLSOpts: TLSOpts{
			CertFile: "../../dev/certs/127.0.0.1.crt",
			KeyFile:  "../../dev/certs/127.0.0.1.key",
		},
	}); err != nil {
		t.Error(err)
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
		wantErr bool
	}{
		{
			"get password requirement error",
			args{&request.CreateAccount{
				Email:    "test@pinpoint.com",
				Name:     "test",
				Password: "1234"}},
			true,
		},
		{
			"get email requirement error",
			args{&request.CreateAccount{
				Email:    "test",
				Name:     "test",
				Password: "1234pass."}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &mocks.FakeDBClient{}
			s := &Service{db: fk}
			_, err := s.CreateAccount(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
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
