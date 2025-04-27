package product_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	product "sago-sample/feature/product/domain"
)

func TestNewProduct(t *testing.T) {
	// Create valid value objects
	id, err := product.NewProductID("prod-123")
	require.NoError(t, err, "Failed to create product ID")

	name, err := product.NewProductName("Test Product")
	require.NoError(t, err, "Failed to create product name")

	desc, err := product.NewProductDescription("This is a test product")
	require.NoError(t, err, "Failed to create product description")

	price, err := product.NewPrice(1000, "USD")
	require.NoError(t, err, "Failed to create price")

	stock := product.NewStock(10)

	// Test creating a valid product
	p, err := product.NewProduct(id, name, desc, price, stock)
	require.NoError(t, err, "Failed to create product")

	// Verify product properties
	assert.Equal(t, id, p.ID(), "Product ID should match")
	assert.Equal(t, name, p.Name(), "Product name should match")
	assert.Equal(t, desc, p.Description(), "Product description should match")
	assert.Equal(t, price, p.Price(), "Product price should match")
	assert.Equal(t, stock, p.Stock(), "Product stock should match")

	// Verify creation time is set
	assert.False(t, p.CreatedAt().IsZero(), "Product creation time should be set")
	assert.False(t, p.UpdatedAt().IsZero(), "Product update time should be set")

	// Test creating a product with empty ID
	emptyID, _ := product.NewProductID("")
	_, err = product.NewProduct(emptyID, name, desc, price, stock)
	assert.Error(t, err, "Should error when creating product with empty ID")
}

func TestProductUpdateMethods(t *testing.T) {
	// Create a valid product
	id, _ := product.NewProductID("prod-123")
	name, _ := product.NewProductName("Test Product")
	desc, _ := product.NewProductDescription("This is a test product")
	price, _ := product.NewPrice(1000, "USD")
	stock := product.NewStock(10)

	p, err := product.NewProduct(id, name, desc, price, stock)
	require.NoError(t, err, "Failed to create product")

	// Record the initial update time
	initialUpdateTime := p.UpdatedAt()

	// Wait a moment to ensure time difference
	time.Sleep(10 * time.Millisecond)

	// Test updating name
	newName, _ := product.NewProductName("Updated Product")
	p.UpdateName(newName)
	assert.Equal(t, newName, p.Name(), "Product name should be updated")
	assert.True(t, p.UpdatedAt().After(initialUpdateTime), "Update time should be updated after name change")

	// Test updating description
	newDesc, _ := product.NewProductDescription("This is an updated product")
	p.UpdateDescription(newDesc)
	assert.Equal(t, newDesc, p.Description(), "Product description should be updated")

	// Test updating price
	newPrice, _ := product.NewPrice(2000, "USD")
	p.UpdatePrice(newPrice)
	assert.Equal(t, newPrice, p.Price(), "Product price should be updated")

	// Test updating stock
	newStock := product.NewStock(20)
	p.UpdateStock(newStock)
	assert.Equal(t, newStock, p.Stock(), "Product stock should be updated")

	// Test decreasing stock
	err = p.DecreaseStock(5)
	assert.NoError(t, err, "Should be able to decrease stock")
	assert.Equal(t, uint(15), p.Stock().Quantity(), "Stock quantity should be decreased")

	// Test decreasing stock with insufficient quantity
	err = p.DecreaseStock(20)
	assert.Error(t, err, "Should error when decreasing stock with insufficient quantity")

	// Test increasing stock
	p.IncreaseStock(5)
	assert.Equal(t, uint(20), p.Stock().Quantity(), "Stock quantity should be increased")
}

func TestProductCategoryMethods(t *testing.T) {
	// Create a valid product
	productID, _ := product.NewProductID("prod-123")
	name, _ := product.NewProductName("Test Product")
	desc, _ := product.NewProductDescription("This is a test product")
	price, _ := product.NewPrice(1000, "USD")
	stock := product.NewStock(10)

	p, err := product.NewProduct(productID, name, desc, price, stock)
	require.NoError(t, err, "Failed to create product")

	// Create a category
	categoryID, _ := product.NewCategoryID("cat-1")
	categoryName, _ := product.NewCategoryName("Electronics")
	category, err := product.NewCategory(categoryID, categoryName)
	require.NoError(t, err, "Failed to create category")

	// Test adding a category
	p.AddCategory(category)
	assert.Len(t, p.Categories(), 1, "Product should have one category")
	assert.Equal(t, category, p.Categories()[0], "Category should be added to product")

	// Test adding the same category again (should not duplicate)
	p.AddCategory(category)
	assert.Len(t, p.Categories(), 1, "Product should still have one category")

	// Test has category
	assert.True(t, p.HasCategory(categoryID), "Product should have the category")

	// Test removing a category
	p.RemoveCategory(categoryID)
	assert.Len(t, p.Categories(), 0, "Product should have no categories")
	assert.False(t, p.HasCategory(categoryID), "Product should not have the category")

	// Test removing a non-existent category (should not error)
	nonExistentID, _ := product.NewCategoryID("cat-2")
	p.RemoveCategory(nonExistentID)
	assert.Len(t, p.Categories(), 0, "Product should still have no categories")
}
