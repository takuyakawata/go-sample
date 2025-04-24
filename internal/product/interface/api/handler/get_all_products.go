package handler

import (
	"net/http"
)

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
