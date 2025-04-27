package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	product "sago-sample/feature/product/usecase"
)

type UpdateProductHandler struct {
	UseCase *product.UpdateProductUseCase
}

func NewUpdateProductHandler(uc *product.UpdateProductUseCase) *UpdateProductHandler {
	return &UpdateProductHandler{UseCase: uc}
}

func (h *UpdateProductHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var in product.UpdateProductInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Set the ID from the URL parameter
	in.ID = id

	out, err := h.UseCase.Execute(r.Context(), in)
	if err != nil {
		if err.Error() == "product not found" {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, out)
}
