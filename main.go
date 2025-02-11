package main

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ordering/api"
	db "ordering/db/sqlc"
	_ "ordering/docs"
	"ordering/middleware"
	"ordering/services"
	"ordering/token"
	"ordering/util"
	"os"
)

// @title Ordering project
// @description This project is a backend service designed for managing orders
// @version 1.0
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msgf("Can't create token maker")
	}

	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't connect to database")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(pool)
	newMiddleware := middleware.NewMiddleware(store, tokenMaker)
	newService := services.NewService(config, store, tokenMaker)

	server, err := api.NewServer(config, newMiddleware, newService)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't start server")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
