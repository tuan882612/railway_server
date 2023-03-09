package database

import (
	"context"
	"log"
)

func CreateTables() {
	client := GetClient()
	tables := []string{
		`CREATE TABLE Passenger (
			ssn INTEGER PRIMARY KEY,
			first_name VARCHAR(50),
			last_name VARCHAR(50),
			address VARCHAR(255),
			city VARCHAR(50),
			county VARCHAR(50),
			phone2 VARCHAR(15),
			bdate VARCHAR(10)
		);`,
		`CREATE TABLE Train (
			train_no INTEGER PRIMARY KEY,
			name VARCHAR(50),
			premium_fare INTEGER,
			general_fare INTEGER,
			source VARCHAR(50),
			destination VARCHAR(50),
			available VARCHAR(255)
		);`,
		`CREATE TABLE Train_Status (
			date date,
			name VARCHAR(50) PRIMARY KEY,
			ps_available integer,
			gs_available integer,
			ps_occupied integer,
			gs_occupied integer, 
			FOREIGN KEY (name) REFERENCES Train (name)
		);`,`
		CREATE TABLE Booked (
			ssn INTEGER,
			train_no INTEGER,
			category VARCHAR(10),
			status VARCHAR(10),
			FOREIGN KEY (ssn) REFERENCES Passenger (ssn),
			FOREIGN KEY (train_no) REFERENCES Train (train_no)
		);`,
	}

	for _, table := range tables {
		_, err := client.Exec(
			context.Background(), 
			table,
		)
		if err != nil {
			log.Fatal("Unable to create table:", err)
		}
	}

}