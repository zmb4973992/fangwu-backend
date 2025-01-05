package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Download struct{}

func (d *Download) Get(c *gin.Context) {
	var param service.ImageGet
	var err error

	param.Id, err = strconv.ParseInt(c.Param("file-id"), 10, 64)
	if err != nil {
		errDetail := util.GetErrDetail(err)
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidUriParams, errDetail),
		)
		return
	}

	result, resCode, errDetail := param.Get()
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	c.FileAttachment(result.Url, result.Name)
}
