package resolver

import (
	"context"
	ctx "eaciit/gopher/dbcontext"
	helper "eaciit/gopher/helper"
	types "eaciit/gopher/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
)

// CreateUser Create User resolver
func (r *RootResolver) RegisterUser(args types.RegisterUserArgs) (*UserResolver, error) {
	// Find user:
	var user *types.User
	validate := validator.New()
	user = &types.User{
		Password: args.Password,
		Email:    args.Email,
	}

	errs := validate.Struct(user)
	if errs != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: errs.Error()}
	}

	encrytPass, err := helper.Encrypt(args.Password)
	user.Password = encrytPass
	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Encryption Error"}
	}
	opts := options.Find().SetSort(bson.D{{"Email", 1}})
	cursor, err := ctx.MainContext("User").Find(context.TODO(), bson.D{{"Email", args.Email}}, opts)
	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Error when query"}
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: errs.Error()}
	}

	if len(results) > 0 {
		return &UserResolver{user}, &types.CommonError{Code: "Existed", Message: "email already existed"}
	}

	_, err = ctx.MainContext("User").InsertOne(context.TODO(),
		bson.D{
			bson.E{Key: "Email", Value: user.Email},
			bson.E{Key: "Password", Value: user.Password},
		},
	)
	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Error inserting"}
	}

	// Didn’t find user:
	return &UserResolver{user}, nil
}
