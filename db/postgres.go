package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arashn2y/go-server/libs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(ctx context.Context) (*sql.DB, error) {
	DSN := libs.GetEnv("DATABASE_URL")
	db, err := sql.Open("pgx", DSN)
	if err != nil {
		return nil, err
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return db, nil
}
