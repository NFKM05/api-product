package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/dto"
	"io.github.com/api-product/internal/middleware"
	"io.github.com/api-product/internal/service"
	"io.github.com/api-product/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}

// CreateUser godoc
// @Summary Criar usuário
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "Dados de criação de usuário"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	user, err := h.service.Create(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, dto.UserResponseFromDomain(user))
}

// Login godoc
// @Summary Realizar login para obter o Bearer token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Dados de login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	token, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.LoginResponse{Token: token})
}

// UpdateName godoc
// @Summary Atualizar o nome do usuário cadastrado
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do usuário"
// @Param request body dto.UpdateNameRequest true "Dados do usuário"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Router /users/{id}/change-name [patch]
func (h *UserHandler) UpdateName(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)

	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	var req dto.UpdateNameRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	user, err := h.service.UpdateName(r.Context(), id, req.Name, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

// UpdateEmail godoc
// @Summary Atualizar o email do cadastro do usuário
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do usuário"
// @Param request body dto.UpdateEmailRequest true "Dados do usuário"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 403 {object} utils.ApiError
// @Router /users/{id}/change-email [patch]
func (h *UserHandler) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	var req dto.UpdateEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	user, err := h.service.UpdateEmail(r.Context(), id, req.Email, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

// UpdatePassword godoc
// @Summary Atualiza a senha de um usuário
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do usuário"
// @Param request body dto.UpdatePasswordRequest true "Dados do usuário"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Router /users/{id}/change-password [patch]
func (h *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	var req dto.UpdatePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	user, err := h.service.UpdatePassword(r.Context(), id, req.Password, req.NewPassword, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

// ListAllUsers godoc
// @Summary Lista todos os usuários cadastrados na base
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} utils.ApiError
// @Router /users [get]
func (h *UserHandler) ListAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.ListAllUsers(r.Context())
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	response := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, dto.UserResponseFromDomain(user))
	}

	utils.WriteJSON(w, http.StatusOK, response)

}

// FindUserById godoc
// @Summary Lista usuários por identificador
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do usuário"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Router /users/{id} [get]
func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	user, err := h.service.FindUserById(r.Context(), id)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

// DesactivateUser godoc
// @Summary
// @Tags
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id do usuário"
// @Success 200 {object} dto.DesactivateUser
// @Failure 404 {object} utils.ApiError
// @Failure 401 {object} utils.ApiError
// @Failure 403 {object} utils.ApiError
// @Router /users/{id}/desactivate [patch]
func (h *UserHandler) DesactivateUser(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	tokenData := middleware.GetTokenData(r)
	if tokenData == nil {
		utils.NewApiError(w, http.StatusUnauthorized, errors.New("must be authenticated to access the resource"))
		return
	}

	user, err := h.service.DesactivateUser(r.Context(), id, tokenData.UserID)
	if err != nil {
		utils.NewApiError(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))

}
