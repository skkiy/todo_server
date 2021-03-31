package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
	task := &model.Task{
		ID:          input.ID,
		Title:       *input.Title,
		Description: *input.Description,
		Deadline:    *input.Deadline,
		IsCompleted: *input.IsCompleted,
	}

	if err := r.taskRepo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (r *queryResolver) Tasks(ctx context.Context, filterCondition *model.FilterCondition, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder) (*model.TaskConnection, error) {
	totalCount, err := r.taskRepo.TotalCounts(filterCondition)
	if err != nil {
		return nil, err
	}
	if totalCount == 0 {
		return model.EmptyTaskConnection(), nil
	}

	totalPage := pageCondition.TotalPage(totalCount)

	pageInfo := &model.PageInfo{
		HasNextPage:     (totalPage - pageCondition.MoveToPageNo()) >= 1,
		HasPreviousPage: pageCondition.MoveToPageNo() > 1,
	}

	tasks, err := r.taskRepo.Tasks(filterCondition, pageCondition, edgeOrder)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return model.EmptyTaskConnection(), nil
	}

	edges := make([]*model.TaskEdge, len(tasks))
	for i, task := range tasks {
		cursor := task.ID
		edges[i] = &model.TaskEdge{
			Cursor: cursor,
			Node: &model.Task{
				ID:          task.ID,
				Title:       task.Title,
				Description: task.Description,
				CreatedAt:   task.CreatedAt,
				Deadline:    task.Deadline,
				IsCompleted: task.IsCompleted,
			},
		}
		if i == 0 {
			pageInfo.StartCursor = cursor
		}
		if i == len(tasks)-1 {
			pageInfo.EndCursor = cursor
		}
	}

	return &model.TaskConnection{
		PageInfo:   pageInfo,
		Edges:      edges,
		TotalCount: totalCount,
	}, nil
}

func (r *queryResolver) Task(ctx context.Context, id *string) (*model.Task, error) {
	task, err := r.taskRepo.FindByID(*id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
