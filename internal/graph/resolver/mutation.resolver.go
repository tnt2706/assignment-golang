package graph

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/model"
	"context"
)

type mutationResolver struct{ *Resolver }

// CreateCarePlan implements graph.MutationResolver.
func (*mutationResolver) CreateCarePlan(ctx context.Context, input model.CarePlanInput) (*model.Response, error) {

	panic("unimplemented")
}

// DeleteUser implements graph.MutationResolver.
func (*mutationResolver) DeleteUser(ctx context.Context, id string) (*model.Response, error) {
	panic("unimplemented")
}

// UpdateUser implements graph.MutationResolver.
func (*mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.UserResponse, error) {
	panic("unimplemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.UserResponse, error) {
	// return &model.Todo{}, nil
	panic("not implemented")
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }
