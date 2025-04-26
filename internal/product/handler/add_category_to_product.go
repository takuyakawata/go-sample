package handler

import (
	"encoding/json"
	"net/http"
	"sago-sample/api"
	"strings"

	product "sago-sample/internal/product/usecase"
)

// AddCategoryToProductRequest represents the request body for adding a category to a product
type AddCategoryToProductRequest struct {
	CategoryID   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// AddCategoryToProduct handles adding a category to a product
func (h *api.ProductHandler) AddCategoryToProduct(w http.ResponseWriter, r *http.Request) {
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
