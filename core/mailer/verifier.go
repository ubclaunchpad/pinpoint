package mailer

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Verifier contains context required verify email
type Verifier struct {
	email     string
	createdAt time.Time
}

// NewVerifier returns a new Verifier using parameter email
func NewVerifier(email string) *Verifier {
	createdAt := time.Now()
	return &Verifier{email, createdAt}
}

// Init sets up verification on the passed email address
func (v *Verifier) Init() (string, error) {
	hash, err := v.generateHash()
	if err != nil {
		return "", err
	}

	err = writeCSV(v.email, hash, false, v.createdAt.Unix())
	if err != nil {
		return "", err
	}

	return hash, nil
}

// generateHash generates a hash used for verifying
func (v *Verifier) generateHash() (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(v.email))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
	return sha, nil
}

func writeCSV(email, hash string, verified bool, createdAt int64) error {
	path, err := filepath.Abs("./tmp/verifies.csv")
	csvfile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(csvfile)
	defer csvfile.Close()
	writer.Write([]string{email, hash, strconv.FormatBool(verified), strconv.FormatInt(createdAt, 10)})
	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
