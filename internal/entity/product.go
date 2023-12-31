package entity

import (
	"errors"
	"go-api-crud/pkg/entity"
	"time"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID = errors.New("id is invalid")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice = errors.New("price is invalid")
)

type Product struct {
	ID					entity.ID     `json:"id"`
	Name				string `json:"name"`
	Price				int    `json:"price"`
	CreatedAt		time.Time  `json:"created_at"`
}


func NewProduct(name string, price int) (*Product, error) {
	return &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}, nil
}


func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}