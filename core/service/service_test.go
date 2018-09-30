package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/ubclaunchpad/pinpoint/grpc/request"
	"github.com/ubclaunchpad/pinpoint/grpc/response"
)

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
