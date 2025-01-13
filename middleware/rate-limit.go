package middleware

import (
	"fangwu-backend/global"
	"fangwu-backend/response"
	"fangwu-backend/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
)

// limitedBy为限制方式（ip、user_id）， rate为限制次数（整数），timeUnit为时间单位（per_second、per_minute、per_hour、per_day、per_week、per_month）
func RateLimit(limitedBy string, rate int, timeUnit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//如果redis未开启，则不限制
		if !global.Config.Redis.Enabled {
			c.Next()
			return
		}

		limiter := redis_rate.NewLimiter(global.Rdb)

		var res *redis_rate.Result
		var err error

		switch limitedBy {
		//如果限制方式为ip
		case "ip":
			key := fmt.Sprintf("ip:%s", c.ClientIP())
			switch timeUnit {
			//如果时间单位为每秒
			case "per_second":
				res, err = limiter.Allow(c, key+":second", redis_rate.PerSecond(rate))
			//如果时间单位为每分钟
			case "per_minute":
				res, err = limiter.Allow(c, key+":minute", redis_rate.PerMinute(rate))
			//如果时间单位为每小时
			case "per_hour":
				res, err = limiter.Allow(c, key+":hour", redis_rate.PerHour(rate))
				//如果时间单位为每天
			case "per_day":
				res, err = limiter.Allow(c, key+":day", redis_rate.Limit{
					Rate:   rate,
					Period: 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为每周
			case "per_week":
				res, err = limiter.Allow(c, key+":week", redis_rate.Limit{
					Rate:   rate,
					Period: 7 * 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为每月
			case "per_month":
				res, err = limiter.Allow(c, key+":month", redis_rate.Limit{
					Rate:   rate,
					Period: 30 * 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为其他
			default:
				c.AbortWithStatusJSON(
					http.StatusOK,
					response.GenerateSingle(
						nil,
						util.ErrorInvalidTimeUnit,
						nil,
					))
				return
			}

		//如果限制方式为user_id
		case "user_id":
			userId, resCode, errDetail := util.GetUserId(c)
			if resCode != util.Success {
				c.AbortWithStatusJSON(
					http.StatusOK,
					response.GenerateSingle(nil, resCode, errDetail),
				)
				return
			}

			key := fmt.Sprintf("user_id:%d", userId)
			switch timeUnit {
			//如果时间单位为每秒
			case "per_second":
				res, err = limiter.Allow(c, key+":second", redis_rate.PerSecond(rate))
			//如果时间单位为每分钟
			case "per_minute":
				res, err = limiter.Allow(c, key+":minute", redis_rate.PerMinute(rate))
			//如果时间单位为每小时
			case "per_hour":
				res, err = limiter.Allow(c, key+":hour", redis_rate.PerHour(rate))
				//如果时间单位为每天
			case "per_day":
				res, err = limiter.Allow(c, key+":day", redis_rate.Limit{
					Rate:   rate,
					Period: 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为每周
			case "per_week":
				res, err = limiter.Allow(c, key+":week", redis_rate.Limit{
					Rate:   rate,
					Period: 7 * 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为每月
			case "per_month":
				res, err = limiter.Allow(c, key+":month", redis_rate.Limit{
					Rate:   rate,
					Period: 30 * 24 * time.Hour,
					Burst:  rate,
				})
			//如果时间单位为其他
			default:
				c.AbortWithStatusJSON(
					http.StatusOK,
					response.GenerateSingle(
						nil,
						util.ErrorInvalidTimeUnit,
						nil,
					))
				return
			}

		//如果限制方式为其他
		default:
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(
					nil,
					util.ErrorInvalidLimitedBy,
					nil,
				))
			return
		}

		//如果limiter.Allow返回错误
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(
					nil,
					util.ErrorRateLimitDoesNotWork,
					util.GetErrDetail(err),
				))
			return
		}

		//如果res.Allowed为0，表示请求过于频繁
		if res.Allowed == 0 {
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(
					nil,
					util.ErrorRequestTooFast,
					nil,
				))
			return
		}

		c.Next()
	}
}
