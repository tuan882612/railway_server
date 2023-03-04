package main

import (
	"main/internal/api"
	"main/tools/config"
	"main/tools/validators"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	config.LoadEnv("../../.env")
	validators.ValidateLoadEnv()
	validators.ValidatePostgres()

	server := api.InitServer()
	server.Start("localhost", 1000)
}