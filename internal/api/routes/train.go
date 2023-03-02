package routes

import "github.com/gin-gonic/gin"

func TrainRoutes(engine *gin.RouterGroup) {
	train := engine.Group("/train")
	train.GET("/")
}