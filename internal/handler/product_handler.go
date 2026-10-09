package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	product, err := h.service.CreateProduct(r.Context(), req.ProductName, req.Quantity, req.Price, &req.Description)
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
	json.NewEncoder(w).Encode(dto.ProductResponseFromDomain(product))
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id não encontrado",
		})
		return
	}

	var req dto.UpdateProduct

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "campos invalidos",
		})
		return
	}

	product, err := h.service.UpdateProduct(r.Context(), int32(id), req.ProductName, req.Quantity, req.Price, &req.Description)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "impossivel processar a requisicao no momento",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.ProductResponseFromDomain(product))
}

func (h *ProductHandler) ProductById(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id nao encontrado",
		})
		return
	}

	product, err := h.service.ProductById(r.Context(), int32(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id nao encontrado",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.ProductResponseFromDomain(product))
}

func (h *ProductHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.AllProducts(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro inesperado",
		})
		return
	}

	response := make([]dto.CreateProductResponse, 0, len(products))

	for _, product := range products {
		response = append(response, dto.ProductResponseFromDomain(product))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id nao encontrado",
		})
		return
	}

	if err := h.service.DeleteProduct(r.Context(), int32(id)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao excluir o produto",
		})
	}

	w.WriteHeader(http.StatusNoContent)
}
