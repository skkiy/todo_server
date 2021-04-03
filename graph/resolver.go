package graph

import "github.com/sk62793/todo_server/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewResolver(tR repository.TaskRepository, uR repository.UserRepository) *Resolver {
	return &Resolver{
		taskRepo: tR,
		userRepo: uR,
	}
}
