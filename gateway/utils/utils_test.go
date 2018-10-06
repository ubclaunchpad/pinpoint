package utils

import "testing"

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
