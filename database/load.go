package database

import (
	"context"

)

func LoadData() {
	client := GetClient()
	defer client.Close(context.Background())
	
	// Used normal SQL to load data
}