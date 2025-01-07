package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type administrativeDivision struct{}

func (a *administrativeDivision) appendAdministrativeDivisionRouterTo(param *gin.RouterGroup) {
	var administrativeDivisionController controller.AdminDiv

	administrativeDivisionRouter := param.Group("/admin-div")
	administrativeDivisionRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := administrativeDivisionRouter.Group("")
	publicRouter.POST("", administrativeDivisionController.GetByName)
	publicRouter.POST("/list", administrativeDivisionController.GetList)
}
