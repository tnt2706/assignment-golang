package graph

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/model"
	"context"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented")
}

// Mutation returns graph.MutationResolver implementation.

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
