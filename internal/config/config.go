package config

import (
	"os"
)

type Config struct {
	PostgresURL   string
	MigrationPath string
}

func LoadConfig() *Config {

	loadEnv()

	return &Config{
		PostgresURL:   getEnv("DATABASE_URL", "postgresql://..."),
		MigrationPath: getEnv("MIGRATION_PATH", "./migrations"),
	}
}

func loadEnv() {
	file, err := os.Open(".env.example")
	if err != nil {
		return
	}

	defer file.Close()

}

func getEnv(key, content string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return content
}
