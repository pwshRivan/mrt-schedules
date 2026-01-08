package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pwshRivan/mrt-schedules/modules/fare"
	"github.com/pwshRivan/mrt-schedules/modules/route"
	"github.com/pwshRivan/mrt-schedules/modules/station"
	"github.com/pwshRivan/mrt-schedules/modules/train"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	var api = router.Group("/v1/api")
	station.Initiate(api)
	train.Initiate(api)
	route.Initiate(api)
	fare.Initiate(api)
	router.Run(":8080")
}
