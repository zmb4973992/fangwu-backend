package main

import (
	"fangwu-backend/controller"
	"fangwu-backend/global"
	"fangwu-backend/middleware"
	"fangwu-backend/model"
	"fangwu-backend/router"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"strconv"
)

func main() {

	//加载全局变量
	global.LoadConfig()
	//加载日志记录器，使用的是zap
	util.LoadLogger()
	//加载id生成器
	var idGenerator util.IdGenerator
	idGenerator.Load()
	//连接到数据库
	model.ConnectToDb()
	//初始化上传路径
	var file controller.Upload
	file.InitUploadingPath()
	//加载验证码
	service.LoadCaptcha()
	//生成引擎
	engine := router.LoadEngine()
	//开启2个协程，用来保存访问记录到数据库
	for i := 0; i < 2; i++ {
		go middleware.SaveRequestLog()
	}
	//开启定时任务
	// cron.Init()

	//运行服务，必须放最后
	err := engine.Run(":" + strconv.Itoa(global.Config.Access.Port))
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}
}
