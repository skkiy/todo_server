package repository

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) (TaskRepository, UserRepository) {
	tR := NewTaskRepository(db)
	uR := NewUserRepository(db)

	return tR, uR
}
