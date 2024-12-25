package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type comment struct{}

func (c *comment) appendCommentRouterTo(param *gin.RouterGroup) {
	var commentController controller.Comment

	commentRouter := param.Group("/comment")
	commentRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := commentRouter.Group("")
	publicRouter.POST("/list", commentController.GetList)

	privateRouter := commentRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.POST("", commentController.Create)
	privateRouter.PATCH("", commentController.Update)
	privateRouter.DELETE("", commentController.Delete)
}
