package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/loader"
	"assignment/internal/model"
	"context"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return loader.GetUserLoader(ctx, obj.User)
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
