package passenger

import (
	"log"

	"github.com/jackc/pgx/v4"
)

type Passenger struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City  	  string `json:"city"`
	County	  string `json:"County"`
	Phone2	  string `json:"phone2"`
	Ssn		  int 	 `json:"ssn"`
	Bdate	  string `json:"bdate"`
}

type Waiting struct {
	Name	  string `json:"Name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City  	  string `json:"city"`
	County	  string `json:"County"`
	Phone2	  string `json:"phone2"`
	Ssn		  int 	 `json:"ssn"`
	Bdate	  string `json:"bdate"`
}

func ValidatePassenger(data pgx.Rows, passenger *Passenger) {
	if err := data.Scan(
		&passenger.FirstName,
		&passenger.LastName,
		&passenger.Address,
		&passenger.City,
		&passenger.County,
		&passenger.Phone2,
		&passenger.Ssn,
		&passenger.Bdate,
	); err != nil {
		log.Fatal("Unable to process data.")
	}
}

func ValidateWaiting(data pgx.Rows, waiting *Waiting) {
	if err := data.Scan(
		&waiting.Name,
		&waiting.FirstName,
		&waiting.LastName,
		&waiting.Address,
		&waiting.City,
		&waiting.County,
		&waiting.Phone2,
		&waiting.Ssn,
		&waiting.Bdate,
	); err != nil {
		log.Fatal("Unable to process data.")
	}
}

func LoadData(data pgx.Rows, err error) []Passenger {
	if err != nil {
		return nil
	}
	defer data.Close()
	
	var arr []Passenger

	for data.Next() {
		var passenger Passenger
		ValidatePassenger(data, &passenger)
		arr = append(arr, passenger)
	}

	return arr
}