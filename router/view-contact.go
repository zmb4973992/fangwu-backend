package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type viewContact struct{}

func (v *viewContact) appendViewContactRouterTo(param *gin.RouterGroup) {
	var viewContactController controller.ViewContact

	viewContactRouter := param.Group("/view-contact")
	viewContactRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	privateRouter := viewContactRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("/:business_type/:business_id", viewContactController.Get)
	privateRouter.POST("", viewContactController.Create)
	privateRouter.DELETE("", viewContactController.Delete)
	privateRouter.GET("/count", viewContactController.Count)
}
