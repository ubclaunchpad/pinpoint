package auth

import (
	"io/ioutil"
	"crypto/rsa"
	"log"
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// the public key is used to verify the tokens 
// the private key is used for signing the tokens
var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
)

// @robert: how should these key be stored ? not sure if this is right!
const (									    // @robert: here are commands I used to generate keys
	privateKeyPath = "auth/keys/app.rsa"    // openssl genrsa -out app.rsa 2048
	publicKeyPath = "auth/keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
	InvalidTokenErr = "invalid token"
)

// Claims (payload) of JWT token
type Claims struct {
	// @robert: not sure what the claims should include ? 
	// using the jwt libaray's standard claims for now
	jwt.StandardClaims
}

// handleErr a helper function to handle the errors
func handleErr(err error) {
	if (err != nil){
		log.Fatal(err)}
	}

// initializeKeys reads the key files
func initializeKeys() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	handleErr(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	handleErr(err)

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	handleErr(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	handleErr(err)
	}

// GenerateToken creates a JWT token using signKey (ie private rsa key)
func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	// just for this poc will only set the expiry and issued at times
    token.Claims = jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 1).Unix(),
        "iat": time.Now().Unix(),
	}
	return (token.SignedString(signKey))
}


// ValidateToken ensures token is valid and returns its metadata
// modified copy from inertia lol
// I am not really sure how the ParseWithClaims work, 
// particulary what are we passing "&Claims{}""
func ValidateToken(tokenString string, lookup jwt.Keyfunc) (*Claims, error) {
	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, lookup)
	if err != nil {
		return nil, err
	}
	
	// Verify signing algorithm and token
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || !token.Valid {
		return nil, errors.New(InvalidTokenErr)
	}

	// Verify the claims and token.
	if claim, ok := token.Claims.(*Claims); ok {
		return claim, nil
	}
	return nil, errors.New(InvalidTokenErr)
}
