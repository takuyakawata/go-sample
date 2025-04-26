package handler

import (
	"encoding/json"
	"net/http"
	"sago-sample/api"
	"strings"

	product "sago-sample/internal/product/usecase"
)

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Currency    string `json:"currency"`
	Stock       uint   `json:"stock"`
}

// UpdateProduct handles the update of an existing product
func (h *api.ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
