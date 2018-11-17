package utils

import (
	"reflect"
	"testing"
)

func TestFirstString(t *testing.T) {
	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"nil", args{nil}, ""},
		{"1 string", args{[]interface{}{"hello"}}, "hello"},
		{"2 strings", args{[]interface{}{"hello", "world"}}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstString(tt.args.strs); got != tt.want {
				t.Errorf("FirstString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{"nil args", args{nil}, nil},
		{"no args", args{[]interface{}{}}, nil},
		{"odd args", args{[]interface{}{"hello"}}, nil},
		{"pair of args", args{[]interface{}{"hello", "world"}}, map[string]interface{}{
			"hello": "world",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
