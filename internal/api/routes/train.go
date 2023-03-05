package routes

import "github.com/gin-gonic/gin"

func TrainRoutes(engine *gin.RouterGroup) {
	train := engine.Group("/train")
	train.GET("/info")
	train.GET("/info/age")
	train.GET("/confirmed")
}

/*
	/info - get info on train on that day from train name.
	/info/age - get train info from age of passenger entered.
	/confirmed - get all psg with confirmed status on train from train name.
*/