package main

import (
	graph "assignment/internal/graph/generate"
	resolver "assignment/internal/graph/resolver"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"assignment/configs"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	log "github.com/sirupsen/logrus"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	mongoConfig := configs.GetMongoConfig()
	clientOptions := options.Client().ApplyURI(mongoConfig.DbCommonConnectString)
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
	log.SetFormatter(&log.JSONFormatter{})

	const port = "4001"

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
