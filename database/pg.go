package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func GetClient() *pgx.Conn {
	db_url := os.Getenv("db_url")
	ctx := context.Background()

	client, err := pgx.Connect(ctx, db_url)
	
	if err != nil {
		log.Fatal("Database connection failed.")
	}

	return client
}