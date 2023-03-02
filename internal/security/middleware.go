package security

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetMiddleware(engine *gin.Engine) {
	origins := []string{
		"http://localhost:3000",
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
	
	engine.Use(cors)
}