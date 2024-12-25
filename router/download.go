package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type downLoad struct{}

func (d *downLoad) appendDownloadRouterTo(param *gin.RouterGroup) {
	var downloadController controller.Download

	downloadRouter := param.Group("/download")
	downloadRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	downloadRouter.GET("/:file-id", downloadController.Get) //下载文件
}
