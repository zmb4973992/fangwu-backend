package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cpatcha struct{}

func (ca *Cpatcha) Get(c *gin.Context) {
	var param service.CaptchaGet

	result, resCode, errDetail := param.Get()
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, nil),
	)
}
