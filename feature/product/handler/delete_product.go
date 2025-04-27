package handler

import (
	"net/http"
	"sago-sample/api"
	"strings"

	product "sago-sample/feature/product/usecase"
)

// DeleteProduct handles the deletion of a product
func (h *api.ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
