package handler

import (
	"net/http"

	product "sago-sample/internal/product/usecase"
)

// ProductHandler handles HTTP requests for product operations
type ProductHandler struct {
	createProductUseCase             *product.CreateProductUseCase
	updateProductUseCase             *product.UpdateProductUseCase
	deleteProductUseCase             *product.DeleteProductUseCase
	getProductByIDUseCase            *product.GetProductByIDUseCase
	getAllProductsUseCase            *product.GetAllProductsUseCase
	addCategoryToProductUseCase      *product.AddCategoryToProductUseCase
	removeCategoryFromProductUseCase *product.RemoveCategoryFromProductUseCase
	getProductsByCategoryUseCase     *product.GetProductsByCategoryUseCase
}

// NewProductHandler creates a new product handler
func NewProductHandler(
	createProductUseCase *product.CreateProductUseCase,
	updateProductUseCase *product.UpdateProductUseCase,
	deleteProductUseCase *product.DeleteProductUseCase,
	getProductByIDUseCase *product.GetProductByIDUseCase,
	getAllProductsUseCase *product.GetAllProductsUseCase,
	addCategoryToProductUseCase *product.AddCategoryToProductUseCase,
	removeCategoryFromProductUseCase *product.RemoveCategoryFromProductUseCase,
	getProductsByCategoryUseCase *product.GetProductsByCategoryUseCase,
) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:             createProductUseCase,
		updateProductUseCase:             updateProductUseCase,
		deleteProductUseCase:             deleteProductUseCase,
		getProductByIDUseCase:            getProductByIDUseCase,
		getAllProductsUseCase:            getAllProductsUseCase,
		addCategoryToProductUseCase:      addCategoryToProductUseCase,
		removeCategoryFromProductUseCase: removeCategoryFromProductUseCase,
		getProductsByCategoryUseCase:     getProductsByCategoryUseCase,
	}
}

// RegisterRoutes registers the product routes
func (h *ProductHandler) RegisterRoutes(mux *http.ServeMux) {
	// Product routes
	mux.HandleFunc("GET /products", h.GetAllProducts)
	mux.HandleFunc("GET /products/{id}", h.GetProductByID)
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", h.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", h.DeleteProduct)

	// Category routes
	mux.HandleFunc("POST /products/{id}/categories", h.AddCategoryToProduct)
	mux.HandleFunc("DELETE /products/{id}/categories/{categoryId}", h.RemoveCategoryFromProduct)
	mux.HandleFunc("GET /categories/{id}/products", h.GetProductsByCategory)
}
