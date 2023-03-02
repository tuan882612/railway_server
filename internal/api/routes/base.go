package routes

import (
	"main/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func BaseRoutes(engine *gin.RouterGroup) {
	engine.GET("/", handler.Default)
	engine.GET("/health", handler.Health)
}