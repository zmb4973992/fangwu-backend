package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type contactInfoBlacklist struct{}

func (c *contactInfoBlacklist) appendContactInfoBlacklistRouterTo(param *gin.RouterGroup) {
	var contactInfoBlacklistController controller.ContactInfoBlacklist

	contactInfoBlacklistRouter := param.Group("/contact-info-blacklist")

	privateRouter := contactInfoBlacklistRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.POST("/verify", contactInfoBlacklistController.Verify)
	privateRouter.POST("/list", contactInfoBlacklistController.GetList)

}
