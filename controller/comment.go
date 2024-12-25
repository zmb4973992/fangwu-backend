package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

func (co *Comment) Create(c *gin.Context) {
	var param service.CommentCreate
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

func (co *Comment) Update(c *gin.Context) {
	var param service.CommentUpdate
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

func (co *Comment) Delete(c *gin.Context) {
	var param service.CommentDelete
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

func (co *Comment) GetList(c *gin.Context) {
	var param service.CommentGetList
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
