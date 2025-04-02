package main

import (
	"context"
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/daniil4142/book-market//book-service/internal/config"
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

	bookServiceConn, err := grpc.DialContext(
		context.Background(),
		cfg.BookServiceAddr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(mwclient.AddAppInfoUnary),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to create client")
	}

	bookServiceClient := grpc_book_service.NewCategoryServiceClient(bookServiceConn)

	productService := book_service.NewService(categoryServiceClient)

	if err := server.NewGrpcServer(productService).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
