package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域配置，参考了gin官方的cors包
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "Authorization", "access_token", "x-requested-with"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{
		"New-Token", "New-Expires-In", "Content-Disposition"}
	return cors.New(config)
}
