package train

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pwshRivan/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	trainService := NewService()

	train := router.Group("/trains")
	train.GET("", func(c *gin.Context) {
		GetAllTrains(c, trainService)
	})

	train.GET("/:id", func(c *gin.Context) {
		GetTrainPosition(c, trainService)
	})
}

func GetAllTrains(c *gin.Context, service Service) {
	datas, err := service.GetAllTrains()
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
		Message: "successfully get all trains",
		Data:    datas,
	})
}

func GetTrainPosition(c *gin.Context, service Service) {
	id := c.Param("id")

	data, err := service.GetTrainPosition(id)
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
		Message: "successfully get train position",
		Data:    data,
	})
}
