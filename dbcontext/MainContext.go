package dbcontext

import (
	"context"
	helper "eaciit/gopher/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// MainContext context of default
func MainContext(name string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_HOST"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		helper.LogFatal(err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		helper.LogFatal(err.Error())
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(name)
	return collection
}
