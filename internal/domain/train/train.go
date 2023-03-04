package train

type Train struct {
	Train_no 	int	     `json:"train_no"`
	Name 		string   `json:"name"`
	Source 		string   `json:"source"`
	Destination string   `json:"destination"`
	Cost_p		float32  `json:"cost_p"`
	Cost_g		float32  `json:"cost_g"`
	Avaibility  []string `json:"avaibility"`
}