package main

import (
	"log"
	"ordering/api"
	util "ordering/utils"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Can't create server", err)
	}
	log.Println("config.HTTPServerAddress", config.HTTPServerAddress)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Can't start server", err)
	}
}
