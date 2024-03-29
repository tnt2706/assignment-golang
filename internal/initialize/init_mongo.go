package initialize

import (
	"assignment/internal/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {

	var (
		ctx    context.Context
		client *mongo.Client
	)

	ctx = context.TODO()

	mongoConfig := config.GetMongoConfig()
	clientOptions := options.Client().ApplyURI(mongoConfig.DbCommonConnectString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic("Connected with mongodb problem")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("Connected with mongo success !")

	return client.Database("golang-api")
}
