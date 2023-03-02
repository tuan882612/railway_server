package api

import (
	"main/internal/api/handler"
	"main/internal/api/routes"
	"main/internal/security"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RailwayServer struct {
	Engine *gin.Engine
}

func InitServer() *RailwayServer {
	return &RailwayServer{}
}

func (r *RailwayServer) ServeHTTP(address string, port int) {
	r.Engine = gin.Default()
	r.ApiRoutes()
	security.SetTrustedProxies(r.Engine)
	security.SetMiddleware(r.Engine)
	r.Engine.Run(address+":"+strconv.Itoa(port))
}

func (r *RailwayServer) ApiRoutes() {
	v1 := r.Engine.Group("/api/v1")
	routes.BaseRoutes(v1)
	routes.PassengerRoutes(v1)
	routes.TrainRoutes(v1)
	r.Engine.NoRoute(handler.None)
}