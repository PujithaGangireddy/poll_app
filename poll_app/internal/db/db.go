package db

import (
	"context"
	"log"
	"os"
	"poll_app/ent"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewDBClient() (*ent.Client, error) {
	dbURL := os.Getenv("DATABASE_URL")
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	return client, nil
}
