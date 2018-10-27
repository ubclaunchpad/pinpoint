package mailer

import (
	"reflect"
	"testing"
)

func TestMailer_New(t *testing.T) {
	type args struct {
		from, pass string
	}
	tests := []struct {
		name    string
		args    args
		want    *Mailer
		wantErr bool
	}{
		{
			"get success",
			args{"test@pinpoint.com", "doesn't check password"},
			&Mailer{"test@pinpoint.com", "doesn't check password"},
			false,
		},
		{
			"get error",
			args{"not a real email", "eh"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.from, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() error = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMailer_emailFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"get true",
			args{"test@pinpoint.com"},
			true,
		},
		{
			"get false",
			args{"testpinpoint.com"},
			false,
		},
		{
			"get false",
			args{"@pinpoint.com"},
			false,
		},
		{
			"get false",
			args{"test"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := emailFormat(tt.args.email)
			if got != tt.want {
				t.Errorf("emailFormat() %v, want %v", got, tt.want)
			}
		})
	}
}
