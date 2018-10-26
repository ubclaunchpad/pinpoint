package mailer

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"io"
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

type verify struct {
	email, hash, createdAt string
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

	err = writeCSV(v.email, hash, v.createdAt.Unix())
	if err != nil {
		return "", err
	}

	return hash, nil
}

// Verify looks up given hash and sets verified to true for the matching email
func Verify(hash string) error {
	path, err := filepath.Abs("./tmp/to_be_verifies.csv")
	csvfile, err := os.Open(path)
	if err != nil {
		return err
	}
	var toBeVerified verify
	reader := csv.NewReader(csvfile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if line[1] == hash {
			toBeVerified = verify{line[0], line[1], line[2]}
		}
	}
	csvfile.Close()

	if toBeVerified != (verify{}) {
		path, err = filepath.Abs("./tmp/verifies.csv")
		csvfile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
		defer csvfile.Close()
		writer := csv.NewWriter(csvfile)
		writer.Write([]string{toBeVerified.email, toBeVerified.hash, strconv.FormatInt(time.Now().Unix(), 10)})
		writer.Flush()
		if err := writer.Error(); err != nil {
			return err
		}
		return nil
	}

	return errors.New("unable to find matching email")
}

// generateHash generates a hash used for verifying
func (v *Verifier) generateHash() (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(v.email))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
	return sha, nil
}

func writeCSV(email, hash string, createdAt int64) error {
	path, err := filepath.Abs("./tmp/to_be_verifies.csv")
	csvfile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(csvfile)
	defer csvfile.Close()
	writer.Write([]string{email, hash, strconv.FormatInt(createdAt, 10)})
	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
