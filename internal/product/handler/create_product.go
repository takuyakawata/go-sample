package handler

import (
	"encoding/json"
	"net/http"
	"sago-sample/api"

	product "sago-sample/internal/product/usecase"
)

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Currency    string `json:"currency"`
	Stock       uint   `json:"stock"`
}

// CreateProduct handles the creation of a new product
func (h *api.ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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
