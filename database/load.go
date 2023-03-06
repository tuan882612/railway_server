package database

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func LoadData() {
	client := GetClient()
	defer client.Close(context.Background())
	
	booked, err := os.Open("./data/booked.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file:", err)
	}
	defer booked.Close()

	passenger, err := os.Open("./data/passenger.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file:", err)
	}
	defer passenger.Close()

	train_status, err := os.Open("./data/train_status.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file:", err)
	}
	defer train_status.Close()

	train, err := os.Open("./data/train.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file:", err)
	}
	defer train.Close()

	copyCount, err := client.CopyFrom(
		context.Background(),
		pgx.Identifier{"my_table"},
		[]string{"column1", "column2", "column3"}, 
		csv.NewReader(booked),
	)
	if err != nil {
		log.Fatal("Unable to copy data:", err)
	}
	log.Printf("Copied %d rows to table", copyCount)
}