package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type upload struct{}

func (f *upload) appendUploadRouterTo(param *gin.RouterGroup) {
	var uploadController controller.Upload

	uploadRouter := param.Group("/upload")
	uploadRouter.Use(
		middleware.JWT(),
		middleware.UploadInterceptor(),
		middleware.RateLimit("ip", 1, "per_second"),
	)

	uploadRouter.POST("", uploadController.Create)
}
