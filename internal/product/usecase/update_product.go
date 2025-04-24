package product

import (
	"context"
	"errors"

	domain "sago-sample/internal/product/domain"
)

// UpdateProductInput represents the input data for updating a product
type UpdateProductInput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
}

// UpdateProductOutput represents the output data after updating a product
type UpdateProductOutput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
}

// UpdateProductUseCase defines the use case for updating a product
type UpdateProductUseCase struct {
	productService *domain.Service
}

// NewUpdateProductUseCase creates a new instance of UpdateProductUseCase
func NewUpdateProductUseCase(productService *domain.Service) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *UpdateProductUseCase) Execute(ctx context.Context, input UpdateProductInput) (*UpdateProductOutput, error) {
	// Create value objects
	productID, err := domain.NewProductID(input.ID)
	if err != nil {
		return nil, err
	}

	productName, err := domain.NewProductName(input.Name)
	if err != nil {
		return nil, err
	}

	productDescription, err := domain.NewProductDescription(input.Description)
	if err != nil {
		return nil, err
	}

	price, err := domain.NewPrice(input.Price, input.Currency)
	if err != nil {
		return nil, err
	}

	stock := domain.NewStock(input.Stock)

	// Call domain service to update product
	updatedProduct, err := uc.productService.UpdateProduct(ctx, productID, productName, productDescription, price, stock)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Map domain entity to output
	return &UpdateProductOutput{
		ID:          updatedProduct.ID().String(),
		Name:        updatedProduct.Name().String(),
		Description: updatedProduct.Description().String(),
		Price:       updatedProduct.Price().Amount(),
		Currency:    updatedProduct.Price().Currency(),
		Stock:       updatedProduct.Stock().Quantity(),
	}, nil
}
