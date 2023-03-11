package train

import (
	"log"

	"github.com/jackc/pgx/v4"
)

type Train struct {
	Train_no 	 int	  `json:"train_no"`
	Name 		 string   `json:"name"`
	Premium_fare int  	  `json:"premium_fare"`
	General_fare int  	  `json:"general_fare"`
	Source 		 string   `json:"source"`
	Destination  string   `json:"destination"`
	Available    string   `json:"available"`
}

type JoinedInfo struct {
	FirstName 	string `json:"firt_name"`
    LastName  	string `json:"last_name"`
    TrainNo	  	int    `json:"train_no"`
    Name	    string `json:"name"`
    Source		string `json:"sourcee"`
    Destination string `json:"destination"`
    Status		string `json:"status"`
    Category	string `json:"category"`
}

type TrainInfo struct {
	Name  string `json:"name"`
    Date  string `json:"date"`
    Count int	 `json:"count"`
}

func ValidateTrain(data pgx.Rows, train *Train) {
	if err := data.Scan(
		&train.Train_no,
		&train.Name,
		&train.Premium_fare,
		&train.General_fare,
		&train.Source,
		&train.Destination,  
		&train.Available,
	); err != nil {
		log.Fatal("Unabkle to process data.")
	}
}