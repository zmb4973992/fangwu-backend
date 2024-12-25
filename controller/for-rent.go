package controller

import (
	"errors"
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ForRent struct{}

func (f *ForRent) Get(c *gin.Context) {
	var param service.ForRentGet
	var err error
	param.Id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidUriParams, util.GetErrDetail(err)),
		)
		return
	}

	result, resCode, errDetail := param.Get()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (f *ForRent) Create(c *gin.Context) {
	var param service.ForRentCreate
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	param.Creator = userId
	result, resCode, errDetail := param.Create()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (f *ForRent) Update(c *gin.Context) {
	var param service.ForRentUpdate
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
	}

	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	param.LastModifier = userId
	result, resCode, errDetail := param.Update()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (f *ForRent) Delete(c *gin.Context) {
	var param service.ForRentDelete
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
		)
		return
	}

	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	param.Deleter = userId
	resCode, errDetail = param.Delete()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(nil, resCode, errDetail),
	)
}

func (f *ForRent) GetList(c *gin.Context) {
	var param service.ForRentGetList
	err := c.ShouldBindJSON(&param)

	//如果json没有传参，会提示EOF错误，这里允许不传参的查询
	//如果是其他错误，就正常报错
	if err != nil && !errors.Is(err, io.EOF) {
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
