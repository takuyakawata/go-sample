package memory_test

import (
	"context"
	"sago-sample/internal/product/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	domain "sago-sample/internal/product/domain"
)

func TestProductRepository_SaveAndFindByID(t *testing.T) {
	// Create repository
	repo := infrastructure.NewProductRepository()

	// Create a product
	productID, _ := domain.NewProductID("prod-123")
	productName, _ := domain.NewProductName("Test Product")
	productDesc, _ := domain.NewProductDescription("This is a test product")
	price, _ := domain.NewPrice(1000, "USD")
	stock := domain.NewStock(10)

	product, err := domain.NewProduct(productID, productName, productDesc, price, stock)
	require.NoError(t, err, "Failed to create product")

	// Save the product
	ctx := context.Background()
	err = repo.Save(ctx, product)
	require.NoError(t, err, "Failed to save product")

	// Find the product by ID
	foundProduct, err := repo.FindByID(ctx, productID)
	require.NoError(t, err, "Failed to find product")

	// Verify the product
	assert.Equal(t, product.ID(), foundProduct.ID(), "Product ID should match")
	assert.Equal(t, product.Name(), foundProduct.Name(), "Product name should match")
	assert.Equal(t, product.Description(), foundProduct.Description(), "Product description should match")
	assert.Equal(t, product.Price(), foundProduct.Price(), "Product price should match")
	assert.Equal(t, product.Stock(), foundProduct.Stock(), "Product stock should match")
}

func TestProductRepository_FindByID_NotFound(t *testing.T) {
	// Create repository
	repo := infrastructure.NewProductRepository()

	// Try to find a non-existent product
	ctx := context.Background()
	productID, _ := domain.NewProductID("non-existent")

	_, err := repo.FindByID(ctx, productID)
	assert.Error(t, err, "Should return error for non-existent product")
	assert.Equal(t, domain.ErrProductNotFound, err, "Should return ErrProductNotFound")
}

func TestProductRepository_FindAll(t *testing.T) {
	// Create repository
	repo := infrastructure.NewProductRepository()

	// Create products
	product1, _ := domain.NewProduct(
		domain.MustNewProductID("prod-1"),
		domain.MustNewProductName("Product 1"),
		domain.MustNewProductDescription("Description 1"),
		domain.MustNewPrice(100, "USD"),
		domain.NewStock(5),
	)

	product2, _ := domain.NewProduct(
		domain.MustNewProductID("prod-2"),
		domain.MustNewProductName("Product 2"),
		domain.MustNewProductDescription("Description 2"),
		domain.MustNewPrice(200, "USD"),
		domain.NewStock(10),
	)

	// Save products
	ctx := context.Background()
	_ = repo.Save(ctx, product1)
	_ = repo.Save(ctx, product2)

	// Find all products
	products, err := repo.FindAll(ctx)
	require.NoError(t, err, "Failed to find all products")

	// Verify products
	assert.Len(t, products, 2, "Should return 2 products")

	// Check if both products are in the result
	foundProduct1 := false
	foundProduct2 := false

	for _, p := range products {
		if p.ID().String() == "prod-1" {
			foundProduct1 = true
		}
		if p.ID().String() == "prod-2" {
			foundProduct2 = true
		}
	}

	assert.True(t, foundProduct1, "Product 1 should be in the result")
	assert.True(t, foundProduct2, "Product 2 should be in the result")
}

func TestProductRepository_Delete(t *testing.T) {
	// Create repository
	repo := infrastructure.NewProductRepository()

	// Create a product
	productID, _ := domain.NewProductID("prod-123")
	productName, _ := domain.NewProductName("Test Product")
	productDesc, _ := domain.NewProductDescription("This is a test product")
	price, _ := domain.NewPrice(1000, "USD")
	stock := domain.NewStock(10)

	product, _ := domain.NewProduct(productID, productName, productDesc, price, stock)

	// Save the product
	ctx := context.Background()
	_ = repo.Save(ctx, product)

	// Delete the product
	err := repo.Delete(ctx, productID)
	require.NoError(t, err, "Failed to delete product")

	// Try to find the deleted product
	_, err = repo.FindByID(ctx, productID)
	assert.Error(t, err, "Should return error for deleted product")
	assert.Equal(t, domain.ErrProductNotFound, err, "Should return ErrProductNotFound")
}

func TestProductRepository_FindByCategory(t *testing.T) {
	// Create repository
	repo := infrastructure.NewProductRepository()

	// Create category
	categoryID, _ := domain.NewCategoryID("cat-1")
	categoryName, _ := domain.NewCategoryName("Electronics")
	category, _ := domain.NewCategory(categoryID, categoryName)

	// Create products
	product1, _ := domain.NewProduct(
		domain.MustNewProductID("prod-1"),
		domain.MustNewProductName("Product 1"),
		domain.MustNewProductDescription("Description 1"),
		domain.MustNewPrice(100, "USD"),
		domain.NewStock(5),
	)

	product2, _ := domain.NewProduct(
		domain.MustNewProductID("prod-2"),
		domain.MustNewProductName("Product 2"),
		domain.MustNewProductDescription("Description 2"),
		domain.MustNewPrice(200, "USD"),
		domain.NewStock(10),
	)

	// Add category to product1 only
	product1.AddCategory(category)

	// Save products
	ctx := context.Background()
	_ = repo.Save(ctx, product1)
	_ = repo.Save(ctx, product2)

	// Find products by category
	products, err := repo.FindByCategory(ctx, categoryID)
	require.NoError(t, err, "Failed to find products by category")

	// Verify products
	assert.Len(t, products, 1, "Should return 1 product")
	assert.Equal(t, "prod-1", products[0].ID().String(), "Should return product 1")
}
