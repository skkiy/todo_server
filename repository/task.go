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
	Tasks(
		filterCondition *model.FilterCondition,
		pageCondition *model.PageCondition,
		edgeOrder *model.EdgeOrder,
	) ([]*model.Task, error)
	TotalCounts(filterCondition *model.FilterCondition) (int, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (r *taskRepository) FindByID(id string) (*model.Task, error) {
	base := r.db.Table("tasks").Where("id = ?", id)
	task := new(model.Task)
	if err := base.First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to find task: record not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find task: %w", err)
	}
	return task, nil
}

func (r *taskRepository) Tasks(filterCondition *model.FilterCondition, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder) ([]*model.Task, error) {
	base := r.db.Table("tasks")
	var tasks []*model.Task
	if err := base.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) TotalCounts(filterCondition *model.FilterCondition) (int, error) {
	base := r.db.Table("tasks")
	base = r.filter(filterCondition, base)

	var counts int64
	if err := base.Count(&counts).Error; err != nil {
		return 0, err
	}
	return int(counts), nil
}

func (r *taskRepository) filter(filterCondition *model.FilterCondition, base *gorm.DB) *gorm.DB {
	return base
}
