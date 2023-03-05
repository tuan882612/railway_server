package routes

import "github.com/gin-gonic/gin"

func PassengerRoutes(engine *gin.RouterGroup) {
	psg := engine.Group("/passenger")
	psg.GET("/area")
	psg.GET("/riding")
	psg.GET("/traveling")
	psg.GET("/traveling/confirmed")
	psg.GET("/waiting")
}

/*
	/area - get all passengers from area code. (desc order)
	/riding - get all trains passenger is riding from first & last name.
	/traveling - get all passengers traveling on day entered. (asc order)
	/traveling/confirmed - same as travel endpoint but for confirmed.
	/waiting - get all passengers that are waitlisted. 
*/