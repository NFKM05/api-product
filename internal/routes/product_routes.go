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

	return mux
}
