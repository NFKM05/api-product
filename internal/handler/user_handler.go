package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/dto"
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

	user, err := h.service.UpdateName(r.Context(), id, req.Name)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

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

	user, err := h.service.UpdateEmail(r.Context(), id, req.Email)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))
}

/*func (h *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

}*/

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

func (h *UserHandler) DesactivateUser(w http.ResponseWriter, r *http.Request) {

	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		utils.NewApiError(w, http.StatusNotFound, errors.New("unable to locate the identifier"))
		return
	}

	var req dto.DesactivateUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewApiError(w, http.StatusBadRequest, errors.New("body does not match the expected format"))
		return
	}

	user, err := h.service.DesactivateUser(r.Context(), id)
	if err != nil {
		utils.NewApiError(w, http.StatusInternalServerError, errors.New("unespected erro during the process"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.UserResponseFromDomain(user))

}
