package utils

import (
	"errors"
	"reflect"
	"testing"
)

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
		{"non-string key", args{[]interface{}{errors.New("asdf"), "hello"}}, nil},
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
