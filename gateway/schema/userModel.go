package schema

type User struct {
	Name     string
	Email    string
	Password string
}

// NewUser creates a new user
func NewUser(n, e, p string) *User {
	return &User{Name: n, Email: e, Password: p}
}
