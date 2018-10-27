package verifier

import (
	"crypto/sha1"
	"encoding/base64"
)

// Init sets up verification on passed email address
func Init(email string) (string, error) {
	hash, err := generateHash(email)
	if err != nil {
		return "", err
	}

	// TODO: Add a database write

	return hash, nil
}

// generateHash generates a hash used for verifying
func generateHash(email string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
	return sha, nil
}
