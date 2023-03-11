package passenger

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

func (r Repository) GetAllPassengers() []Passenger {
	ctx := context.Background()
	data, err := r.db.Query(ctx, "SELECT * from passenger;")

	return LoadData(data, err)
}

func (r Repository) GetWaiting() []Waiting {
	query := `
	SELECT 
		train.name, 
		passenger.*
	FROM passenger
	JOIN booked
	on passenger.ssn = booked.ssn
	JOIN train
	on train.train_no = booked.train_no
	WHERE booked.status = 'WaitL';`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []Waiting

	for data.Next() {
		var passenger Waiting
		ValidateWaiting(data, &passenger)
		arr = append(arr, passenger)
	}

	return arr
}

func (r Repository) GetAllPsgAreaCode(code string) []Passenger {
	query := `
	SELECT * FROM passenger
    WHERE phone2 LIKE '`+code+`%'
    ORDER BY phone2 DESC;`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	return LoadData(data, err)
}

func (r Repository) GetTrainsPsgRiding(firstName string, lastName string) []string {
	query := `
	SELECT train.name FROM booked
    JOIN train
    ON booked.train_no = train.train_no
    JOIN passenger
    on booked.ssn = passenger.ssn
    WHERE 
        passenger.first_name = '`+firstName+`' AND
        passenger.last_name = '`+lastName+`' AND
        booked.status = 'Booked';`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []string
	
	for data.Next() {
		var name string
		if err := data.Scan(&name); err != nil {
			log.Fatal("")
		}
		arr = append(arr, name)
	}

	return arr
}

func (r Repository) GetPsgRiding(day string) []Passenger {
	query := `
    SELECT passenger.* FROM passenger
    JOIN booked
    on booked.ssn = passenger.ssn
    JOIN train
    on booked.train_no = train.train_no
    WHERE
        train.available LIKE '%`+day+`%' AND
        booked.status = 'Booked';`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	return LoadData(data, err)
}

func (r Repository) GetPsgRidingConfirmed(day string) []string {
	query := `
	SELECT 
        passenger.first_name,
		passenger.last_name 
    FROM passenger
    JOIN booked
    ON passenger.ssn = booked.ssn
    JOIN train
    ON train.train_no = booked.train_no
    WHERE 
        train.available LIKE '%`+day+`%'
    ORDER BY passenger ASC;`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	if err != nil {
		return nil
	}
	defer data.Close()

	var arr []string
	
	for data.Next() {
		var firstName string
		var lastName string
		if err := data.Scan(
			&firstName,
			&lastName,
		); err != nil {
			log.Fatal("Unable to process data.")
		}
		arr = append(arr, firstName+" "+lastName)
	}

	return arr
}

func (r Repository) GetPsgConfirmedTrainName(name string) []Passenger {
	query := `
	SELECT passenger.* FROM passenger
	JOIN booked
	on passenger.ssn = booked.ssn
	JOIN train
	on train.train_no = booked.train_no
	WHERE 
		booked.status = 'Booked' AND 
		train.name = '`+name+`';`

	ctx := context.Background()
	data, err := r.db.Query(ctx, query)

	return LoadData(data, err)
}