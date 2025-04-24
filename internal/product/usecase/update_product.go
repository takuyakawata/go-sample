package product

import (
	"context"
	"errors"

	"github.com/sago-sample/internal/product/domain"
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
	productService *product.Service
}

// NewUpdateProductUseCase creates a new instance of UpdateProductUseCase
func NewUpdateProductUseCase(productService *product.Service) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *UpdateProductUseCase) Execute(ctx context.Context, input UpdateProductInput) (*UpdateProductOutput, error) {
	// Create value objects
	productID, err := product.NewProductID(input.ID)
	if err != nil {
		return nil, err
	}

	productName, err := product.NewProductName(input.Name)
	if err != nil {
		return nil, err
	}

	productDescription, err := product.NewProductDescription(input.Description)
	if err != nil {
		return nil, err
	}

	price, err := product.NewPrice(input.Price, input.Currency)
	if err != nil {
		return nil, err
	}

	stock := product.NewStock(input.Stock)

	// Call domain service to update product
	updatedProduct, err := uc.productService.UpdateProduct(ctx, productID, productName, productDescription, price, stock)
	if err != nil {
		if errors.Is(err, product.ErrProductNotFound) {
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
