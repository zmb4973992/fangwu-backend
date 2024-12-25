package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type complaint struct{}

func (c *complaint) appendComplaintRouterTo(param *gin.RouterGroup) {
	var complaintController controller.Complaint

	complaintRouter := param.Group("/complaint")
	complaintRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	privateRouter := complaintRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("/:id", complaintController.Get)
	privateRouter.POST("", complaintController.Create)
	privateRouter.POST("/list", complaintController.GetList)
}
