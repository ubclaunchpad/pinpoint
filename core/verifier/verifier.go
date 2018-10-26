package verifier

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"github.com/ubclaunchpad/pinpoint/core/poc"
)

type verify struct {
	email, hash, createdAt string
}

// Init sets up verification on passed email address
func Init(email string) (string, error) {
	hash, err := generateHash(email)
	if err != nil {
		return "", err
	}

	// TODO: Replace with database write
	line := []string{email, hash, strconv.FormatInt(time.Now().Unix(), 10)}
	if err = poc.WriteCSV("./tmp/to_be_verified_accounts.csv", line); err != nil {
		return "", err
	}

	return hash, nil
}

// Verify looks up given hash and sets verified to true for the matching email
func Verify(hash string) error {
	// TODO: Replace with database lookup
	match, err := poc.FindHash(hash)
	if err != nil {
		return err
	}

	if len(match) > 0 {
		// TODO: Replace with database write
		toBeVerified := verify{match[0], match[1], match[2]}
		line := []string{toBeVerified.email, toBeVerified.hash, strconv.FormatInt(time.Now().Unix(), 10)}
		if err = poc.WriteCSV("./tmp/verified_accounts.csv", line); err != nil {
			return err
		}
		return nil
	}

	return errors.New("unable to find matching email")
}

// generateHash generates a hash used for verifying
func generateHash(email string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
	return sha, nil
}
