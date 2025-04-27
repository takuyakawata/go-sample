package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"

	product "sago-sample/feature/product/usecase"
)

func (h ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")

	if categoryID == "" {
		respondWithError(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	input := usecase.GetProductsByCategoryInput{
		CategoryID: categoryID,
	}

	// --- ↓↓↓ 以下は元のコードとほぼ同じ ---
	output, err := h.getProductsByCategoryUseCase.Execute(r.Context(), input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []ProductResponse

	respondWithJSON(w, http.StatusOK, response)
}
