package bootstrap

import (
	"database/sql"
	"net/http"

	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/handler"
	"io.github.com/api-product/internal/repository"
	"io.github.com/api-product/internal/routes"
	"io.github.com/api-product/internal/service"
)

func InjectionDependency(sqlConn *sql.DB) http.Handler {
	mux := http.NewServeMux()

	queries := db.New(sqlConn)                                  // *sql.DB -> *db.Queries (sqlc)
	productRepo := repository.NewProductRepository(queries)     // injeta queries no repository
	productService := service.NewProductService(productRepo)    // injeta repository no service
	productHandler := handler.NewProductHandler(productService) // injeta service no handler
	routes.RegisterProductRoutes(mux, productHandler)           // registra as rotas

	return mux
}
