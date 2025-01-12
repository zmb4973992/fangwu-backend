package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type adminDiv struct{}

func (a *adminDiv) appendAdminDivRouterTo(param *gin.RouterGroup) {
	var adminDivController controller.AdminDiv

	adminDivRouter := param.Group("/admin-div")
	adminDivRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := adminDivRouter.Group("")
	publicRouter.POST("/list", adminDivController.GetList)
}
