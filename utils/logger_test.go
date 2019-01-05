package utils

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		dev  bool
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"dev", args{true, ""}, false},
		{"prod", args{false, ""}, false},
		{"prod with logpath", args{false, "tmp/logs"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSugar, err := NewLogger(tt.args.dev, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotSugar == nil {
				t.Error("got unexpected nil logger")
			}
		})
	}
}
