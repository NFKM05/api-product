package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"io.github.com/api-product/internal/dto"
	"io.github.com/api-product/internal/middleware"
	"io.github.com/api-product/internal/service"
	"io.github.com/api-product/internal/utils"
)

// Struct iniciliaza e recebe o ponteiro da camada service
type ProductHandler struct {
	service *service.ProductService
}

// Realiza a injeção de dependencia da camada
func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{service: svc}
}

// Cria um novo produto
// CreateProduct godoc
// @Summary criação de um novo produto
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateProduct true "Dados de criação do produto"
// @Success 201 {object} dto.CreateProductResponse
// @Failure 400 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Router /products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProduct

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req.ProductName, req.Quantity, req.Price, &req.Description, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected error during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, dto.ProductResponseFromDomain(product))
}

// Atualiza um produto existente
// UpdateProduct godoc
// @Summary Atualizar os dados de um produto existente
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do produto"
// @Param request body dto.UpdateProduct true "Dados do produto"
// @Success 200 {object} dto.CreateProductResponse
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 403 {object} utils.ApiError
// @Router /products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	var req dto.UpdateProduct

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	product, err := h.service.UpdateProduct(r.Context(), int32(id), req.ProductName, req.Quantity, req.Price, &req.Description, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.ProductResponseFromDomain(product))
}

// Busca um produto por um identificador
// ProductById godoc
// @Summary Pesquisa um produto pelo seu identificador
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do produto"
// @Success 200 {object} dto.CreateProductResponse
// @Failure 400 {object} utils.ApiError
// @Failure 404 {object} utils.ApiError
// @Router /products/{id} [get]
func (h *ProductHandler) ProductById(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	product, err := h.service.ProductById(r.Context(), int32(id))
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.ProductResponseFromDomain(product))
}

// Lista todos os produtos cadastrados
// ListAllProducts godoc
// @Summary Lista todos os produtos cadastrados na base
// @Tags products
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.CreateProductResponse
// @Failure 500 {object} utils.ApiError
// @Router /products [get]
func (h *ProductHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.AllProducts(r.Context())
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected error during the process"))
		return
	}

	response := make([]dto.CreateProductResponse, 0, len(products))

	for _, product := range products {
		response = append(response, dto.ProductResponseFromDomain(product))
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// Exclui um produto cadastrado
// DeleteProduct godoc
// @Summary Exclui um produto da base
// @Tags products
// @Security BearerAuth
// @Param id path int true "Id do produto"
// @Success 204
// @Failure 404 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 403 {object} utils.ApiError
// @Router /products/{id} [delete]

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	if err := h.service.DeleteProduct(r.Context(), int32(id), tokenData.UserID); err != nil {
		utils.NewApiError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
