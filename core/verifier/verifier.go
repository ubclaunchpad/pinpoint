package verifier

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

type verify struct {
	email, hash, createdAt string
}

// Init sets up verification on passed email address
func Init(email string) (string, error) {
	hash, err := generateHash(email)
	if err != nil {
		return "", err
	}

	err = writeCSV(email, hash, time.Now().Unix())
	if err != nil {
		return "", err
	}

	return hash, nil
}

// Verify looks up given hash and sets verified to true for the matching email
func Verify(hash string) error {
	path, err := filepath.Abs("./tmp/to_be_verified_accounts.csv")
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
		path, err = filepath.Abs("./tmp/verified_accounts.csv")
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
func generateHash(email string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum([]byte{}))
	return sha, nil
}

// writeCSV writes the to-be verified email into a CSV
func writeCSV(email, hash string, createdAt int64) error {
	// TODO: Replace tmp files with database
	path, err := filepath.Abs("./tmp/to_be_verified_accounts.csv")
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
