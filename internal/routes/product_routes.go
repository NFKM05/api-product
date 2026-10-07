package routes

import (
	"net/http"

	"io.github.com/api-product/internal/handler"
)

type ProductRoute struct {
	handler handler.ProductHandler
}

func NewRoleRoutes(route ProductRoute) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", route.handler.Create)
	mux.HandleFunc("PUT /products/{id}", route.handler.Update)
	mux.HandleFunc("GET /products/{id}", route.handler.ProductById)
	mux.HandleFunc("GET /products", route.handler.ListAll)
	mux.HandleFunc("DELETE /products/{id}", route.handler.Delete)

	return mux
}
