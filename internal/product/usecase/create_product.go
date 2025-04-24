package product

import (
	"context"
	"errors"

	domain "sago-sample/internal/product/domain"
)

// CreateProductInput represents the input data for creating a product
type CreateProductInput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
}

// CreateProductOutput represents the output data after creating a product
type CreateProductOutput struct {
	ID          string
	Name        string
	Description string
	Price       uint
	Currency    string
	Stock       uint
}

// CreateProductUseCase defines the use case for creating a product
type CreateProductUseCase struct {
	productService *domain.Service
}

// NewCreateProductUseCase creates a new instance of CreateProductUseCase
func NewCreateProductUseCase(productService *domain.Service) *CreateProductUseCase {
	return &CreateProductUseCase{
		productService: productService,
	}
}

// Execute runs the use case
func (uc *CreateProductUseCase) Execute(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error) {
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

	// Call domain service to create product
	createdProduct, err := uc.productService.CreateProduct(ctx, productID, productName, productDescription, price, stock)
	if err != nil {
		if errors.Is(err, domain.ErrProductExists) {
			return nil, errors.New("product with this ID already exists")
		}
		return nil, err
	}

	// Map domain entity to output
	return &CreateProductOutput{
		ID:          createdProduct.ID().String(),
		Name:        createdProduct.Name().String(),
		Description: createdProduct.Description().String(),
		Price:       createdProduct.Price().Amount(),
		Currency:    createdProduct.Price().Currency(),
		Stock:       createdProduct.Stock().Quantity(),
	}, nil
}
