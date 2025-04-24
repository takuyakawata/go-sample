package product

import (
	"context"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductExists   = errors.New("product already exists")
)

type Repository interface {
	FindByID(ctx context.Context, id ProductID) (*Product, error)
	FindAll(ctx context.Context) ([]*Product, error)
	FindByCategory(ctx context.Context, categoryID CategoryID) ([]*Product, error)
	Save(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id ProductID) error
}
