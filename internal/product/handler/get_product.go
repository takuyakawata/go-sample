package handler

import (
	"net/http"
	"strings"

	usecase "sago-sample/internal/product/usecase"
)

func (h *api.ProductHandler) NewGetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	input := usecase.GetProductInput{
		ID: id,
	}

	output, err := h.getProductUseCase.Execute(r.Context(), input)
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
