package crypto

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	errSameUsernamePassword = errors.New("Username and password must be different")
	errInvalidUsername      = errors.New("Username must be at least 3 characters. Only alphanumeric characters, underscores, and dashes are allowed")
	errInvalidPassword      = errors.New("Password must be at least 5 characters. Only alphanumeric characters, and symbols are alowed")
)

// hashAndSalt hashes and salts the given user password 
func hashAndSalt(password string) (string, error) {

	bytePwd := []byte(password)
    hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	
	if err != nil {
		return "", errors.New("Password hashing unsuccessful: " + err.Error())
    }
    return string(hash), nil
}

// comparePasswords checks if given password maps correctly to the given hash
func comparePasswords(hashedPassword string, password string) bool {
	
	byteHash := []byte(hashedPassword)
	bytePwd := []byte(password)
	
	return bcrypt.CompareHashAndPassword(byteHash, bytePwd) == nil
}

// ValidateCredentialValues verifies that the chosen username and passwords are valid
// A valid password must be at least 5 characters long
// A valid username must be at least 3 characters and contains only legal characters
func ValidateCredentialValues(username string, password string) error {
	if username == password {
		return errSameUsernamePassword
	}
	if len(password) < 5 || len(password) >= 128 || !IsLegalPassword(password) {
		return errInvalidPassword
	}
	if len(username) < 3 || len(username) >= 128 || !IsLegalUserName(username) {
		return errInvalidUsername
	}
	return nil
}

// IsLegalUserName returns true if the chosen username only contains characters [A-Z], [a-z], or '_' or '-'
func IsLegalUserName(username string) bool {
	for _, c := range username {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < 48 || c > 57) && c != '_' && c != '-' {
			return false
		}
	}
	return true
}

// IsLegalPassword returns true if the chosen password does not contain illegal characters
// Only alphanumeric characters and symbols are alowed. These correspond to 33-126 range in ASCII table
func IsLegalPassword(password string) bool {
	for _, c := range password {
		if (c < 33 || c > 126) {
			return false
		}
	}
	return true
}
