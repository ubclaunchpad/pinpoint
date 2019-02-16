package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Constant error messages
const (
	InvalidTokenErr = "invalid token"
)

// Claims (payload) of JWT token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken creates a JWT token using signKey (ie private rsa key)
func (c *Claims) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signKey, err := GetAPIPrivateKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(signKey)
}

// ValidateToken ensures token is valid and returns its metadata
func ValidateToken(tokenString string, lookup jwt.Keyfunc) (*Claims, error) {
	// Parse takes the token string and a function for looking up the key.
	// For default lookup function, use GetAPIPrivateKey()
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
