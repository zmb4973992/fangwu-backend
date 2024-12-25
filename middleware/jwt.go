package middleware

import (
	"fangwu-backend/response"
	"fangwu-backend/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("access_token")
		//如果请求头没有携带access_token
		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(nil, util.ErrorAccessTokenNotFound, nil))
			return
		}

		//开始校验access_token
		res, err := util.ParseToken(token)
		//如果token错误
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK,
				response.GenerateSingle(
					nil,
					util.ErrorInvalidAccessToken,
					util.GetErrDetail(err),
				))
			return
		}
		//如果token过期
		if res.ExpiresAt.Unix() < time.Now().Unix() {
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(nil, util.ErrorAccessTokenExpired, nil),
			)
			return
		}

		c.Next()
	}
}
