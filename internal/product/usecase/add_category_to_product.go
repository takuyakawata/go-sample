package product

import (
	"context"
	"errors"

	domain "sago-sample/internal/product/domain"
)

// AddCategoryToProductInput represents the input data for adding a category to a product
type AddCategoryToProductInput struct {
	ProductID    string
	CategoryID   string
	CategoryName string
}

// AddCategoryToProductOutput represents the output data after adding a category to a product
type AddCategoryToProductOutput struct {
	ProductID   string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
	Categories  []CategoryOutput
}

// CategoryOutput represents a category in the output
type CategoryOutput struct {
	ID   string
	Name string
}

// AddCategoryToProductUseCase defines the use case for adding a category to a product
type AddCategoryToProductUseCase struct {
	productService *domain.Service
}

// NewAddCategoryToProductUseCase creates a new instance of AddCategoryToProductUseCase
func NewAddCategoryToProductUseCase(productService *domain.Service) *AddCategoryToProductUseCase {
	return &AddCategoryToProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *AddCategoryToProductUseCase) Execute(ctx context.Context, input AddCategoryToProductInput) (*AddCategoryToProductOutput, error) {
	// Create value objects
	productID, err := domain.NewProductID(input.ProductID)
	if err != nil {
		return nil, err
	}

	categoryID, err := domain.NewCategoryID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	categoryName, err := domain.NewCategoryName(input.CategoryName)
	if err != nil {
		return nil, err
	}

	// Create category
	category, err := domain.NewCategory(categoryID, categoryName)
	if err != nil {
		return nil, err
	}

	// Call domain service to add category to product
	updatedProduct, err := uc.productService.AddCategoryToProduct(ctx, productID, category)
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

	return &AddCategoryToProductOutput{
		ProductID:   updatedProduct.ID().String(),
		Name:        updatedProduct.Name().String(),
		Description: updatedProduct.Description().String(),
		Price:       updatedProduct.Price().Amount(),
		Currency:    updatedProduct.Price().Currency(),
		Stock:       updatedProduct.Stock().Quantity(),
		Categories:  categories,
	}, nil
}
