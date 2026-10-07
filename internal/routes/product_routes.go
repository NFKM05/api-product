package routes

import (
	"net/http"

	"io.github.com/api-product/internal/handler"
)

type ProductRoute struct {
	handler handler.ProductHandler
}

func NewRoleRoutes(h ProductRoute) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", h.handler.Create)
	mux.HandleFunc("PUT /products/{id}", h.handler.Update)
	mux.HandleFunc("GET /products/{id}", h.handler.ProductById)
	mux.HandleFunc("GET /products", h.handler.ListAll)

	return mux
}
