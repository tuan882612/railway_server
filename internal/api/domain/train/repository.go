package train

type TrainRepository interface {

}

func Init_Repo() TrainRepository {
	return &Repository{}
}