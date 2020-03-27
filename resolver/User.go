package resolver

import (
	"context"
	ctx "eaciit/gopher/dbcontext"
	graphql "github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// User jdjfdj
func (r *RootResolver) User(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	// Find user:
	for _, user := range users {
		if args.UserID == user.UserID {
			// Found user:
			return &UserResolver{user}, nil
		}
	}

	// db.getCollection('User').aggregate([{
	// 	$lookup:
	// 	 {
	// 	   from: 'Notes',
	// 	   localField: 'UserID',
	// 	   foreignField: 'UserID',
	// 	   as: 'Notes'
	// 	 }
	//  }])
	groupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "Notes"},
			{Key: "localField", Value: "UserID"},
			{Key: "foreignField", Value: "UserID"},
			{Key: "as", Value: "Notes"},
		}},
	}
	cursor, err := ctx.MainContext("User").Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	// Didn’t find user:
	return nil, nil
}

// Email prop
func (r *UserResolver) Email() string {
	return r.u.Email
}
