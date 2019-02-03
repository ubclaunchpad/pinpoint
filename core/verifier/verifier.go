package verifier

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ubclaunchpad/pinpoint/core/mailer"
)

// Verifier manages verification
type Verifier struct {
	Email  string
	Hash   string
	Expiry int64

	m *mailer.Mailer
}

// New sets up verification on passed email address
func New(email string, m *mailer.Mailer) Verifier {
	return Verifier{Email: email, Hash: generateHash(email), Expiry: int64(time.Now().Add(24 * 7 * time.Hour).Unix())}
}

// SendVerification sends a verification email
func (v *Verifier) SendVerification() error {
	if v.m == nil {
		return nil
	}
	return v.m.Send(v.Email,
		"Welcome to Pinpoint!",
		fmt.Sprintf("Visit localhost:8081/user/verify?hash=%s to verify your email.", v.Hash))
}

// generateHash generates a hash used for verifying
func generateHash(email string) string {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	return base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
}
