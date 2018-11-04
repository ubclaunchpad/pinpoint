package schema

// CreateUser defines a request to create a new user
type CreateUser struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"confirmPassword"`
	ESub      bool   `json:"emailSubscribe"`
}
