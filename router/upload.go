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
	uploadRouter.Use(middleware.JWT())
	uploadRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	//上传单个文件
	uploadSingleRouter := uploadRouter.Group("")
	uploadSingleRouter.Use(middleware.UploadFilterForSingle())
	uploadSingleRouter.POST("", uploadController.Create)

	//上传多个文件
	uploadBatchRouter := uploadRouter.Group("")
	uploadBatchRouter.Use(middleware.UploadFilterForBatch())
	uploadBatchRouter.POST("/batch", uploadController.BatchCreate)
}
