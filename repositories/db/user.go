package db

import (
	"context"
	"gorm.io/gorm"
	"todolist/models"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *models.User) error {
	return dao.Create(&user).Error
}

func (dao *UserDao) DeleteUser(name string) error {
	return dao.Where("user_name = ?", name).Unscoped().Delete(&models.User{}).Error
}
