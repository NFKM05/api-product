package config

import (
	"bufio"
	"os"
	"strings"
)

// Struct que contém todos os dados de configuração da aplicação a partir de variáveis de amb.
type Config struct {
	PostgresURL   string
	MigrationPath string
	ServicePort   string
}

// Função que inicializa e carrega as variáveis de ambiente do arquivo .env e monta a config
func LoadConfig() *Config {

	loadEnv()

	return &Config{
		PostgresURL:   getEnv("DATABASE_URL", "postgresql://..."),
		MigrationPath: getEnv("MIGRATION_PATH", "./migrations"),
		ServicePort:   getEnv("SERVICE_PORT", "8080"),
	}
}

// Função que lê o arquivo .env linha a linha, e injeta os pares de chave=valor
func loadEnv() {
	file, err := os.Open(".env.example")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scanner lê o arquivo linha por linha, sem carregar tudo na
	// memória de uma vez.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Separa só no PRIMEIRO "=" (limite 2). Isso é essencial:
		// uma DATABASE_URL tem várias "=" dentro (query params do
		// Postgres), e SplitN com limite garante que elas fiquem
		// intactas dentro do valor, em vez de serem cortadas.
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove aspas duplas envolvendo o valor, se existirem.
		// Permite escrever DATABASE_URL="postgres://..." no arquivo.
		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) && len(value) >= 2 {
			value = value[1 : len(value)-1]
		}

		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

}

func getEnv(key, content string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return content
}
