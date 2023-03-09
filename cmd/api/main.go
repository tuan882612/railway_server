package main

import (
	// "main/internal/api"
	"main/database"
	"main/pkg/config"
	"main/pkg/validators"
	// "github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.DebugMode)

	config.LoadEnv(".env")
	validators.ValidateLoadEnv()
	validators.ValidatePostgres()

	// server := api.InitServer("localhost",2000)
	// server.LoadConfig()
	// server.Run()
	database.CreateTables()
}