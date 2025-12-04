package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	
	// Connection pool settings
	MaxConns     int32
	MinConns     int32
	MaxConnLife  time.Duration
	MaxConnIdle  time.Duration
}

func NewDBPool(ctx context.Context, cfg DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	// parse config
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %v", err)
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v\n", err)
	}

	return dbpool, nil
}

func Close(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
	}
}
