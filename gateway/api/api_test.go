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
	api, err := New(l)
	if err != nil {
		t.Error(err)
		return
	}

	// stub run
	go api.Run("localhost", "8080", RunOpts{})
	time.Sleep(time.Millisecond)
	api.Stop()

	// stub run with core tls
	go api.Run("localhost", "8080", RunOpts{
		CoreOpts: CoreOpts{
			Host:     "localhost",
			Port:     "9111",
			CertFile: "../../dev/certs/127.0.0.1.crt",
		},
	})
	time.Sleep(time.Millisecond)
	api.Stop()

	// stub run with gateway tls
	go api.Run("localhost", "8081", RunOpts{
		GatewayOpts: GatewayOpts{
			CertFile: "../../dev/certs/127.0.0.1.crt",
			KeyFile:  "../../dev/certs/127.0.0.1.key",
		},
	})
	time.Sleep(time.Millisecond)
	api.Stop()
}
