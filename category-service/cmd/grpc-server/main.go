package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/daniil4142/book-market/category-service/internal/config"
	"github.com/daniil4142/book-market/category-service/internal/server"
	"github.com/daniil4142/book-market/category-service/internal/service/category"
	cat_repository "github.com/daniil4142/book-market/category-service/internal/service/category/repository"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	initCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(initCtx, cfg.Database.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	categoryRepository := cat_repository.New(db)
	categoryService := category.New(categoryRepository)

	if err := server.NewGrpcServer(
		categoryService,
	).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
