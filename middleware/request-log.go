package middleware

import (
	"time"

	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"

	"github.com/gin-gonic/gin"
)

var channelForRequestLog = make(chan model.RequestLog, 5)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		// 执行耗时（毫秒）
		cost := int(endTime.Sub(startTime).Milliseconds())
		var requestLog model.RequestLog
		userId, resCode, _ := util.GetUserId(c)
		if resCode == util.Success {
			requestLog.Creator = &userId
			requestLog.LastModifier = &userId
		}

		//获取访问路径
		tempPath := c.FullPath()
		requestLog.Path = &tempPath

		//获取URI参数
		//requestLog.URIParams = c.Params

		//获取请求方式
		tempMethod := c.Request.Method
		requestLog.Method = &tempMethod

		//获取ip
		tempIP := c.ClientIP()
		requestLog.IP = &tempIP

		//获取响应码
		tempCode := c.Writer.Status()
		requestLog.ResponseCode = &tempCode

		//获取开始时间和执行耗时(毫秒)
		requestLog.StartTime = &startTime
		requestLog.TimeElapsed = &cost

		//获取用户的浏览器标识
		tempUserAgent := c.Request.UserAgent()
		requestLog.UserAgent = &tempUserAgent

		//把日志放到通道中，等待保存到数据库
		channelForRequestLog <- requestLog

	}
}

func SaveRequestLog() {
	for {
		select {
		case requestLog := <-channelForRequestLog:
			global.Db.Create(&requestLog)
		}
	}
}
