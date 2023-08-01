package graph

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/graph/model"
	"context"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (r *mutationResolver) UpdateToDo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteToDo(ctx context.Context, id string) (*model.Todo, error) {
	panic("not implemented")
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (*mutationResolver) DeleteCarePlan(ctx context.Context, id string) (*model.CarePlan, error) {
	panic("unimplemented")
}
