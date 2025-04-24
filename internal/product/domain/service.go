package product

import (
	"context"
	"errors"
)

// Service provides domain operations for products
type Service struct {
	repo Repository
}

// NewService creates a new product service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(ctx context.Context, id ProductID, name ProductName, description ProductDescription, price Price, stock Stock) (*Product, error) {
	// Check if product with the same ID already exists
	existingProduct, err := s.repo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, ErrProductNotFound) {
		return nil, err
	}

	if existingProduct != nil {
		return nil, ErrProductExists
	}

	// Create new product
	product, err := NewProduct(id, name, description, price, stock)
	if err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct updates an existing product
func (s *Service) UpdateProduct(ctx context.Context, id ProductID, name ProductName, description ProductDescription, price Price, stock Stock) (*Product, error) {
	// Find existing product
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update product fields
	product.UpdateName(name)
	product.UpdateDescription(description)
	product.UpdatePrice(price)
	product.UpdateStock(stock)

	// Save to repository
	if err := s.repo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct deletes a product
func (s *Service) DeleteProduct(ctx context.Context, id ProductID) error {
	// Check if product exists
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete from repository
	return s.repo.Delete(ctx, id)
}

// GetProductByID retrieves a product by ID
func (s *Service) GetProductByID(ctx context.Context, id ProductID) (*Product, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAllProducts retrieves all products
func (s *Service) GetAllProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.FindAll(ctx)
}

// GetProductsByCategory retrieves products by category
func (s *Service) GetProductsByCategory(ctx context.Context, categoryID CategoryID) ([]*Product, error) {
	return s.repo.FindByCategory(ctx, categoryID)
}

// AddCategoryToProduct adds a category to a product
func (s *Service) AddCategoryToProduct(ctx context.Context, productID ProductID, category *Category) (*Product, error) {
	// Find existing product
	product, err := s.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// Add category to product
	product.AddCategory(category)

	// Save to repository
	if err := s.repo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// RemoveCategoryFromProduct removes a category from a product
func (s *Service) RemoveCategoryFromProduct(ctx context.Context, productID ProductID, categoryID CategoryID) (*Product, error) {
	// Find existing product
	product, err := s.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	// Remove category from product
	product.RemoveCategory(categoryID)

	// Save to repository
	if err := s.repo.Save(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
