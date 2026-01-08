package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pwshRivan/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	routeService := NewService()

	route := router.Group("/routes")
	route.GET("", func(c *gin.Context) {
		GetAllRoutes(c, routeService)
	})

	route.GET("/:id", func(c *gin.Context) {
		GetRouteById(c, routeService)
	})
}

func GetAllRoutes(c *gin.Context, service Service) {
	datas, err := service.GetAllRoutes()
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
		Message: "successfully get all routes",
		Data:    datas,
	})
}

func GetRouteById(c *gin.Context, service Service) {
	id := c.Param("id")

	data, err := service.GetRouteById(id)
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
		Message: "successfully get route",
		Data:    data,
	})
}
