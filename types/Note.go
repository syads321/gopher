package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Note sdjdj
type Note struct {
	UserID graphql.ID
	Data   string
}
