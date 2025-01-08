package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"os"
	"path/filepath"
	"strconv"

	"fangwu-backend/util"
)

type ImageGet struct {
	Id int64 `json:"id"`
}

type imageGetList struct {
	businessType string `json:"-"`
	businessId   int64  `json:"-"`
}

type ImageResult struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

func (d *ImageGet) Get() (result *ImageResult, resCode int, errDetail *util.ErrDetail) {
	//获取文件记录
	var record model.File
	err := global.Db.
		Where("id = ?", d.Id).
		First(&record).Error
	if err != nil {
		return nil, util.ErrorFailToGetFileRecord, util.GetErrDetail(err)
	}

	//拼接文件在服务器中的保存路径
	storagePath := global.Config.Upload.StoragePath
	filePath := filepath.Join(storagePath, strconv.FormatInt(record.Id, 10)+filepath.Ext(record.Name))
	println(filePath)
	//看该文件是否存在于服务器的文件夹中
	_, err = os.Stat(filePath)
	if err != nil {
		return nil, util.ErrorFileNotFound, util.GetErrDetail(err)
	}

	return &ImageResult{Url: filePath, Name: record.Name},
		util.Success, nil
}

func (d *imageGetList) GetList() (results []ImageResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.File{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("business_type =?", d.businessType).
		Where("business_id =?", d.businessId)

	// count
	var count int64
	db.Count(&count)

	var files []model.File
	db.Find(&files)

	for _, file := range files {
		var result ImageResult

		//拼接文件在服务器中的保存路径
		storagePath := global.Config.Upload.StoragePath
		filePath := filepath.Join(storagePath,
			strconv.FormatInt(file.Id, 10)+filepath.Ext(file.Name))

		//看该文件是否存在于服务器的文件夹中
		_, err := os.Stat(filePath)
		if err != nil {
			continue
		}

		result.Id = file.Id
		result.Name = file.Name
		result.Url = "http://" + global.Config.Download.PublicIp +
			":" + strconv.Itoa(global.Config.Access.Port) +
			"/image/" + strconv.FormatInt(file.Id, 10) +
			filepath.Ext(file.Name)
		results = append(results, result)
	}

	var tmpPaging response.Paging
	tmpPaging.Page = 1
	tmpPaging.TotalRecords = int(count)
	tmpPaging.TotalPages = 1

	return results, &tmpPaging, util.Success, nil
}
