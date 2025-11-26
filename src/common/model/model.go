package model

type User struct {
	Username string
	Role     string // "admin", "user", "guest"
	Active   bool
}
