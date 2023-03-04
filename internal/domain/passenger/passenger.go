package passenger

type Passenger struct {
	Name 	 string `json:"name"`
	Age		 int	`json:"age"`
	Address  string `json:"address"`
	Category string `json:"category"`
	Status	 string `json:"status"`
}