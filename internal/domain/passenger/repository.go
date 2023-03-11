package passenger

import "main/database"

type PassengerRepository interface {
	GetAllPassengers() []Passenger 
	GetWaiting() []Waiting
	GetAllPsgAreaCode(code string) []Passenger
	GetTrainsPsgRiding(firstName string, lastName string) []string
	GetPsgRiding(day string) []Passenger
	GetPsgRidingConfirmed(day string) []string
	GetPsgConfirmedTrainName(name string) []Passenger
}

func InitRepo() PassengerRepository {
	return &Repository{db: database.GetClient()}
}