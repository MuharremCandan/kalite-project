package main

import (
	"context"
	"go-backend-test/pkg/api"
	"go-backend-test/pkg/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := api.NewServer(&config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.StartServer(ctx); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
