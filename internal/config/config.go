package config

type Config struct {
	PostgresURL   string
	MigrationPath string
}

func LoadConfig() *Config {
	return &Config{
		PostgresURL:   "",
		MigrationPath: "",
	}
}
