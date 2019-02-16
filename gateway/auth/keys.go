package auth

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	privateKeyPath = "dev/keys/app.rsa"
	publicKeyPath  = "dev/keys/app.rsa.pub"
)

// GetAPIPrivateKey returns the private RSA key to authenticate
// HTTP requests sent to the daemon.
func GetAPIPrivateKey() ([]byte, error) {
	return getPrivateKeyFromPath(getProjectRelativePath(privateKeyPath))
}

func getPrivateKeyFromPath(path string) ([]byte, error) {
	signBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return signBytes, nil
}

// GetAPIPublicKey returns public RSA key
func GetAPIPublicKey() ([]byte, error) {
	return getPublicKeyFromPath(getProjectRelativePath(publicKeyPath))
}

func getPublicKeyFromPath(path string) ([]byte, error) {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return verifyBytes, nil
}

func getProjectRelativePath(path string) string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "pinpoint") {
		wd = filepath.Dir(wd)
	}

	return wd + "/" + path
}
