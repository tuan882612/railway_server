package routes

import (
	"main/internal/api/handler"
	"main/internal/domain/train"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(engine *gin.RouterGroup) {
	hHandler := &handler.TrainController{Repo: train.InitRepo()}
	train := engine.Group("/train")
	train.GET("/all", hHandler.GetAll)
	train.GET("/info", hHandler.GetInfo)
	train.GET("/info/age", hHandler.GetInfoAge)
}
