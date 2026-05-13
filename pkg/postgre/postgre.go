package postgre

import (
	"context"
	"e-commerce-api/internal/constants"
	"e-commerce-api/internal/infrastructure/config"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(cfg config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&&search_path=public",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	poolConfig.MaxConns = constants.MaxDBPoolconns
	poolConfig.MinConns = constants.MinDBPoolconns

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
