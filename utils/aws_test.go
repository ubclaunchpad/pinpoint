package utils

import (
	"testing"
)

func TestAWSConfig(t *testing.T) {
	type args struct {
		dev bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"dev", args{true}},
		{"prod", args{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// skip any checks for now
			_ = AWSConfig(tt.args.dev)
		})
	}
}

func TestAWSSession(t *testing.T) {
	if _, err := AWSSession(); err != nil {
		t.Error(err)
	}
}
