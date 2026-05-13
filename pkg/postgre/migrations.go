package postgre

import (
	"e-commerce-api/internal/constants"
	"e-commerce-api/internal/infrastructure/config"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg config.DBConfig) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&&search_path=public",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

	m, err := migrate.New("file://"+constants.PathOfMigrations, dbURL)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return nil
		}
		return fmt.Errorf("could not apply migrations: %w", err)
	}

	return nil
}
