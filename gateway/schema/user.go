package schema

// CreateUser defines a request to create a new user
type CreateUser struct {
	Name     string
	Email    string
	Password string
}
