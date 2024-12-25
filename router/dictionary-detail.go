package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type dictionaryDetail struct{}

func (d *dictionaryDetail) appendDictionaryDetailRouterTo(param *gin.RouterGroup) {
	var dictionaryDetailController controller.DictionaryDetail

	dictionaryDetailRouter := param.Group("/dictionary-detail")
	dictionaryDetailRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	dictionaryDetailRouter.POST("/list", dictionaryDetailController.GetList)
}
