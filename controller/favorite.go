package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Favorite struct{}

func (f *Favorite) Get(c *gin.Context) {
	var param service.FavoriteGet
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

func (f *Favorite) Create(c *gin.Context) {
	var param service.FavoriteCreate
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

func (f *Favorite) Delete(c *gin.Context) {
	var param service.FavoriteDelete
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

func (f *Favorite) GetList(c *gin.Context) {
	var param service.FavoriteGetList
	err := c.ShouldBindJSON(&param)

	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateList(nil, nil, util.ErrorInvalidJsonParams, util.GetErrDetail(err)),
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

	param.UserId = userId
	result, paging, resCode, errDetail := param.GetList()
	c.JSON(
		http.StatusOK,
		response.GenerateList(result, paging, resCode, errDetail),
	)
}
