package handler

import (
	"encoding/json"
	"net/http"
	handler2 "sago-sample/feature/handler"
	product "sago-sample/feature/product/usecase"
)

type CreateProductHandler struct {
	UseCase *product.CreateProductUseCase
}

func NewCreateProductHandler(uc *product.CreateProductUseCase) *CreateProductHandler {
	return &CreateProductHandler{UseCase: uc}
}

func (h *CreateProductHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var in product.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler2.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	out, err := h.UseCase.Execute(r.Context(), in)
	if err != nil {
		handler2.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler2.respondWithJSON(w, http.StatusCreated, out)
}
