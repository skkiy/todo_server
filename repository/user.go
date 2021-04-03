package repository

import (
	"errors"
	"fmt"

	"github.com/sk62793/todo_server/graph/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	base := r.db.Table("users").Where("id = ?", id)
	user := new(model.User)
	if err := base.First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to find user: record not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	if err := r.db.Table("users").Where("id = ?", user.ID).Updates(model.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Twitter: user.Twitter,
	}).Error; err != nil {
		return err
	}
	return nil
}
