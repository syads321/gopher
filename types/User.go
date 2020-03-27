package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// User jdjdjdj
type User struct {
	UserID    graphql.ID
	Username  string
	Emoji     string
	Notes     []*Note
	Password  string
	Email     string `validate:"required,email"`
	Token     string
	ExpiresAt string
}
