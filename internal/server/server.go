package server

import (
	graph "assignment/internal/graph/generate"
	resolver "assignment/internal/graph/resolver"
	"assignment/internal/initialize"
	repo "assignment/internal/repository"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	log "github.com/sirupsen/logrus"
)

func Init() {
	// Init mongo
	db := initialize.ConnectMongo()

	var (
		userRepo = repo.NewUserRepo(db.Collection("users"))
	)

	fmt.Print(userRepo)

	log.SetFormatter(&log.JSONFormatter{})

	const port = "4001"

	context := graph.Config{Resolvers: &resolver.Resolver{
		// UserRepo: userRepo,
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(context))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
