package resolver

import (
	"context"
	ctx "eaciit/gopher/dbcontext"
	types "eaciit/gopher/types"
	graphql "github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (r *RootResolver) Users() ([]*UserResolver, error) {
	var userRxs []*UserResolver
	// for _, u := range users {
	// 	userRxs = append(userRxs, &UserResolver{u})
	// }

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

	for _, result := range results {
		//	fmt.Println(result)
		notes := []*types.Note{}
		resultNotes := result["Notes"].(bson.A)
		for _, note := range resultNotes {
			noteItem := note.(bson.M)
			notes = append(notes, &types.Note{
				UserID: graphql.ID(noteItem["UserID"].(string)),
				Data:   noteItem["Data"].(string),
			})
		}
		userRxs = append(userRxs, &UserResolver{
			&types.User{
				UserID:   graphql.ID(result["UserID"].(string)),
				Username: result["Username"].(string),
				Emoji:    result["Emoji"].(string),
				Notes:    notes,
			},
		})

	}

	return userRxs, nil
}
