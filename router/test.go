package router

import (
	"fangwu-backend/controller"

	"github.com/gin-gonic/gin"
)

type test struct{}

func (t *test) appendTestRouterTo(param *gin.RouterGroup) {
	var testController controller.Test

	testRouter := param.Group("")

	testRouter.GET("/test", testController.Respond)
}
