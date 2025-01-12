package util

import (
	"fangwu-backend/global"
	"os"
)

func InitUploadPath() {
	//检查上传文件夹是否存在
	var directory Directory
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
