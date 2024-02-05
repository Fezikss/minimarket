package main

import (
	"context"
	"log"


	"main.go/api"
	"main.go/config"
	"main.go/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(context.Background(),cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}
	defer store.Close()

	server := api.New(store)

	if err = server.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
