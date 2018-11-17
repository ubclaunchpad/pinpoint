package crypto

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	errPasswordContainsUsername = errors.New("password must not contain username")
	errInvalidEmail             = errors.New("email does not appear to be a valid email")
	errInvalidPassword          = errors.New("password must be at least 5 characters. only alphanumeric and symbols are alowed")
)

// HashAndSalt hashes and salts the given user password
func HashAndSalt(password string) (string, error) {
	bytePwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("password hashing unsuccessful: " + err.Error())
	}
	return string(hash), nil
}

// ComparePasswords checks if given password maps correctly to the given hash
func ComparePasswords(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePwd := []byte(password)
	return bcrypt.CompareHashAndPassword(byteHash, bytePwd) == nil
}

// ValidateCredentialValues verifies that the chosen username and passwords are valid
func ValidateCredentialValues(email, password string) error {
	if strings.Contains(email, password) {
		return errPasswordContainsUsername
	}
	if len(email) < 3 || len(email) >= 128 || !isEmailFormat(email) {
		return errInvalidEmail
	}

	if len(password) < 5 || len(password) >= 128 || !isLegalPassword(password) {
		return errInvalidPassword
	}

	return nil
}

// isLegalPassword returns true if the chosen password does not contain illegal characters
// Only alphanumeric characters and symbols are alowed. These correspond to 33-126 range in ASCII table
func isLegalPassword(password string) bool {
	for _, c := range password {
		if c < 33 || c > 126 {
			return false
		}
	}
	return true
}

// isEmailFormat checks that a given email is formatted correctly
func isEmailFormat(email string) bool {
	regex := regexp.MustCompile(`^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	return regex.MatchString(email)
}
