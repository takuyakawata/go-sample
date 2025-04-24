package handler

import (
	"net/http"
	"strings"

	product "sago-sample/internal/product/usecase"
)

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
