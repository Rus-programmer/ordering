package main

import (
	"context"
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
