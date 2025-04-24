package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Currency    string `json:"currency"`
	Stock       uint   `json:"stock"`
}

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Currency    string `json:"currency"`
	Stock       uint   `json:"stock"`
}

// CategoryResponse represents a category in the response
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProductResponse represents the response body for product operations
type ProductResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       uint               `json:"price"`
	Currency    string             `json:"currency"`
	Stock       uint               `json:"stock"`
	Categories  []CategoryResponse `json:"categories"`
}

// AddCategoryToProductRequest represents the request body for adding a category to a product
type AddCategoryToProductRequest struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
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

// CreateProduct handles the creation of a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	input := product.CreateProductInput{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
		Stock:       req.Stock,
	}

	output, err := h.createProductUseCase.Execute(r.Context(), input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Map categories
	categories := make([]CategoryResponse, 0)

	response := ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		Currency:    output.Currency,
		Stock:       output.Stock,
		Categories:  categories,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// UpdateProduct handles the update of an existing product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	input := product.UpdateProductInput{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Currency:    req.Currency,
		Stock:       req.Stock,
	}

	output, err := h.updateProductUseCase.Execute(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Map categories
	categories := make([]CategoryResponse, 0)

	response := ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		Currency:    output.Currency,
		Stock:       output.Stock,
		Categories:  categories,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// DeleteProduct handles the deletion of a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	input := product.DeleteProductInput{
		ID: id,
	}

	err := h.deleteProductUseCase.Execute(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetProductByID handles the retrieval of a product by ID
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	input := product.GetProductByIDInput{
		ID: id,
	}

	output, err := h.getProductByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Map categories
	categories := make([]CategoryResponse, 0, len(output.Categories))
	for _, c := range output.Categories {
		categories = append(categories, CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	response := ProductResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		Currency:    output.Currency,
		Stock:       output.Stock,
		Categories:  categories,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetAllProducts handles the retrieval of all products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	output, err := h.getAllProductsUseCase.Execute(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []ProductResponse
	for _, p := range output.Products {
		// Map categories
		categories := make([]CategoryResponse, 0, len(p.Categories))
		for _, c := range p.Categories {
			categories = append(categories, CategoryResponse{
				ID:   c.ID,
				Name: c.Name,
			})
		}

		response = append(response, ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Currency:    p.Currency,
			Stock:       p.Stock,
			Categories:  categories,
		})
	}

	respondWithJSON(w, http.StatusOK, response)
}

// respondWithError returns an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

// RemoveCategoryFromProduct handles removing a category from a product
func (h *ProductHandler) RemoveCategoryFromProduct(w http.ResponseWriter, r *http.Request) {
	// Extract product ID and category ID from URL
	path := strings.TrimPrefix(r.URL.Path, "/products/")
	parts := strings.Split(path, "/categories/")
	if len(parts) != 2 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	productID := parts[0]
	categoryID := parts[1]

	input := product.RemoveCategoryFromProductInput{
		ProductID:  productID,
		CategoryID: categoryID,
	}

	output, err := h.removeCategoryFromProductUseCase.Execute(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Map categories
	categories := make([]CategoryResponse, 0, len(output.Categories))
	for _, c := range output.Categories {
		categories = append(categories, CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	response := ProductResponse{
		ID:          output.ProductID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		Currency:    output.Currency,
		Stock:       output.Stock,
		Categories:  categories,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// AddCategoryToProduct handles adding a category to a product
func (h *ProductHandler) AddCategoryToProduct(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL
	productID := strings.TrimPrefix(r.URL.Path, "/products/")
	productID = strings.TrimSuffix(productID, "/categories")

	var req AddCategoryToProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	input := product.AddCategoryToProductInput{
		ProductID:    productID,
		CategoryID:   req.CategoryID,
		CategoryName: req.CategoryName,
	}

	output, err := h.addCategoryToProductUseCase.Execute(r.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Map categories
	categories := make([]CategoryResponse, 0, len(output.Categories))
	for _, c := range output.Categories {
		categories = append(categories, CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	response := ProductResponse{
		ID:          output.ProductID,
		Name:        output.Name,
		Description: output.Description,
		Price:       output.Price,
		Currency:    output.Currency,
		Stock:       output.Stock,
		Categories:  categories,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetProductsByCategory handles retrieving products by category
func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from URL
	categoryID := strings.TrimPrefix(r.URL.Path, "/categories/")
	categoryID = strings.TrimSuffix(categoryID, "/products")

	input := product.GetProductsByCategoryInput{
		CategoryID: categoryID,
	}

	output, err := h.getProductsByCategoryUseCase.Execute(r.Context(), input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []ProductResponse
	for _, p := range output.Products {
		// Map categories
		categories := make([]CategoryResponse, 0, len(p.Categories))
		for _, c := range p.Categories {
			categories = append(categories, CategoryResponse{
				ID:   c.ID,
				Name: c.Name,
			})
		}

		response = append(response, ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Currency:    p.Currency,
			Stock:       p.Stock,
			Categories:  categories,
		})
	}

	respondWithJSON(w, http.StatusOK, response)
}

// respondWithJSON returns a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
