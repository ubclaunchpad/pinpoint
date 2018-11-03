package api

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	gateutil "github.com/ubclaunchpad/pinpoint/gateway/utils"
	"github.com/ubclaunchpad/pinpoint/protobuf/mocks"
	"github.com/ubclaunchpad/pinpoint/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// NewMockAPI is used to create an API with a mocked client for use in tests
func NewMockAPI(l *zap.SugaredLogger, t *testing.T) (*API, *mocks.MockCoreClient, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockCoreClient(ctrl)
	a, err := New(l, CoreOpts{})
	if err != nil {
		t.Fatal(err)
	}
	a.c = mock
	return a, mock, ctrl
}

func TestNew(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		opts CoreOpts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"without opts", args{CoreOpts{}}, false},
		{"with opts", args{CoreOpts{
			Host:     "localhost",
			Port:     "9111",
			CertFile: "../../dev/certs/127.0.0.1.crt",
		}}, false},
		{"with invalid opts", args{CoreOpts{
			Host:     "localhost",
			Port:     "9111",
			CertFile: "../../README.md",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(l, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Run(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	type args struct {
		host string
		opts RunOpts
	}
	tests := []struct {
		name       string
		args       args
		clientFail bool
	}{
		{"with client failure", args{"localhost", RunOpts{}}, true},
		{"with invalid host", args{"", RunOpts{}}, true},
		{"no options", args{"localhost", RunOpts{}}, false},
		{"with gateway TLS", args{"localhost", RunOpts{
			CertFile: "../../dev/certs/127.0.0.1.crt",
			KeyFile:  "../../dev/certs/127.0.0.1.key",
		}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set up mock controller
			api, mock, ctrl := NewMockAPI(l, t)
			defer ctrl.Finish()

			if tt.args.host != "" {
				if tt.clientFail {
					// set client to fail
					mock.EXPECT().
						GetStatus(gomock.Any(), gomock.Any(), gomock.Any()).
						DoAndReturn(func(...interface{}) (interface{}, error) {
							return nil, errors.New("oh no")
						}).
						Times(1)
					mock.EXPECT().
						Handshake(gomock.Any(), gomock.Any(), gomock.Any()).
						DoAndReturn(func(...interface{}) (interface{}, error) {
							return nil, errors.New("oh no")
						}).
						Times(1)
				} else {
					// expect exactly one call to GetStatus && Handshake with anything
					mock.EXPECT().
						GetStatus(gomock.Any(), gomock.Any(), gomock.Any()).
						Times(1)
					mock.EXPECT().
						Handshake(gomock.Any(), gomock.Any(), gomock.Any()).
						Times(1)
				}
			}

			// run the server!
			go api.Run(tt.args.host, "", tt.args.opts)
			time.Sleep(time.Millisecond)
			api.Stop()
		})
	}
}

//Check if gateway is properly adding token to context
func TestSecureContext(t *testing.T) {
	type args struct {
		host       string
		context    context.Context
		setcontext bool
		opts       RunOpts
	}
	tests := []struct {
		name       string
		args       args
		clientFail bool
	}{
		{"Secure Context", args{"localhost",
			context.Background(), true,
			RunOpts{}},
			false},
		{"Original Context", args{"localhost",
			context.Background(), false,
			RunOpts{}},
			true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assign token values
			os.Setenv("PINPOINT_CORE_TOKEN", "valid_token")
			os.Setenv("PINPOINT_GATEWAY_TOKEN", "valid_token")
			ctx := tt.args.context
			if tt.args.setcontext {
				ctx = gateutil.SecureContext(ctx)
			}
			meta, ok := metadata.FromOutgoingContext(ctx)
			if !ok && !tt.clientFail {
				t.Errorf("missing context metadata")
				return
			}
			if len(meta["token"]) != 1 && !tt.clientFail {
				t.Errorf("invalid token")
				return
			}
		})
	}
}
