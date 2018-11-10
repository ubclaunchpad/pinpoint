package verifier

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/ubclaunchpad/pinpoint/core/mailer"
)

// Init sets up verification on passed email address
func Init(email string, m *mailer.Mailer) (string, error) {
	if email == "" {
		return "", errors.New("invalid email")
	}

	hash := generateHash(email)

	// Send verification email - TODO: Change to get email address from user session
	if m != nil {
		if err := m.Send(email,
			"Welcome to Pinpoint!",
			fmt.Sprintf("Visit localhost:8081/user/verify?hash=%s to verify your email.", hash)); err != nil {
			return hash, fmt.Errorf("failed to send email: %s", err.Error())
		}
	}

	return hash, nil
}

// generateHash generates a hash used for verifying
func generateHash(email string) string {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	return base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
}
