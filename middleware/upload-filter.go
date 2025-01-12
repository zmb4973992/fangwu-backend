package middleware

import (
	"fangwu-backend/global"
	"fangwu-backend/response"
	"fangwu-backend/util"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析表单数据中的文件
		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusOK,
				response.GenerateSingle(nil, util.ErrorFailToUpload, util.GetErrDetail(err)),
			)
			return
		}

		// 获取所有上传的文件
		fileHeaders := form.File["files"]
		for _, fileHeader := range fileHeaders {
			// 检查文件大小是否超过限制
			if fileHeader.Size > global.Config.Upload.MaxSize {
				c.AbortWithStatusJSON(
					http.StatusOK,
					response.GenerateSingle(nil, util.ErrorFileTooLarge, util.GetErrDetail(err)),
				)
				return
			}

			// 获取文件扩展名
			fileExt := filepath.Ext(fileHeader.Filename)

			// 检查文件扩展名是否在允许的列表中
			var allowedExts = global.Config.Upload.AllowedExts
			var isAllowed bool
			for _, ext := range allowedExts {
				// 将文件扩展名转换为小写进行比较
				if strings.ToLower(fileExt) == ext {
					isAllowed = true
					break
				}
			}

			// 如果文件类型不在允许的列表中，则返回错误信息
			if !isAllowed {
				c.AbortWithStatusJSON(
					http.StatusOK,
					response.GenerateSingle(nil, util.ErrorUnsupportedFileType, nil),
				)
				return
			}
		}

		// 如果文件符合要求，则继续处理请求
		c.Next()
	}
}
