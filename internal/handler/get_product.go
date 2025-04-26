package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	product "sago-sample/internal/product/usecase"
)

type GetProductHandler struct {
	UseCase       *product.GetProductUseCase
	GetAllUseCase *product.GetAllProductsUseCase
}

func NewGetProductHandler(uc *product.GetProductUseCase, getAllUc *product.GetAllProductsUseCase) *GetProductHandler {
	return &GetProductHandler{UseCase: uc, GetAllUseCase: getAllUc}
}

func (h *GetProductHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllUseCase.Execute(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, output.Products)
}

func (h *GetProductHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	out, err := h.UseCase.Execute(r.Context(), product.GetProductInput{ID: id})
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

// Helper functions for JSON responses
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
