package api

import (
	"main/internal/api/handler"
	"main/internal/api/routes"
	"main/internal/security"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RailwayServer struct {
	Engine  *gin.Engine
	Address string
	Port	int
}

func InitServer(address string, port int) *RailwayServer {
	return &RailwayServer{gin.Default(), address, port}
}

func (r *RailwayServer) Run() {
	r.Engine.Run(r.Address+":"+strconv.Itoa(r.Port))
}

func (r *RailwayServer) LoadConfig() {
	r.ApiRoutes()
	security.SetTrustedProxies(r.Engine)
	security.SetMiddleware(r.Engine)
}

func (r *RailwayServer) ApiRoutes() {
	v1 := r.Engine.Group("/api/v1")
	routes.BaseRoutes(v1)
	routes.PassengerRoutes(v1)
	routes.TrainRoutes(v1)
	r.Engine.NoRoute(handler.None)
}