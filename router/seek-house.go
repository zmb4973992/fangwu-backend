package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type seekHouse struct{}

func (s *seekHouse) appendSeekHouseRouterTo(param *gin.RouterGroup) {
	var seekHouseController controller.SeekHouse

	seekHouseRouter := param.Group("/seek-house")
	seekHouseRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := seekHouseRouter.Group("")
	publicRouter.POST("/list", seekHouseController.GetList)

	privateRouter := seekHouseRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("/:id", seekHouseController.Get)
	privateRouter.POST("", seekHouseController.Create)
	privateRouter.PATCH("", seekHouseController.Update)
	privateRouter.DELETE("", seekHouseController.Delete)

}
