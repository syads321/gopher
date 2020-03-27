package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type CreateNoteArgs struct {
	UserID graphql.ID
	Note   NoteInput
}
