package main

import (
	"main/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := api.InitServer()
	server.ServeHTTP("0.0.0.0", 1000)
}