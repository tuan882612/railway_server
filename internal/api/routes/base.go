package routes

import "github.com/gin-gonic/gin"

func BaseRoutes(engine *gin.RouterGroup) {
	engine.GET("/")
}