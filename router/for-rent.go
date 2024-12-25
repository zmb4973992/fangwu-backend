package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type forRent struct{}

func (f *forRent) appendForRentRouterTo(param *gin.RouterGroup) {
	var forRentController controller.ForRent

	forRentRouter := param.Group("/for-rent")
	forRentRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := forRentRouter.Group("")
	publicRouter.POST("/list", forRentController.GetList)

	privateRouter := forRentRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("/:id", forRentController.Get)
	privateRouter.POST("", forRentController.Create)
	privateRouter.PATCH("", forRentController.Update)
	privateRouter.DELETE("", forRentController.Delete)

}
