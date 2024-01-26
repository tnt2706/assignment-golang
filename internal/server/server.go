package server

import (
	"assignment/internal/graph/directive"
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/loader"
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"

	resolver "assignment/internal/graph/resolver"
	"assignment/internal/initialize"
	repo "assignment/internal/repository"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	log "github.com/sirupsen/logrus"
)

func Init() {
	const PORT = "4001"
	router := chi.NewRouter()

	// Init mongo
	db := initialize.ConnectMongo()

	var (
		userRepo = repo.NewUserRepo(db.Collection("users"))
		todoRepo = repo.NewTodoRepo(db.Collection("todos"))
	)

	log.SetFormatter(&log.JSONFormatter{})
	router.Use(middleware.AuthMiddleware)

	c := graph.Config{
		Resolvers: &resolver.Resolver{
			UserRepo: userRepo,
			TodoRepo: todoRepo,
		},
	}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*model.Role) (res interface{}, err error) {
		return directive.Auth(ctx, next, roles)
	}

	dbLoader := loader.DBLoader{
		UserRepo: userRepo,
	}

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(c))
	srv = loader.Middleware(&dbLoader, srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
