package bootstrap

import (
	"database/sql"
	"net/http"

	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/handler"
	"io.github.com/api-product/internal/middleware"
	"io.github.com/api-product/internal/repository"
	"io.github.com/api-product/internal/routes"
	"io.github.com/api-product/internal/service"
	"io.github.com/api-product/internal/utils"
)

// Função que injeta, cria e agrupa todas as instâncias por camadas e fluxo de relação
func InjectionDependency(sqlConn *sql.DB) http.Handler {
	mux := http.NewServeMux()
	jwt := utils.NewJwtService()
	queries := db.New(sqlConn) // *sql.DB -> *db.Queries (sqlc)

	// Agrupa o fluxo de dependência de Produtos
	productRepo := repository.NewProductRepository(queries)     // Injeta queries no repository
	productService := service.NewProductService(productRepo)    // Injeta repository no service
	productHandler := handler.NewProductHandler(productService) // Injeta service no handler
	routes.RegisterProductRoutes(mux, productHandler, jwt)      // Registra as rotas

	// Agrupa o fluxo de dependência de Usuários
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo, jwt)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(mux, userHandler, jwt)

	// Documentação da API com Swagger
	routes.RegisterSwaggerRoutes(mux)

	// Log que audita toda rota chamada na aplicação antes de chegar ao Handler final
	return middleware.LoggerMiddleware(mux)
}
