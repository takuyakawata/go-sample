package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	domain "sago-sample/internal/product/domain"
	usecase "sago-sample/internal/product/usecase"
)

// We're reusing the MockProductRepository from add_category_to_product_test.go
// No need to redefine it here since both files are in the same package

func TestCreateProductUseCase_Execute(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockProductRepository)

	// Create real service with mock repository
	productService := domain.NewService(mockRepo)

	// Create use case
	useCase := usecase.NewCreateProductUseCase(productService)

	// Test data
	ctx := context.Background()
	input := usecase.CreateProductInput{
		ID:          "prod-123",
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       1000,
		Currency:    "USD",
		Stock:       10,
	}

	// Create product ID for expectation
	productID, _ := domain.NewProductID(input.ID)

	// Setup expectations
	// First, the service will check if the product exists
	mockRepo.On("FindByID", ctx, productID).Return(nil, domain.ErrProductNotFound)
	// Then, it will save the new product
	mockRepo.On("Save", ctx, mock.Anything).Return(nil)

	// Execute use case
	output, err := useCase.Execute(ctx, input)

	// Assertions
	require.NoError(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.ID, output.ID)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Description, output.Description)
	assert.Equal(t, input.Price, output.Price)
	assert.Equal(t, input.Currency, output.Currency)
	assert.Equal(t, input.Stock, output.Stock)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestCreateProductUseCase_Execute_Error(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockProductRepository)

	// Create real service with mock repository
	productService := domain.NewService(mockRepo)

	// Create use case
	useCase := usecase.NewCreateProductUseCase(productService)

	// Test data
	ctx := context.Background()
	input := usecase.CreateProductInput{
		ID:          "prod-123",
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       1000,
		Currency:    "USD",
		Stock:       10,
	}

	// Create product ID for expectation
	productID, _ := domain.NewProductID(input.ID)

	// Setup expectations - product already exists
	mockRepo.On("FindByID", ctx, productID).Return(&domain.Product{}, nil)

	// Execute use case
	output, err := useCase.Execute(ctx, input)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, output)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}
