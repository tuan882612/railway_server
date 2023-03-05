package train

type Train struct {
	Train_no 	 int	  `json:"train_no"`
	Name 		 string   `json:"name"`
	Source 		 string   `json:"source"`
	Destination  string   `json:"destination"`
	Premium_fare int  	  `json:"premium_fare"`
	General_fare int  	  `json:"general_fare"`
	Available    []string `json:"available"`
}

type TrainStatus struct {
	Name 	 	string `json:"name"`
	Date	 	string `json:"date"`
	PsAvaiable  string `json:"ps_avaiable"`
	PsTaken 	string `json:"ps_taken"`
	GsAvaiable  string `json:"gs_avaiable"`
	GsTaken 	string `json:"gs_taken"`
}

type Booked struct {
	Ssn		 int    `json:"ssn"`
	TrainNo	 string `json:"train_no"`
	Category string `json:"category"`
	Status   string `json:"status"`
}