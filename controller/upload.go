package controller

import (
	"fangwu-backend/global"
	"fangwu-backend/response"
	"fangwu-backend/service"
	"fangwu-backend/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func (u *Upload) InitUploadingPath() {
	//检查上传文件夹是否存在
	var directory util.Directory
	exists := directory.PathExistsOrNot(global.Config.Upload.StoragePath)
	//如果不存在就创建
	if !exists {
		err := os.MkdirAll(global.Config.Upload.StoragePath, os.ModePerm)
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}

	exists = directory.PathExistsOrNot(global.Config.Upload.TmpStoragePath)
	if !exists {
		err := os.MkdirAll(global.Config.Upload.TmpStoragePath, os.ModePerm)
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}
}

func (u *Upload) Create(c *gin.Context) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorFailToUpload, util.GetErrDetail(err)),
		)
		return
	}

	var param service.UploadCreate
	param.FileHeader = fileHeader

	userId, resCode, errDetail := util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	param.UserId = userId

	result, resCode, errDetail := param.Create()
	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}

func (u *Upload) BatchCreate(c *gin.Context) {
	var param service.UploadBatchCreate
	var resCode int
	var errDetail *util.ErrDetail

	param.UserId, resCode, errDetail = util.GetUserId(c)
	if resCode != util.Success {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, resCode, errDetail),
		)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(
			http.StatusOK,
			response.GenerateSingle(nil, util.ErrorFailToParseMultipartForm, util.GetErrDetail(err)),
		)
		return
	}

	var result []*service.UploadResult
	param.FileHeader = form.File["file"]
	result, resCode, errDetail = param.BatchCreate()

	c.JSON(
		http.StatusOK,
		response.GenerateSingle(result, resCode, errDetail),
	)
}
