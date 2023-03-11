package train

import "main/database"

type TrainRepository interface {
	GetAllTrains() []Train
	GetInfoAgeRange(start string, end string) []JoinedInfo
	GetTrainInfoSum(name string) []TrainInfo
}

func InitRepo() TrainRepository {
	return &Repository{db: database.GetClient()}
}