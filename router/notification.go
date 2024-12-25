package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type notification struct{}

func (n *notification) appendNotificationRouterTo(param *gin.RouterGroup) {
	var notificationController controller.Notification

	notificationRouter := param.Group("/notification")
	notificationRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	privateRouter := notificationRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.POST("", notificationController.Create)
	privateRouter.PATCH("", notificationController.Update)
	privateRouter.POST("/list", notificationController.GetList)
}
