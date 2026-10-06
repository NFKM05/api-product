package handler

import (
	"encoding/json"
	"net/http"

	"io.github.com/api-product/internal/dto"
	"io.github.com/api-product/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{service: svc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProduct

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "body fora do formato esperado",
		})
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req.ProductName, req.Quantity, req.Price, req.Description)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro inesperado durante a criação",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.ProductResponseToDomain(product))
}
