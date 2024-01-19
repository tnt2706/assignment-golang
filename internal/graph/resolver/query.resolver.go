package graph

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/model"
	"context"
)

type queryResolver struct{ *Resolver }

// User implements graph.QueryResolver.
func (*queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return &model.User{ID: "1111"}, nil
}

// User implements graph.QueryResolver.

func (r *queryResolver) Users(ctx context.Context, filter model.UserFilterInput) ([]*model.User, error) {

	panic("not implemented")
}

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }
