package connection

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ConnectConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdle     time.Time
	ConnMaxLifetime time.Time
	ForceUTC        bool
}

func MigrateDatabase() error {
	databaseURL := os.Getenv("DATABASE_URL")
	migrationPath := os.Getenv("MIGRATION_PATH")

	if databaseURL == "" {
		return fmt.Errorf("Database URL is not available")
	}
	if migrationPath == "" {
		return fmt.Errorf("Migration Path is not available")
	}

	absPathMigration, err := filepath.Abs(migrationPath)
	if err != nil {
		return fmt.Errorf("Absolute migration path is not available: %w", err)
	}

	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("Failed to open connection with database: %w", err)
	}
	defer CloseConnection(conn)

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Failed to create database: %w", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://"+absPathMigration,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("Failed to create migrate instance: %w", err)
	}

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to apply migrations: %w", err)
	}

	return nil
}

func CloseConnection(conn *sql.DB) error {
	return conn.Close()
}

func OpenConnection(ctx context.Context, config ConnectConfig) (*sql.DB, error) {

	if config.DSN == "" {
		config.DSN = os.Getenv("DATABASE_URL")
	}

	if config.DSN == "" {
		return nil, fmt.Errorf("Database not provided and Postgres URL not set")
	}

	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = 25
	}

	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 25
	}

	db, err := sql.Open("postgres", config.DSN)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect database: %w", err)
	}

	return db, nil
}
