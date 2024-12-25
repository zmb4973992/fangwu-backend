package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type userBlacklist struct{}

func (u *userBlacklist) appendUserBlacklistRouterTo(param *gin.RouterGroup) {
	var userBlacklistController controller.UserBlacklist

	userBlacklistRouter := param.Group("/user-blacklist")
	userBlacklistRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	privateRouter := userBlacklistRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.POST("/verify", userBlacklistController.Verify)
	privateRouter.POST("", userBlacklistController.Create)
	privateRouter.DELETE("", userBlacklistController.Delete)
	privateRouter.POST("/list", userBlacklistController.GetList)

}
