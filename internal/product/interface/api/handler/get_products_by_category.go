package handler

import (
	"net/http"
	"strings"

	product "sago-sample/internal/product/usecase"
)

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
