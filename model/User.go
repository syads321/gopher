package model

import (
	ctx "eaciit/gopher/dbcontext"
	"go.mongodb.org/mongo-driver/mongo"
)

// User model
type User struct {
}

// Collection of user raw collection
func (r *User) Collection() *mongo.Collection {
	return ctx.MainContext("User")
}
