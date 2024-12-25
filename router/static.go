package router

import (
	"fangwu-backend/global"
	"fangwu-backend/middleware"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type static struct{}

func (s *static) appendStaticRouterTo(param *gin.RouterGroup) {
	staticRouter := param.Group("")
	staticRouter.Use(middleware.RateLimit("ip", 100, "per_minute"))
	staticRouter.Use(middleware.RateLimit("ip", 1000, "per_hour"))
	staticRouter.Use(middleware.RateLimit("ip", 10000, "per_day"))

	staticRouter.Static("/image", filepath.Join(global.Config.Upload.StoragePath))
}
