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

// MockProductService is a mock implementation of the product.Service
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(ctx context.Context, id domain.ProductID, name domain.ProductName, description domain.ProductDescription, price domain.Price, stock domain.Stock) (*domain.Product, error) {
	args := m.Called(ctx, id, name, description, price, stock)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(ctx context.Context, id domain.ProductID, name domain.ProductName, description domain.ProductDescription, price domain.Price, stock domain.Stock) (*domain.Product, error) {
	args := m.Called(ctx, id, name, description, price, stock)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(ctx context.Context, id domain.ProductID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductService) GetProductByID(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductService) AddCategoryToProduct(ctx context.Context, productID domain.ProductID, categoryID domain.CategoryID, categoryName domain.CategoryName) (*domain.Product, error) {
	args := m.Called(ctx, productID, categoryID, categoryName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) RemoveCategoryFromProduct(ctx context.Context, productID domain.ProductID, categoryID domain.CategoryID) (*domain.Product, error) {
	args := m.Called(ctx, productID, categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductService) GetProductsByCategory(ctx context.Context, categoryID domain.CategoryID) ([]*domain.Product, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func TestCreateProductUseCase_Execute(t *testing.T) {
	// Create mock service
	mockService := new(MockProductService)

	// Create use case
	useCase := usecase.NewCreateProductUseCase(mockService)

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

	// Create expected domain objects
	productID, _ := domain.NewProductID(input.ID)
	productName, _ := domain.NewProductName(input.Name)
	productDesc, _ := domain.NewProductDescription(input.Description)
	price, _ := domain.NewPrice(input.Price, input.Currency)
	stock := domain.NewStock(input.Stock)

	// Create expected product
	expectedProduct, _ := domain.NewProduct(productID, productName, productDesc, price, stock)

	// Setup expectations
	mockService.On("CreateProduct", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedProduct, nil)

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
	mockService.AssertExpectations(t)
}

func TestCreateProductUseCase_Execute_Error(t *testing.T) {
	// Create mock service
	mockService := new(MockProductService)

	// Create use case
	useCase := usecase.NewCreateProductUseCase(mockService)

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

	// Setup expectations - return an error
	mockService.On("CreateProduct", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, domain.ErrProductExists)

	// Execute use case
	output, err := useCase.Execute(ctx, input)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, output)

	// Verify expectations
	mockService.AssertExpectations(t)
}
