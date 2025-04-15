package main

import (
	"context"

	"github.com/daniil4142/book-market/book-service/internal/config"
	"github.com/daniil4142/book-market/book-service/internal/pkg/db"
	"github.com/daniil4142/book-market/book-service/migrations"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	const cmd = "up"

	err = goose.RunContext(ctx, cmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}

}
