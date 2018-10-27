package model

// Club info
type Club struct {
	ID          string `json:"pk"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ClubTable   string `json:"ct"`
}

// User info
type User struct {
	Email string `json:"pk"`
	Name  string `json:"name"`
	Salt  string `json:"salt"`
}

// ClubUser manages the relationship of a club to a user
type ClubUser struct {
	ClubID   string `json:"pk"`
	Email    string `json:"sk"`
	UserName string `json:"name"`
	Role     string `json:"role"`
}
