package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type user struct{}

func (u *user) appendUserRouterTo(param *gin.RouterGroup) {
	var userController controller.User

	userRouter := param.Group("/user")
	userRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	publicRouter := userRouter.Group("")
	publicRouter.POST("", userController.Register) //注册用户

	privateRouter := userRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("", userController.Get)       //获取用户详情
	privateRouter.PATCH("", userController.Update)  //修改用户详情
	privateRouter.DELETE("", userController.Delete) //删除用户

}
