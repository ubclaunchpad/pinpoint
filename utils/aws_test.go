package utils

import (
	"testing"
)

func TestAWSConfig(t *testing.T) {
	l, _ := NewLogger(true, "")
	type args struct {
		dev    bool
		logger Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{"dev", args{true, nil}},
		{"prod", args{false, nil}},
		{"with logger", args{false, l}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
