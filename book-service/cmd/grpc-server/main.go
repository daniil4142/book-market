package main

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	grpc_category_service "github.com/daniil4142/book-market/category-service/pkg/category-service"

	"github.com/daniil4142/book-market/book-service/internal/config"
	"github.com/daniil4142/book-market/book-service/internal/pkg/db"
	mwclient "github.com/daniil4142/book-market/book-service/internal/pkg/mw/client"
	"github.com/daniil4142/book-market/book-service/internal/server"
	book_service "github.com/daniil4142/book-market/book-service/internal/service/book"
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

	categoryServiceConn, err := grpc.NewClient(
		cfg.CategoryServiceAddr,
		grpc.WithUnaryInterceptor(mwclient.AddAppInfoUnary),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to create client")
	}

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	categoryServiceClient := grpc_category_service.NewCategoryServiceClient(categoryServiceConn)

	bookService := book_service.NewService(categoryServiceClient, conn)

	if err := server.NewGrpcServer(bookService).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
