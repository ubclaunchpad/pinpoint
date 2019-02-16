package auth

import (
	"io/ioutil"
)

const (
	privateKeyPath = "dev/keys/app.rsa"
	publicKeyPath  = "dev/keys/app.rsa.pub"
)

// GetAPIPrivateKey returns the private RSA key to authenticate
// HTTP requests sent to the daemon.
func GetAPIPrivateKey() ([]byte, error) {
	return getPrivateKeyFromPath(privateKeyPath)
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
	return getPublicKeyFromPath(publicKeyPath)
}

func getPublicKeyFromPath(path string) ([]byte, error) {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return verifyBytes, nil
}
