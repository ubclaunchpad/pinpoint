package verifier

import (
	"testing"

	"github.com/ubclaunchpad/pinpoint/core/mailer"
)

func TestInit(t *testing.T) {
	type args struct {
		email string
		m     *mailer.Mailer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid email", args{"", nil}, true},
		{"valid email but no mailer", args{"test@pinpoint.com", nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = New(tt.args.email, tt.args.m)
		})
	}
}

func TestVerifier_generateHash(t *testing.T) {
	hash := generateHash("test@pinpoint.com")
	if hash != "NmSdjumzjHOF7IAnafAK74LAPug=" {
		t.Error("Unexpected hash")
		return
	}
}
