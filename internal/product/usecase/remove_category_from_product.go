package product

import (
	"context"
	"errors"

	domain "sago-sample/internal/product/domain"
)

// RemoveCategoryFromProductInput represents the input data for removing a category from a product
type RemoveCategoryFromProductInput struct {
	ProductID  string
	CategoryID string
}

// RemoveCategoryFromProductOutput represents the output data after removing a category from a product
type RemoveCategoryFromProductOutput struct {
	ProductID   string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

// RemoveCategoryFromProductUseCase defines the use case for removing a category from a product
type RemoveCategoryFromProductUseCase struct {
	productService *domain.Service
}

// NewRemoveCategoryFromProductUseCase creates a new instance of RemoveCategoryFromProductUseCase
func NewRemoveCategoryFromProductUseCase(productService *domain.Service) *RemoveCategoryFromProductUseCase {
	return &RemoveCategoryFromProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *RemoveCategoryFromProductUseCase) Execute(ctx context.Context, input RemoveCategoryFromProductInput) (*RemoveCategoryFromProductOutput, error) {
	// Create value objects
	productID, err := domain.NewProductID(input.ProductID)
	if err != nil {
		return nil, err
	}

	categoryID, err := domain.NewCategoryID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	// Call domain service to remove category from product
	updatedProduct, err := uc.productService.RemoveCategoryFromProduct(ctx, productID, categoryID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Map domain entity to output
	categories := make([]CategoryOutput, 0, len(updatedProduct.Categories()))
	for _, c := range updatedProduct.Categories() {
		categories = append(categories, CategoryOutput{
			ID:   c.ID().String(),
			Name: c.Name().String(),
		})
	}

	return &RemoveCategoryFromProductOutput{
		ProductID:   updatedProduct.ID().String(),
		Name:        updatedProduct.Name().String(),
		Description: updatedProduct.Description().String(),
		Price:       updatedProduct.Price().Amount(),
		Currency:    updatedProduct.Price().Currency(),
		Stock:       updatedProduct.Stock().Quantity(),
		Categories:  categories,
	}, nil
}
