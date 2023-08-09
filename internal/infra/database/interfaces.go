package database

import "go-api-crud/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindAll() ([]*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}