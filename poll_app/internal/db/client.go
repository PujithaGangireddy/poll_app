package db

import (
	"context"
	"fmt"
	"poll_app/ent"
	"poll_app/ent/migrate"

	_ "github.com/lib/pq"
)

func NewClient() (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "polling_app")

	client, err := ent.Open("postgres", dsn)
	fmt.Println(client, err)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		return nil, err
	}

	return client, nil
}
