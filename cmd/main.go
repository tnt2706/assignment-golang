package main

import (
	graph "assignment/internal/graph/generate"
	resolver "assignment/internal/graph/resolver"
	"net/http"

	"assignment/internal/initialize"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	log "github.com/sirupsen/logrus"
)

func init() {
	initialize.ConnectMongo()
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
