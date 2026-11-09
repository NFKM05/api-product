package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"io.github.com/api-product/internal/handler"
	"io.github.com/api-product/internal/middleware"
	"io.github.com/api-product/internal/utils"
)

// Função que registra todas as rotas de Products no mux
func RegisterProductRoutes(mux *http.ServeMux, h *handler.ProductHandler, jwtService *utils.JwtService) {
	auth := middleware.Auth(jwtService)

	// Rotas que estão com a função auth() são manipuladas pelo Middleware para acesso
	mux.Handle("POST /products", auth(http.HandlerFunc(h.Create)))
	mux.Handle("PUT /products/{id}", auth(http.HandlerFunc(h.Update)))
	mux.Handle("GET /products/{id}", auth(http.HandlerFunc(h.ProductById)))
	mux.Handle("GET /products", auth(http.HandlerFunc(h.ListAll)))
	mux.Handle("DELETE /products/{id}", auth(http.HandlerFunc(h.Delete)))
}

// Função que registra todas as rotas de Users no mux
func RegisterUserRoutes(mux *http.ServeMux, h *handler.UserHandler, jwtService *utils.JwtService) {
	auth := middleware.Auth(jwtService)

	// Duas rotas públicas para fazer o acesso sem necessitar de autenticação
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("POST /users", h.CreateUser)

	// Rotas que estão com a função auth() são manipuladas pelo Middleware para acesso
	mux.Handle("PATCH /users/{id}/change-name", auth(http.HandlerFunc(h.UpdateName)))
	mux.Handle("PATCH /users/{id}/change-email", auth(http.HandlerFunc(h.UpdateEmail)))
	mux.Handle("PATCH /users/{id}/change-password", auth(http.HandlerFunc(h.UpdatePassword)))
	mux.Handle("PATCH /users/{id}/desactivate", auth(http.HandlerFunc(h.DesactivateUser)))
	mux.Handle("GET /users", auth(http.HandlerFunc(h.ListAllUsers)))
	mux.Handle("GET /users/{id}", auth(http.HandlerFunc(h.FindUserById)))
}

// Registra a rota da documentação da API no Swagger
func RegisterSwaggerRoutes(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
