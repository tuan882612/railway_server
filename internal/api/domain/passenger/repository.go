package passenger

type PassengerRepository interface {

}

func Init_Repo() PassengerRepository {
	return &Repository{}
}