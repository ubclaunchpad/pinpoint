package poc

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
)

// WriteCSV takes a relative path and an array of string and appends to file
func WriteCSV(relativePath string, line []string) error {
	path, err := filepath.Abs(relativePath)
	csvfile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Write(line)
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

// FindHash looks up the given hash and tries to return a match
// Note: No error if not found, just empty result
func FindHash(hash string) ([]string, error) {
	path, err := filepath.Abs("./tmp/to_be_verified_accounts.csv")
	csvfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var matchingLine []string
	reader := csv.NewReader(csvfile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if line[1] == hash {
			matchingLine = line
		}
	}

	return matchingLine, nil
}
