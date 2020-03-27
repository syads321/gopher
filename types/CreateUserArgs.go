package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type CreateUserArgs struct {
	UserID   graphql.ID
	Username string
	Emoji    string
	Password string
	Notes    *[]NoteInput
	Email    string
}
