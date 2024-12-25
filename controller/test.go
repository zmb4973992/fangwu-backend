package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct{}

func (t *Test) Respond(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(nil, util.Success, nil),
	)
}
