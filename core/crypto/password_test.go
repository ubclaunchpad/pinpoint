package crypto

import (
	"testing"
)

func TestValidateCredentialValues(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantedErr error
	}{
		{"valid case", args{"robert", "pinpoint@LP2018"}, false, nil},
		{"valid case", args{"robert_lp", "pinpoint@LP$2018"}, false, nil},
		{"valid case", args{"robert-lp", "#pinpoint@LP$2018"}, false, nil},
		{"same user password", args{"mojave", "mojave"}, true, errPasswordContainsUsername},
		{"invalid username", args{"robert ", "pinpoint@LP2018"}, true, errInvalidUsername},
		{"invalid username", args{"robert@lp", "pinpoint@LP2018"}, true, errInvalidUsername},
		{"invalid username", args{"ROBERT.lp", "pinpoint@f$T#4%7"}, true, errInvalidUsername},
		{"invalid password", args{"robert", "pinpoint f$T@4%7"}, true, errInvalidPassword},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if err = ValidateCredentialValues([]string{tt.args.username}, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCredentialValues() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && tt.wantedErr != nil {
				if err != tt.wantedErr {
					t.Errorf("wanted %s, got %s", tt.wantedErr, err)
				}
			}
		})
	}
}

func Test_hashAndSalt(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantedErr error
	}{
		{"valid case", args{"Hxt-f$T@4%7"}, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashAndSalt(tt.args.password)
			if hash == tt.args.password {
				t.Errorf("HashAndSalt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("HashAndSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != tt.wantedErr {
				t.Errorf("HashAndSalt() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func Test_comparePasswords(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{"valid case", args{"Hxt-f$T@4%7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := HashAndSalt(tt.args.password)
			diffHash, err2 := HashAndSalt("Hxt-f$T%7")
			if result := ComparePasswords(hashed, tt.args.password); result != true {
				t.Errorf("ComparePasswords() = %v, error %v", result, err)
			}
			if result := ComparePasswords(diffHash, tt.args.password); result != false {
				t.Errorf("ComparePasswords() = %v, error %v", result, err2)
			}
		})
	}
}
