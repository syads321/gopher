package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
)

func (r *UserResolver) UserID() graphql.ID {
	return r.u.UserID
}
