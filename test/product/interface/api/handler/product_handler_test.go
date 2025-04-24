package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sago-sample/internal/product/interface/api/handler"
	usecase "sago-sample/internal/product/usecase"
)

// MockCreateProductUseCase mocks the CreateProductUseCase
type MockCreateProductUseCase struct {
	mock.Mock
}

func (m *MockCreateProductUseCase) Execute(ctx context.Context, input usecase.CreateProductInput) (*usecase.CreateProductOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.CreateProductOutput), args.Error(1)
}

// MockGetProductByIDUseCase mocks the GetProductByIDUseCase
type MockGetProductByIDUseCase struct {
	mock.Mock
}

func (m *MockGetProductByIDUseCase) Execute(ctx context.Context, input usecase.GetProductByIDInput) (*usecase.GetProductByIDOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.GetProductByIDOutput), args.Error(1)
}

// MockUpdateProductUseCase mocks the UpdateProductUseCase
type MockUpdateProductUseCase struct {
	mock.Mock
}

func (m *MockUpdateProductUseCase) Execute(ctx context.Context, input usecase.UpdateProductInput) (*usecase.UpdateProductOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.UpdateProductOutput), args.Error(1)
}

// MockDeleteProductUseCase mocks the DeleteProductUseCase
type MockDeleteProductUseCase struct {
	mock.Mock
}

func (m *MockDeleteProductUseCase) Execute(ctx context.Context, input usecase.DeleteProductInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

// MockGetAllProductsUseCase mocks the GetAllProductsUseCase
type MockGetAllProductsUseCase struct {
	mock.Mock
}

func (m *MockGetAllProductsUseCase) Execute(ctx context.Context) (*usecase.GetAllProductsOutput, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.GetAllProductsOutput), args.Error(1)
}

func TestProductHandler_CreateProduct(t *testing.T) {
	// Create mocks
	createUseCase := new(MockCreateProductUseCase)
	updateUseCase := new(MockUpdateProductUseCase)
	deleteUseCase := new(MockDeleteProductUseCase)
	getByIDUseCase := new(MockGetProductByIDUseCase)
	getAllUseCase := new(MockGetAllProductsUseCase)

	// Create handler
	productHandler := handler.NewProductHandler(
		createUseCase,
		updateUseCase,
		deleteUseCase,
		getByIDUseCase,
		getAllUseCase,
	)

	// Create request body
	reqBody := map[string]interface{}{
		"id":          "prod-123",
		"name":        "Test Product",
		"description": "This is a test product",
		"price":       1000,
		"currency":    "USD",
		"stock":       10,
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Setup expectations
	expectedInput := usecase.CreateProductInput{
		ID:          "prod-123",
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       1000,
		Currency:    "USD",
		Stock:       10,
	}

	expectedOutput := &usecase.CreateProductOutput{
		ID:          "prod-123",
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       1000,
		Currency:    "USD",
		Stock:       10,
	}

	createUseCase.On("Execute", mock.Anything, expectedInput).Return(expectedOutput, nil)

	// Call handler
	productHandler.CreateProduct(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "prod-123", response["id"])
	assert.Equal(t, "Test Product", response["name"])
	assert.Equal(t, "This is a test product", response["description"])
	assert.Equal(t, float64(1000), response["price"])
	assert.Equal(t, "USD", response["currency"])
	assert.Equal(t, float64(10), response["stock"])

	// Verify expectations
	createUseCase.AssertExpectations(t)
}

func TestProductHandler_GetProductByID(t *testing.T) {
	// Create mocks
	createUseCase := new(MockCreateProductUseCase)
	updateUseCase := new(MockUpdateProductUseCase)
	deleteUseCase := new(MockDeleteProductUseCase)
	getByIDUseCase := new(MockGetProductByIDUseCase)
	getAllUseCase := new(MockGetAllProductsUseCase)

	// Create handler
	productHandler := handler.NewProductHandler(
		createUseCase,
		updateUseCase,
		deleteUseCase,
		getByIDUseCase,
		getAllUseCase,
	)

	// Create request
	req := httptest.NewRequest(http.MethodGet, "/products/prod-123", nil)

	// Create response recorder
	w := httptest.NewRecorder()

	// Setup expectations
	expectedInput := usecase.GetProductByIDInput{
		ID: "prod-123",
	}

	expectedOutput := &usecase.GetProductByIDOutput{
		ID:          "prod-123",
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       1000,
		Currency:    "USD",
		Stock:       10,
		Categories:  []usecase.CategoryOutput{},
	}

	getByIDUseCase.On("Execute", mock.Anything, expectedInput).Return(expectedOutput, nil)

	// Call handler
	productHandler.GetProductByID(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "prod-123", response["id"])
	assert.Equal(t, "Test Product", response["name"])
	assert.Equal(t, "This is a test product", response["description"])
	assert.Equal(t, float64(1000), response["price"])
	assert.Equal(t, "USD", response["currency"])
	assert.Equal(t, float64(10), response["stock"])

	// Verify expectations
	getByIDUseCase.AssertExpectations(t)
}
