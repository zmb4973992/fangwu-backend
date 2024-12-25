package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/middleware"

	"github.com/gin-gonic/gin"
)

type favorite struct{}

func (f *favorite) appendFavoriteRouterTo(param *gin.RouterGroup) {
	var favoriteController controller.Favorite

	favoriteRouter := param.Group("/favorite")
	favoriteRouter.Use(middleware.RateLimit("ip", 1, "per_second"))

	privateRouter := favoriteRouter.Group("")
	privateRouter.Use(middleware.JWT())
	privateRouter.GET("/:id", favoriteController.Get)
	privateRouter.POST("", favoriteController.Create)
	privateRouter.DELETE("", favoriteController.Delete)
	privateRouter.POST("/list", favoriteController.GetList)

}
