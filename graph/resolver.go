package graph

import "github.com/sk62793/todo_server/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	taskRepo repository.TaskRepository
}

func NewResolver(tR repository.TaskRepository) *Resolver {
	return &Resolver{
		taskRepo: tR,
	}
}
