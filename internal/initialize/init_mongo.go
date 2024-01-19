package initialize

import (
	"assignment/internal/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx    context.Context
	client *mongo.Client
)

func ConnectMongo() *mongo.Client {

	ctx := context.TODO()

	mongoConfig := config.GetMongoConfig()
	clientOptions := options.Client().ApplyURI(mongoConfig.DbCommonConnectString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("Mongo connection established")
	return client
}

var DB *mongo.Client = ConnectMongo()

func GetCollection(client *mongo.Client, collectionName string) (*mongo.Collection, context.Context) {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection, ctx
}
