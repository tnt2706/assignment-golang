package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"assignment/configs"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	mongoConfig := configs.GetMongoConfig()
	clientOptions := options.Client().ApplyURI(mongoConfig.COMMON_DB)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
}

func main() {

}
