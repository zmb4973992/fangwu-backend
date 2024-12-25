package router

import (
	"fangwu-backend/controller"

	"github.com/gin-gonic/gin"
)

type fileConfirm struct{}

func (f *fileConfirm) appendFileConfirmRouterTo(param *gin.RouterGroup) {
	var fileConfirmController controller.FileConfirm

	fileConfirmRouter := param.Group("/file-confirm")

	//确认上传多个文件
	fileConfirmRouter.POST("/batch", fileConfirmController.BatchConfirm)
}
