package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/sk62793/todo_server/graph/generated"
	"github.com/sk62793/todo_server/graph/model"
	"github.com/sk62793/todo_server/tools"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	task := &model.Task{
		ID:          tools.UUID(),
		Title:       input.Title,
		Description: *input.Description,
		CreatedAt:   time.Now(),
		Deadline:    input.Deadline,
		IsCompleted: false,
	}

	if err := r.taskRepo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context, orderKey *model.TaskOrderKey, orderDirection *model.OrderDirection) (*model.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
