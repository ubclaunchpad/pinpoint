package api

import (
	"errors"
	"testing"
	"time"

	"github.com/ubclaunchpad/pinpoint/protobuf/fakes"
	"go.uber.org/zap/zaptest"
)

// newMockAPI is used to create an API with a mocked client for use in tests
func newMockAPI(t *testing.T) (*API, *fakes.FakeCoreClient) {
	fake := &fakes.FakeCoreClient{}
	a, err := New(zaptest.NewLogger(t).Sugar(), CoreOpts{})
	if err != nil {
		t.Fatal(err)
	}
	a.c = fake
	return a, fake
}

func TestNew(t *testing.T) {
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
			var l = zaptest.NewLogger(t).Sugar()
			if _, err := New(l, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAPI_Run(t *testing.T) {
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
			api, fake := newMockAPI(t)

			if tt.clientFail {
				// set client to fail
				fake.GetStatusReturns(nil, errors.New("oh no"))
				fake.HandshakeReturns(nil, errors.New("oh no"))
			}

			// run the server!
			go api.Run(tt.args.host, "", tt.args.opts)
			time.Sleep(time.Millisecond)
			api.Stop()
		})
	}
}
