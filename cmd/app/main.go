package main

import (
	"fmt"
	"log"
	"net/http"
	"sago-sample/api"
	product "sago-sample/feature/product/domain"
	"sago-sample/feature/product/infrastructure"
	productUseCase "sago-sample/feature/product/usecase"
)

func main() {
	// Create repositories
	productRepo := infrastructure.NewProductRepository()

	// Create domain services
	productService := product.NewService(productRepo)

	// Create use cases
	createProductUseCase := productUseCase.NewCreateProductUseCase(productService)
	updateProductUseCase := productUseCase.NewUpdateProductUseCase(productService)
	deleteProductUseCase := productUseCase.NewDeleteProductUseCase(productService)
	getProductByIDUseCase := productUseCase.NewGetProductByIDUseCase(productService)
	getAllProductsUseCase := productUseCase.NewGetAllProductsUseCase(productService)

	// Create category-related use cases
	addCategoryToProductUseCase := productUseCase.NewAddCategoryToProductUseCase(productService)
	removeCategoryFromProductUseCase := productUseCase.NewRemoveCategoryFromProductUseCase(productService)
	getProductsByCategoryUseCase := productUseCase.NewGetProductsByCategoryUseCase(productService)

	// Create handlers
	productHandler := api.NewProductHandler(
		createProductUseCase,
		updateProductUseCase,
		deleteProductUseCase,
		getProductByIDUseCase,
		getAllProductsUseCase,
		addCategoryToProductUseCase,
		removeCategoryFromProductUseCase,
		getProductsByCategoryUseCase,
	)

	// Create router
	mux := http.NewServeMux()

	// Register routes
	productHandler.RegisterRoutes(mux)

	// Start server
	port := 8080
	fmt.Printf("Server running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
