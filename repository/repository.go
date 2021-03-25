package repository

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) TaskRepository {
	tR := NewTaskRepository(db)

	return tR
}
