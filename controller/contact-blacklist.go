package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactInfoBlacklist struct{}

func (co *ContactInfoBlacklist) Verify(c *gin.Context) {
	var param service.ContactBlackListVerify
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	isBlocked, resCode, errDetail := param.Verify()
	result := map[string]bool{
		"is_blocked": isBlocked,
	}
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (co *ContactInfoBlacklist) GetList(c *gin.Context) {
	var param service.ContactBlacklistGetList
	err := c.ShouldBindJSON(&param)

	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateList(nil, nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	result, resCode, errDetail := param.GetList()
	c.JSON(
		http.StatusOK,
		response.GenerateList(result, nil, resCode, errDetail),
	)
}
