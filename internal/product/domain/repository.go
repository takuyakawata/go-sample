package product

import (
	"context"
	"errors"
)

// Common repository errors
var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductExists   = errors.New("product already exists")
)

// Repository defines the interface for product persistence
type Repository interface {
	// FindByID finds a product by its ID
	FindByID(ctx context.Context, id ProductID) (*Product, error)

	// FindAll returns all products
	FindAll(ctx context.Context) ([]*Product, error)

	// FindByCategory finds products by category ID
	FindByCategory(ctx context.Context, categoryID CategoryID) ([]*Product, error)

	// Save persists a product
	Save(ctx context.Context, product *Product) error

	// Delete removes a product
	Delete(ctx context.Context, id ProductID) error
}
