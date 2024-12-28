package controller

import (
	"net/http"

	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) Login(c *gin.Context) {
	var param service.UserLogin
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(
				nil,
				util.ErrorInvalidJsonParams,
				util.GetErrDetail(err),
			))
		return
	}

	var needCaptcha bool = false

	if needCaptcha {
		permitted := param.Verify()
		if !permitted {
			c.JSON(
				http.StatusOK,
				response.GenerateSingle(nil, util.ErrorWrongCaptcha, nil))
			return
		}
	}

	result, resCode, errDetail := param.Login()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (u *User) Get(c *gin.Context) {
	var param service.UserGet
	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail))
		return
	}

	param.Id = userId
	output, resCode, errDetail := param.Get()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(output, resCode, errDetail))
}

func (u *User) Register(c *gin.Context) {
	var param service.UserCreate
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(
				nil,
				util.ErrorInvalidJsonParams,
				util.GetErrDetail(err)))
		return
	}

	// ip := "38.39.55.123"
	ip := c.ClientIP()
	param.Ip = &ip

	result, resCode, errDetail := param.Create()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail))
}

func (u *User) Update(c *gin.Context) {
	var param service.UserUpdate
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(
				nil,
				util.ErrorInvalidJsonParams,
				util.GetErrDetail(err)))
		return
	}

	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail))
		return
	}

	param.Id = userId

	resCode, errDetail = param.Update()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(nil, resCode, errDetail),
	)
}

func (u *User) Delete(c *gin.Context) {
	var param service.UserDelete
	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	param.Id = userId
	resCode, errDetail = param.Delete()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(nil, resCode, errDetail),
	)
}
