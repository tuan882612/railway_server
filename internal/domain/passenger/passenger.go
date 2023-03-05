package passenger

type Passenger struct {
	Ssn		  int 	 `json:"ssn"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bdate	  string `json:"bdate"`
	Phone2	  string `json:"phone2"`
	Address   string `json:"address"`
	City  	  string `json:"city"`
	County	  string `json:"County"`
}