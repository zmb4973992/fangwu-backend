package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DictionaryDetail struct{}

func (d *DictionaryDetail) GetList(c *gin.Context) {
	// 将json参数绑定到结构体
	var dictionaryType service.DictionaryTypeGet
	err := c.ShouldBindJSON(&dictionaryType)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateList(nil, nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	//获取字典类型的name，去获取字典类型的id
	result, resCode, errDetail := dictionaryType.Get()
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateList(nil, nil, resCode, errDetail),
		)
		return
	}

	//根据字典类型的id，去获取字典详情的列表
	var param service.DictionaryDetailGetList
	param.DictionaryTypeId = result.Id
	result1, paging, resCode, errDetail := param.GetList()
	c.JSON(
		http.StatusOK,
		response.GenerateList(result1, paging, resCode, errDetail),
	)
}
