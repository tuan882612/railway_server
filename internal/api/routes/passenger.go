package routes

import "github.com/gin-gonic/gin"

func PassengerRoutes(engine *gin.RouterGroup) {
	psg := engine.Group("/passenger")
	psg.GET("/")
}