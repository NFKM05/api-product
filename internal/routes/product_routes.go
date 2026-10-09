package routes

import (
	"net/http"

	"io.github.com/api-product/internal/handler"
)

type ProductRoute struct {
	handler handler.ProductHandler
}

func RegisterProductRoutes(mux *http.ServeMux, h *handler.ProductHandler) {
	mux.HandleFunc("POST /products", h.Create)
	mux.HandleFunc("PUT /products/{id}", h.Update)
	mux.HandleFunc("GET /products/{id}", h.ProductById)
	mux.HandleFunc("GET /products", h.ListAll)
	mux.HandleFunc("DELETE /products/{id}", h.Delete)
}
