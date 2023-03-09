package train

type Train struct {
	Train_no 	 int	  `json:"train_no"`
	Name 		 string   `json:"name"`
	Premium_fare int  	  `json:"premium_fare"`
	General_fare int  	  `json:"general_fare"`
	Source 		 string   `json:"source"`
	Destination  string   `json:"destination"`
	Available    []string `json:"available"`
}

type TrainStatus struct {
	Date	 	string `json:"date"`
	Name 	 	string `json:"name"`
	PsAvaiable  string `json:"ps_avaiable"`
	GsAvaiable  string `json:"gs_avaiable"`
	Ps_Occupied	string `json:"ps_occupied"`
	Gs_Occupied	string `json:"gs_occupied"`
}

type Booked struct {
	Ssn		 int    `json:"ssn"`
	TrainNo	 string `json:"train_no"`
	Category string `json:"category"`
	Status   string `json:"status"`
}