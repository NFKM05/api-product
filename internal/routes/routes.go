package routes

import (
	"errors"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"io.github.com/api-product/internal/handler"
	"io.github.com/api-product/internal/middleware"
	"io.github.com/api-product/internal/utils"
)

func RegisterProductRoutes(mux *http.ServeMux, h *handler.ProductHandler, jwtService *utils.JwtService) {
	auth := middleware.Auth(jwtService)

	mux.Handle("POST /products", auth(http.HandlerFunc(h.Create)))
	mux.Handle("PUT /products/{id}", auth(http.HandlerFunc(h.Update)))
	mux.Handle("GET /products/{id}", auth(http.HandlerFunc(h.ProductById)))
	mux.Handle("GET /products", auth(http.HandlerFunc(h.ListAll)))
	mux.Handle("DELETE /products/{id}", auth(http.HandlerFunc(h.Delete)))
	mux.HandleFunc("/", GetNotFound)
}

func RegisterUserRoutes(mux *http.ServeMux, h *handler.UserHandler, jwtService *utils.JwtService) {
	auth := middleware.Auth(jwtService)

	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.Handle("PATCH /users/{id}/change-name", auth(http.HandlerFunc(h.UpdateName)))
	mux.Handle("PATCH /users/{id}/change-email", auth(http.HandlerFunc(h.UpdateEmail)))
	mux.Handle("PATCH /users/{id}/change-password", auth(http.HandlerFunc(h.UpdatePassword)))
	mux.Handle("PATCH /users/{id}/desactivate", auth(http.HandlerFunc(h.DesactivateUser)))
	mux.Handle("GET /users", auth(http.HandlerFunc(h.ListAllUsers)))
	mux.Handle("GET /users/{id}", auth(http.HandlerFunc(h.FindUserById)))
}

func RegisterSwaggerRoutes(mux *http.ServeMux) {
	mux.Handle("/swagger", httpSwagger.WrapHandler)
}

func GetNotFound(w http.ResponseWriter, r *http.Request) {
	utils.NewApiError(w, http.StatusNotFound, errors.New("the requested resource was not found"))
}
