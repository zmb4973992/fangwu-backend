package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type login struct{}

func (l *login) appendLoginRouterTo(param *gin.RouterGroup) {
	var userController controller.User

	loginRouter := param.Group("")
	loginRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	loginRouter.POST("/login", userController.Login)
}
