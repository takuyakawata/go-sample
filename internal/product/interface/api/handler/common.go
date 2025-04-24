package handler

import (
	"encoding/json"
	"net/http"
)

// CategoryResponse represents a category in the response
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProductResponse represents the response body for product operations
type ProductResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       uint               `json:"price"`
	Currency    string             `json:"currency"`
	Stock       uint               `json:"stock"`
	Categories  []CategoryResponse `json:"categories"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// respondWithError returns an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

// respondWithJSON returns a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
