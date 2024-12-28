package router

import (
	"fangwu-backend/controller"
	"fangwu-backend/global"
	"fangwu-backend/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 路由增强器，入参为路由组，用来给路由组指定中间件和路由，同时也为了拆分路由组为单文件
type routerEnhancer struct {
	test
	login
	downLoad
	dictionaryDetail
	static
	user
	upload
	forRent
	fileConfirm
	seekHouse
	contactInfoBlacklist
	complaint
	comment
	userBlacklist
	notification
	captcha
	favorite
	viewContact
	administrativeDivision
}

// LoadEngine 初始化引擎,最终返回*gin.Engine类型，给main调用
func LoadEngine() *gin.Engine {
	//设置运行模式
	gin.SetMode(global.Config.Gin.Mode)
	fmt.Printf("当前运行模式为：%s", gin.Mode())
	engine := gin.New()

	//全局中间件
	engine.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.Cors(),
		// middleware.RequestLog(),
	)

	//路由不匹配时的处理
	noRouteController := new(controller.NoRoute)
	engine.NoRoute(noRouteController.Respond)

	defaultGroup := engine.Group("")

	var routerEnhancer routerEnhancer
	routerEnhancer.appendTestRouterTo(defaultGroup)
	routerEnhancer.appendLoginRouterTo(defaultGroup)
	routerEnhancer.appendDownloadRouterTo(defaultGroup)
	routerEnhancer.appendDictionaryDetailRouterTo(defaultGroup)
	routerEnhancer.appendStaticRouterTo(defaultGroup)
	routerEnhancer.appendUserRouterTo(defaultGroup)
	routerEnhancer.appendUploadRouterTo(defaultGroup)
	routerEnhancer.appendForRentRouterTo(defaultGroup)
	routerEnhancer.appendFileConfirmRouterTo(defaultGroup)
	routerEnhancer.appendSeekHouseRouterTo(defaultGroup)
	routerEnhancer.appendContactInfoBlacklistRouterTo(defaultGroup)
	routerEnhancer.appendComplaintRouterTo(defaultGroup)
	routerEnhancer.appendCommentRouterTo(defaultGroup)
	routerEnhancer.appendUserBlacklistRouterTo(defaultGroup)
	routerEnhancer.appendNotificationRouterTo(defaultGroup)
	routerEnhancer.appendCaptchaRouterTo(defaultGroup)
	routerEnhancer.appendFavoriteRouterTo(defaultGroup)
	routerEnhancer.appendViewContactRouterTo(defaultGroup)
	routerEnhancer.appendAdministrativeDivisionRouterTo(defaultGroup)

	//引擎配置完成后，返回
	return engine
}
