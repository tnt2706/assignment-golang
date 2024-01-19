package initialize

import (
	"assignment/internal/config"
	"assignment/internal/repository/repoimpl"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {

	var (
		ctx    context.Context
		client *mongo.Client
	)

	ctx = context.TODO()

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

	repoimpl.NewUserService(client.Database("golang-api").Collection("users"), ctx)
	return client
}
