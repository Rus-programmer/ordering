package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"ordering/api"
	db "ordering/db/sqlc"
	util "ordering/utils"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	store := db.NewStore(pool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can't create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Can't start server", err)
	}
}
