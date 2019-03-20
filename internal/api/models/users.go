package models

import "time"

// User is a user object
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token is a token object
type Token struct {
	Name    string
	Token   string
	Expires time.Time
}
