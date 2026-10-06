package config

type Config struct {
	PostgresURL   string
	MigrationPath string
}

func LoadConfig() error {
	return nil
}
