package model

import "time"

// Club info
type Club struct {
	ID          string `json:"pk"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// User info
type User struct {
	Email    string `json:"pk"`
	Name     string `json:"name"`
	Salt     string `json:"salt"`
	Verified bool   `json:"verified"`
}

// ClubUser manages the relationship of a club to a user
type ClubUser struct {
	ClubID   string `json:"pk"`
	Email    string `json:"sk"`
	UserName string `json:"name"`
	Role     string `json:"role"`
}

// EmailVerification denotes a pending email verification
type EmailVerification struct {
	Hash   string    `json:"pk"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}
