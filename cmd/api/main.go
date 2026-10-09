package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"io.github.com/api-product/internal/bootstrap"
	"io.github.com/api-product/internal/config"
	"io.github.com/api-product/internal/connection"
)

func main() {

	ctx := context.Background()

	cfg := config.LoadConfig()

	fmt.Println("Inicializando Servidor", cfg)

	if err := connection.MigrateDatabase(); err != nil {
		log.Fatalf("Failed to instance database %v", err)
	}

	sqlConn, err := connection.OpenConnection(ctx, connection.NewConnectConfig())
	if err != nil {
		log.Fatalf("Failed to open connection %v", err)
	}

	defer sqlConn.Close()

	dependencies := bootstrap.InjectionDependency(sqlConn)

	port := cfg.ServicePort

	address := ":" + port
	log.Println("Server executing on port ", address)

	log.Fatal(http.ListenAndServe(address, dependencies))

}
