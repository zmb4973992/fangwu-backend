package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdministrativeDivision struct{}

func (a *AdministrativeDivision) GetList(c *gin.Context) {
	var param service.AdministrativeDivisionGetList
	err := c.ShouldBindJSON(&param)

	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateList(nil, nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	result, paging, resCode, errDetail := param.GetList()
	c.JSON(
		http.StatusOK,
		response.GenerateList(result, paging, resCode, errDetail),
	)
}
