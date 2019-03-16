package utils

import (
	"os"
	"testing"

	"go.uber.org/zap/zaptest"
)

func TestAWSConfig(t *testing.T) {
	var l = zaptest.NewLogger(t).Sugar()
	type args struct {
		dev    bool
		logger Logger
		debug  bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"dev", args{true, nil, false}},
		{"prod", args{false, nil, false}},
		{"with logger", args{false, l, false}},
		{"with debug", args{false, nil, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.debug {
				os.Setenv(AWSDebugEnv, "true")
				defer os.Setenv(AWSDebugEnv, "")
			}

			// skip any checks for now
			_ = AWSConfig(tt.args.dev, tt.args.logger)
		})
	}
}

func TestAWSSession(t *testing.T) {
	if _, err := AWSSession(); err != nil {
		t.Error(err)
	}
}
