package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func (u *Upload) Create(c *gin.Context) {
	var param service.UploadCreate
	var resCode int
	var errDetail *util.ErrDetail

	param.UserId, resCode, errDetail = util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorFailToParseMultipartForm, util.GetErrDetail(err)),
		)
		return
	}

	var result []*service.UploadResult
	param.FileHeader = form.File["file"]
	result, resCode, errDetail = param.Create()

	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}
