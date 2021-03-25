package repository

import (
	"fmt"

	"github.com/sk62793/todo_server/graph/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Create(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}
