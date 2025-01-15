package db

import (
	"context"
	"emb/pkg/db/ent"
	"log"
)

var Client *ent.Client

func Init() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=gotask2 password=admin sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	Client = client
}
