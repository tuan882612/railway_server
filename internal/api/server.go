package api

import (
	"main/internal/api/routes"
	"strconv"

	"github.com/gin-contrib/cors"
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
	r.SetTrustedProxies()
	r.SetMiddleware()
	r.ApiRoutes()
	r.Engine.Run(address+":"+strconv.Itoa(port))
}

func (r *RailwayServer) ApiRoutes() {
	v1 := r.Engine.Group("/api/v1")
	routes.BaseRoutes(v1)
	routes.PassengerRoutes(v1)
	routes.TrainRoutes(v1)
}

func (r *RailwayServer) SetTrustedProxies() {
	proxies := []string{
		"192.168.1.2", 
		"localhost",
		"0.0.0.0",
		"127.0.0.1",
	}

	r.Engine.SetTrustedProxies(proxies)
}

func (r *RailwayServer) SetMiddleware() {
	origins := []string{
		"http://localhost:3000",
		"http://localhost:1001",
		"http://localhost:1002",
		"http://localhost:1003",
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = origins
	config.AllowCredentials = true
	config.AddAllowHeaders(
		"Access-Control-Allow-Headers",
		"access-control-allow-origin, access-control-allow-headers",
		"Content-Type", "X-XSRF-TOKEN",
		"Accept", "Origin", "X-Requested-With", "Authorization")
	config.AddAllowMethods("*")
	cors := cors.New(config)
	
	r.Engine.Use(cors)
}