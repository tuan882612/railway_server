package passenger

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