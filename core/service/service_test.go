package service

import (
	"context"
	"reflect"
	"testing"
	"time"

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
	if _, err = New(aws, l); err != nil {
		t.Error(err)
		return
	}
}

func TestService_Run(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	aws, _ := utils.AWSSession(utils.AWSConfig(true))
	s, err := New(aws, l)
	if err != nil {
		t.Error(err)
		return
	}

	// stub run
	go s.Run("", "", RunOpts{})
	time.Sleep(time.Millisecond)
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
			args{nil, &request.Status{Callback: "hi"}},
			&response.Status{Callback: "hi"},
			false,
		},
		{
			"get error",
			args{nil, &request.Status{Callback: "I don't like launch pad"}},
			&response.Status{Callback: "I don't like launch pad"},
			true,
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
