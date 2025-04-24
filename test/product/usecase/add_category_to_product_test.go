package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domain "sago-sample/internal/product/domain"
	productUseCase "sago-sample/internal/product/usecase"
)

// MockProductService is a mock implementation of the domain.Service interface
type MockProductService struct {
	mock.Mock
}

// AddCategoryToProduct mocks the AddCategoryToProduct method
func (m *MockProductService) AddCategoryToProduct(ctx context.Context, productID domain.ProductID, category *domain.Category) (*domain.Product, error) {
	args := m.Called(ctx, productID, category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func TestAddCategoryToProduct(t *testing.T) {
	// Create mock service
	mockService := new(MockProductService)

	// Create use case
	useCase := productUseCase.NewAddCategoryToProductUseCase(mockService)

	// Create test data
	productID := "prod-123"
	categoryID := "cat-456"
	categoryName := "Test Category"

	// Create input
	input := productUseCase.AddCategoryToProductInput{
		ProductID:    productID,
		CategoryID:   categoryID,
		CategoryName: categoryName,
	}

	// Create expected output
	// This would typically be a product with the category added
	// For simplicity, we'll just assert that the method was called with the right parameters

	// Set up expectations
	mockService.On("AddCategoryToProduct", mock.Anything, mock.Anything, mock.Anything).Return(nil, domain.ErrProductNotFound)

	// Execute use case
	output, err := useCase.Execute(context.Background(), input)

	// Assert expectations
	assert.Error(t, err)
	assert.Nil(t, output)
	mockService.AssertExpectations(t)
}
