package repository

import (
	"errors"
	"fmt"

	"github.com/sk62793/todo_server/graph/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	Tasks(orderKey *model.TaskOrderKey, orderDirection *model.OrderDirection) ([]*model.Task, error)
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

func (t *taskRepository) FindByID(id string) (*model.Task, error) {
	base := t.db.Table("tasks").Where("id = ?", id)
	task := new(model.Task)
	if err := base.First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to find task: record not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find task: %w", err)
	}
	return task, nil
}

func (t *taskRepository) Tasks(orderKey *model.TaskOrderKey, orderDirection *model.OrderDirection) ([]*model.Task, error) {
	base := t.db.Table("tasks")
	var tasks []*model.Task
	if err := base.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
