package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domain "sago-sample/feature/product/domain"
	productUseCase "sago-sample/feature/product/usecase"
)

// MockProductRepository is a mock implementation of the domain.Repository interface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByID(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductRepository) FindByCategory(ctx context.Context, categoryID domain.CategoryID) ([]*domain.Product, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductRepository) Save(ctx context.Context, product *domain.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id domain.ProductID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestAddCategoryToProduct(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockProductRepository)

	// Create real service with mock repository
	productService := domain.NewService(mockRepo)

	// Create use case
	useCase := productUseCase.NewAddCategoryToProductUseCase(productService)

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

	// Set up expectations
	mockRepo.On("FindByID", mock.Anything, mock.Anything).Return(nil, domain.ErrProductNotFound)

	// Execute use case
	output, err := useCase.Execute(context.Background(), input)

	// Assert expectations
	assert.Error(t, err)
	assert.Nil(t, output)
	mockRepo.AssertExpectations(t)
}
