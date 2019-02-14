package auth

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "dev/keys/app.rsa"
	publicKeyPath  = "dev/keys/app.rsa.pub"
)

// GetAPIPrivateKey returns the private RSA key to authenticate
// HTTP requests sent to the daemon.
func GetAPIPrivateKey() (*rsa.PrivateKey, error) {
	return getPrivateKeyFromPath(privateKeyPath)
}

func getPrivateKeyFromPath(path string) (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(signBytes)
}

// GetAPIPublicKey returns public RSA key
func GetAPIPublicKey() (*rsa.PublicKey, error) {
	return getPublicKeyFromPath(publicKeyPath)
}

func getPublicKeyFromPath(path string) (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
}
