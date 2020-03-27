package resolver

import (
	"context"
	ctx "eaciit/gopher/dbcontext"
	helper "eaciit/gopher/helper"
	types "eaciit/gopher/types"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
)

// CreateUser Create User resolver
func (r *RootResolver) CreateUser(args types.CreateUserArgs) (*UserResolver, error) {
	// Find user:

	var user *types.User
	var notes []*types.Note
	var notesdocs = []interface{}{}
	validate := validator.New()
	user = &types.User{
		UserID:   args.UserID,
		Username: args.Username,
		Emoji:    args.Emoji,
		Password: args.Password,
		Email:    args.Email,
	}
	if r.Session == "" {
		return &UserResolver{user}, &types.CommonError{Code: "NotFound", Message: "Access Denied"}
	}
	errs := validate.Struct(user)
	if errs != nil {
		return &UserResolver{user}, &types.CommonError{Code: "NotFound", Message: errs.Error()}
	}

	if args.UserID == "err" {
		return &UserResolver{user}, &types.CommonError{Code: "NotFound", Message: "userID must valid"}
	}
	for _, t := range *args.Notes {
		notes = append(notes, &types.Note{
			Data:   t.Data,
			UserID: args.UserID,
		})

		notesdocs = append(notesdocs,
			bson.D{
				bson.E{Key: "Data", Value: t.Data},
				bson.E{Key: "UserID", Value: args.UserID},
			},
		)
	}

	_, err := ctx.MainContext("Notes").InsertMany(context.TODO(), notesdocs)
	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "NotFound", Message: "Error when insert notes "}
	}
	encrytPass, err := helper.Encrypt(args.Password)
	user.Password = encrytPass
	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "NotFound", Message: "Error when insert notes "}
	}
	opts := options.Find().SetSort(bson.D{{"UserName", 1}})
	cursor, err := ctx.MainContext("User").Find(context.TODO(), bson.D{{"Username", args.Username}}, opts)
	if err != nil {
		//log.Fatal(err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		//log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	if len(results) > 0 {
		return &UserResolver{user}, &types.CommonError{Code: "Existed", Message: "userID already existed"}
	}

	userResult, err := ctx.MainContext("User").InsertOne(context.TODO(),
		bson.D{
			bson.E{Key: "UserID", Value: user.UserID},
			bson.E{Key: "Username", Value: user.Username},
			bson.E{Key: "Emoji", Value: user.Emoji},
			bson.E{Key: "Password", Value: user.Password},
		},
	)
	if err == nil {
		fmt.Println("Inserted a single document: ", userResult.InsertedID)
	}

	users = append(users, user)
	// Didn’t find user:
	return &UserResolver{user}, nil
}
