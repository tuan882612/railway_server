package train

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

func (r Repository) GetAllTrains() []Train {
	ctx := context.Background()
	data, err := r.db.Query(ctx, "SELECT * FROM Train")

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []Train

	for data.Next() {
		var train Train
		ValidateTrain(data, &train)
		arr = append(arr, train)
	}

	return arr
}

func (r Repository) GetInfoAgeRange(start string, end string) []JoinedInfo {
	query := `
	SELECT
		passenger.first_name, 
		passenger.last_name,
		train.train_no,
		train.name,
		train.source,
		train.destination,
		booked.status,
		booked.category
	FROM booked
	JOIN train
	on train.train_no = booked.train_no
	JOIN passenger
	on passenger.ssn = booked.ssn
	WHERE 
		2023-EXTRACT(YEAR FROM CAST(passenger.bdate AS DATE)) >= `+start+` AND
		2023-EXTRACT(YEAR FROM CAST(passenger.bdate AS DATE)) <= `+end+`;`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []JoinedInfo

	for data.Next() {
		var info JoinedInfo

		if err := data.Scan(
			&info.FirstName, &info.LastName,
			&info.TrainNo,&info.Name, &info.Source, &info.Destination,
			&info.Status, &info.Category,
		); err != nil {
			log.Fatal("Unable to process data.")
		}

		arr = append(arr, info)
	}

	return arr
}

func (r Repository) GetTrainInfoSum(name string) []TrainInfo {
	query := `
	SELECT 
		train_status.name,
		train_status.date,
		train_status.ps_occupied+train_status.gs_occupied
	FROM train_status
	WHERE train_status.name = '`+name+`';`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []TrainInfo

	for data.Next() {
		var info TrainInfo

		if err := data.Scan(&info.Name, &info.Date, &info.Count); err != nil {
			log.Fatal("Unable to process data.")
		}

		arr = append(arr, info)
	}

	return arr
}
