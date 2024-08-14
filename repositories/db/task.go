package db

import (
	"context"
	"gorm.io/gorm"
	"todolist/models"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(task *models.Task) error {
	return dao.Create(&task).Error
}
