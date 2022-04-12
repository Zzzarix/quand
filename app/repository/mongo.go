package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var ctx = context.Background()

func mongoClient(db Db) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@maincluster0.ak4l7.mongodb.net/%s?retryWrites=true&w=majority", db.Username, db.Password, db.Name)).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	return client
}
