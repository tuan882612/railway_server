package validators

import (
	"context"
	"log"
	"main/tools/config"
	"time"
)

func ValidatePostgres() bool {
	client := config.GetClient()

	query := "SELECT NOW()"
	ctx := context.Background()

	var example time.Time
	err := client.QueryRow(ctx, query).Scan(&example)

	if err != nil {
		log.Fatal("Unable to connect to Postgres.", err)
		return false
	}
	return true
}