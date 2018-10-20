package model

// Club info
type Club struct {
	ID          string   `json:"pk"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Periods     []string `json:"periods"`
}

// User info
type User struct {
	Email string `json:"pk"`
	Name  string `json:"name"`
	Salt  string `json:"salt"`
}

// Clubuser manages the relationship of a club to a user
type Clubuser struct {
	ClubID   string `json:"pk"`
	Email    string `json:"sk"`
	UserName string `json:"name"`
	Role     string `json:"role"`
}
