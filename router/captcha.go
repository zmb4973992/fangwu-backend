package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type captcha struct{}

func (c *captcha) appendCaptchaRouterTo(param *gin.RouterGroup) {
	var captchaController controller.Cpatcha

	captchaRouter := param.Group("/captcha")
	captchaRouter.Use(middleware.RateLimit("ip", 3, "per_second"))

	publicRouter := captchaRouter.Group("")
	publicRouter.GET("", captchaController.Get)
}
