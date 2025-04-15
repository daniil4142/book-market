package main

import (
	"log"

	"github.com/daniil4142/book-market/category-service/internal/config"
	"github.com/daniil4142/book-market/category-service/internal/pkg/db"
	"github.com/daniil4142/book-market/category-service/migrations"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatalf("config.ReadConfigYML(): %v", err)
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatalf("sql.Open() error: %v", err)
	}
	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(conn.DB, ".")
	if err != nil {
		log.Fatalf("goose.Up(): %v", err)
	}
}
