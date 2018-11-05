package utils

import (
	"context"
	"os"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestFirstString(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"nil", args{nil}, ""},
		{"1 string", args{[]string{"hello"}}, "hello"},
		{"2 strings", args{[]string{"hello", "world"}}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstString(tt.args.strs); got != tt.want {
				t.Errorf("FirstString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecureContext(t *testing.T) {
	type args struct {
		host       string
		context    context.Context
		setcontext bool
	}
	tests := []struct {
		name       string
		args       args
		clientFail bool
	}{
		{"Secure Context", args{"localhost",
			context.Background(), true},
			false},
		{"Original Context", args{"localhost",
			context.Background(), false},
			true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assign token values
			os.Setenv("PINPOINT_CORE_TOKEN", "valid_token")
			os.Setenv("PINPOINT_GATEWAY_TOKEN", "valid_token")
			ctx := tt.args.context
			if tt.args.setcontext {
				ctx = SecureContext(ctx)
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
