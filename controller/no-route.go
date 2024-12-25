package controller

import (
	"fangwu-backend/response"
	"fangwu-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoRoute struct{}

func (n *NoRoute) Respond(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(nil, util.ErrorInvalidRequest, nil),
	)
}
