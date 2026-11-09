package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "io.github.com/api-product/docs"
	"io.github.com/api-product/internal/bootstrap"
	"io.github.com/api-product/internal/config"
	"io.github.com/api-product/internal/connection"
)

// @title api-product
// @version 1.0
// @description Api de cadastro de produtos com autenticação de usuários
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Token de acesso no formato Bearer {token}

func main() {

	// Contexto inicial do sistema
	ctx := context.Background()

	// Carrega o arquivo com as variáveis de ambiente e monta a configuração da aplicação
	cfg := config.LoadConfig()

	fmt.Println("Inicializando Servidor", cfg)

	// Aplica todas as Migrations pendentes no banco
	if err := connection.MigrateDatabase(); err != nil {
		log.Fatalf("Failed to instance database %v", err)
	}

	// Abre o pool de conexões com o banco de dados que a aplicação vai utilizar durante a execução
	sqlConn, err := connection.OpenConnection(ctx, connection.NewConnectConfig())
	if err != nil {
		log.Fatalf("Failed to open connection %v", err)
	}

	// Fecha a pool o mais tarde possível
	defer sqlConn.Close()

	// Injetando todas as dependências entre as camadas (Repository -> Service -> Handler)
	dependencies := bootstrap.InjectionDependency(sqlConn)

	// Pegando o valor de variavel de ambiente da porta da aplicação
	port := cfg.ServicePort

	address := ":" + port
	log.Println("Server executing on port ", address)

	// Inicializando o servidor HTTP ou ocorre um erro fatal ao tentar subir a aplicação
	log.Fatal(http.ListenAndServe(address, dependencies))

}
