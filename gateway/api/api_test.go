package api

import (
	"context"
	"os"
	"testing"
	"time"

	gateutil "github.com/ubclaunchpad/pinpoint/gateway/utils"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

func TestAPI_establishConnection(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		opts RunOpts
	}

	testopts := args{RunOpts{
		GatewayOpts: GatewayOpts{
			CertFile: "../../dev/certs/127.0.0.1.crt",
			KeyFile:  "../../dev/certs/127.0.0.1.key",
		},
	}}

	md := metadata.Pairs("token", "WRONG_TOKEN")
	wrongctx := metadata.NewOutgoingContext(context.Background(), md)

	tests := []struct {
		name          string
		ctx           context.Context
		args          args
		errorexpected bool
	}{
		{"Blank Context", context.Background(), testopts, true},
		{"Incorrect Token Used", wrongctx, testopts, true},
		{"Token Inserted", gateutil.SecureContext(context.Background()), testopts, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			os.Setenv("PINPOINT_CORE_TOKEN", "valid_token")
			os.Setenv("PINPOINT_GATEWAY_TOKEN", "valid_token")

			a, err := New(l)
			if err != nil {
				t.Error(err)
				return
			}

			// opts := tt.args.opts

			// set up server
			a.srv.Addr = "localhost" + ":" + "9111"
			// set up parameters
			dialOpts := make([]grpc.DialOption, 0)

			creds, err := credentials.NewClientTLSFromFile("../../dev/certs/127.0.0.1.crt", "")
			if err != nil {
				t.Error(err)
				return
			}
			dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))

			// connect to core service
			a.l.Infow("connecting to core",
				"core.host", "localhost",
				"core.port", "9111",
				"core.tls", "../../dev/certs/127.0.0.1.crt" != "")
			conn, err := grpc.Dial("localhost"+":"+"9111", dialOpts...)
			if err != nil {
				t.Error(err)
				return
			}

			a.c = pinpoint.NewCoreClient(conn)
			defer conn.Close()

			// Exchange auth tokens with core
			if err := a.establishConnection(tt.ctx); err != nil {
				a.l.Infow("Closing connection")
				conn.Close()

				// Should not have gotten error
				if !tt.errorexpected {
					t.Error(err)
				}
				return
			}

			// Should have gotten error
			if tt.errorexpected {
				t.Error(err)
			}
			return
		})
	}
}
