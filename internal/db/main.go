package main

import (
	"context"

	"github.com/arashn0uri/go-server/internal/config"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	conn, err := pgx.Connect(ctx, cfg.DB.DSN)
	if err != nil {
		logrus.Fatal("failed to connect to the database", "error", err)
	} else {
		logrus.Info("connected to the database successfully")
	}
	defer conn.Close(ctx)

	db := repository.New(conn)

	if err := Seed(ctx, db); err != nil {
		logrus.Fatal("failed to seed the database", "error", err)
	}

	logrus.Info("database seeded successfully")
}
