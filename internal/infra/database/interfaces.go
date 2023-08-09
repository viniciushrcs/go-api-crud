package database

import "go-api-crud/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindAll() ([]*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type ProductDBInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindByID(id int) (*entity.Product, error)
}