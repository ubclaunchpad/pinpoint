package api

import (
	"testing"
	"time"

	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestAPI_New(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
	}
	_, _ = New(l)
}

func TestAPI_Run(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		opts RunOpts
	}
	tests := []struct {
		name string
		args args
	}{
		{"no options", args{RunOpts{}}},
		{"with core TLS", args{RunOpts{
			CoreOpts: CoreOpts{
				Host:     "localhost",
				Port:     "9111",
				CertFile: "../../dev/certs/127.0.0.1.crt",
			},
		}}},
		{"with gateway TLS", args{RunOpts{
			GatewayOpts: GatewayOpts{
				CertFile: "../../dev/certs/127.0.0.1.crt",
				KeyFile:  "../../dev/certs/127.0.0.1.key",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := New(l)
			if err != nil {
				t.Error(err)
				return
			}
			go api.Run("localhost", "8080", tt.args.opts)
			time.Sleep(time.Millisecond)
			api.Stop()
		})
	}
}
