package fare

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pwshRivan/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	fareService := NewService()

	fare := router.Group("/fares")
	fare.GET("", func(c *gin.Context) {
		GetAllFares(c, fareService)
	})

	fare.GET("/check", func(c *gin.Context) {
		CheckFare(c, fareService)
	})
}

func GetAllFares(c *gin.Context, service Service) {
	datas, err := service.GetAllFares()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "successfully get all fares",
		Data:    datas,
	})
}

func CheckFare(c *gin.Context, service Service) {
	fromStation := c.Query("from")
	toStation := c.Query("to")

	if fromStation == "" || toStation == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "from and to query parameters are required",
			Data:    nil,
		})
		return
	}

	data, err := service.GetFare(fromStation, toStation)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "successfully get fare",
		Data:    data,
	})
}
