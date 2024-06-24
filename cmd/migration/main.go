package main

import (
	"context"
	"go-grpc-connect/config"
	"go-grpc-connect/ent"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dbConfig := config.LoadDBConfig()
	client, err := ent.Open(config.DRIVER_NAME, dbConfig.ConnectionString())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
