package routes

import (
	"main/internal/api/handler"
	"main/internal/domain/passenger"

	"github.com/gin-gonic/gin"
)

func PassengerRoutes(engine *gin.RouterGroup) {
	pHandler := &handler.PassengerController{Repo: passenger.InitRepo()}
	psg := engine.Group("/passenger")
	psg.GET("/all", pHandler.GetAll)
	psg.GET("/waiting", pHandler.GetWaiting)
	psg.GET("/area", pHandler.GetArea)
	psg.GET("/riding", pHandler.GetTrainsRiding)
	psg.GET("/confirmed", pHandler.GetAllwTrainName)
	psg.GET("/traveling", pHandler.GetTraveling)
	psg.GET("/traveling/confirmed", pHandler.GetTravelingConfirmed)
}
