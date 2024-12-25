package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type UploadCreate struct {
	UserId     int64
	FileHeader *multipart.FileHeader
}

type UploadBatchCreate struct {
	UserId     int64
	FileHeader []*multipart.FileHeader
}

type UploadResult struct {
	FileId int64  `json:"file_id,omitempty"`
	Url    string `json:"url,omitempty"`
}

func (u *UploadCreate) Create() (result *UploadResult, resCode int, errDetail *util.ErrDetail) {
	//读取临时的存储路径
	tmpStoragePath := global.Config.Upload.TmpStoragePath

	// 创建文件记录
	file := model.File{
		Base:   model.Base{Creator: &u.UserId},
		Name:   u.FileHeader.Filename,
		SizeKb: math.Round(float64(u.FileHeader.Size)/1024*100) / 100,
	}

	err := global.Db.Create(&file).Error
	if err != nil {
		return nil, util.ErrorFailToCreateFileRecord, util.GetErrDetail(err)
	}

	// 生成新的文件名
	newFileName := strconv.FormatInt(file.Id, 10) + filepath.Ext(u.FileHeader.Filename)

	//生成文件路径
	destination := filepath.Join(tmpStoragePath, newFileName)

	// 保存文件
	resCode, errDetail = saveFile(u.FileHeader, destination)
	if resCode != util.Success {
		return nil, resCode, errDetail
	}

	//如果是png图片
	if strings.HasSuffix(u.FileHeader.Filename, ".png") {
		//将png图片转换为jpg图片
		newDestination, resCode, errDetail := pngToJpg(destination)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}

		//将jpg图片裁剪、压缩、保存
		resCode, errDetail = jpgResize(newDestination, 900, 80)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}

		//删除png图片
		err := os.Remove(destination)
		if err != nil {
			return nil, util.ErrorFailToDeleteFile, util.GetErrDetail(err)
		}

		//更新file表的文件名
		err = global.Db.Model(&file).
			Where("id = ?", file.Id).
			Update("name", strings.Trim(file.Name, ".png")+".jpg").
			Error
		if err != nil {
			return nil, util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
		}
	}

	//如果是bmp图片
	if strings.HasSuffix(u.FileHeader.Filename, ".bmp") {
		//将bmp图片转换为jpg图片
		newDestination, resCode, errDetail := bmpToJpg(destination)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}

		//将jpg图片裁剪、压缩、保存
		resCode, errDetail = jpgResize(newDestination, 900, 80)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}

		//删除bmp文件
		err := os.Remove(destination)
		if err != nil {
			return nil, util.ErrorFailToDeleteFile, util.GetErrDetail(err)
		}

		//更新file表的文件名
		err = global.Db.Model(&file).
			Where("id =?", file.Id).
			Update("name", strings.Trim(file.Name, ".bmp")+".jpg").Error
		if err != nil {
			return nil, util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
		}
	}

	//如果是jpg、jpeg图片
	if strings.HasSuffix(u.FileHeader.Filename, ".jpg") ||
		strings.HasSuffix(u.FileHeader.Filename, ".jpeg") {
		//将jpg、jpeg图片裁剪、压缩、保存
		resCode, errDetail = jpgResize(destination, 900, 80)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}
	}

	result = &UploadResult{FileId: file.Id}

	return result, util.Success, nil
}

func (u *UploadBatchCreate) BatchCreate() (result []*UploadResult, resCode int, errDetail *util.ErrDetail) {
	for _, fileHeader := range u.FileHeader {
		//读取临时的存储路径
		tmpStoragePath := global.Config.Upload.TmpStoragePath

		// 创建文件记录
		file := model.File{
			Base:   model.Base{Creator: &u.UserId},
			Name:   fileHeader.Filename,
			SizeKb: math.Round(float64(fileHeader.Size)/1024*100) / 100,
		}

		err := global.Db.Create(&file).Error
		if err != nil {
			return nil, util.ErrorFailToCreateFileRecord, util.GetErrDetail(err)
		}

		// 生成新的文件名
		newFileName := strconv.FormatInt(file.Id, 10) + filepath.Ext(fileHeader.Filename)

		//生成文件路径
		destination := filepath.Join(tmpStoragePath, newFileName)

		// 保存文件
		resCode, errDetail = saveFile(fileHeader, destination)
		if resCode != util.Success {
			return nil, resCode, errDetail
		}

		tmpResult := &UploadResult{FileId: file.Id}
		result = append(result, tmpResult)

		//如果是png图片
		if strings.HasSuffix(fileHeader.Filename, ".png") {
			//将png图片转换为jpg图片
			newDestination, resCode, errDetail := pngToJpg(destination)
			if resCode != util.Success {
				return nil, resCode, errDetail
			}

			//将jpg图片裁剪、压缩、保存
			resCode, errDetail = jpgResize(newDestination, 900, 80)
			if resCode != util.Success {
				return nil, resCode, errDetail
			}

			//删除png文件
			err := os.Remove(destination)
			if err != nil {
				return nil, util.ErrorFailToDeleteFile, util.GetErrDetail(err)
			}

			//更新file表的文件名
			err = global.Db.Model(&file).
				Where("id = ?", file.Id).
				Update("name", strings.Trim(file.Name, ".png")+".jpg").
				Error
			if err != nil {
				return nil, util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
			}
		}

		//如果是bmp图片
		if strings.HasSuffix(fileHeader.Filename, ".bmp") {
			//将bmp图片转换为jpg图片
			newDestination, resCode, errDetail := bmpToJpg(destination)
			if resCode != util.Success {
				return nil, resCode, errDetail
			}

			//将jpg图片裁剪、压缩、保存
			resCode, errDetail = jpgResize(newDestination, 900, 80)
			if resCode != util.Success {
				return nil, resCode, errDetail
			}

			//删除bmp文件
			err := os.Remove(destination)
			if err != nil {
				return nil, util.ErrorFailToDeleteFile, util.GetErrDetail(err)
			}

			//更新file表的文件名
			err = global.Db.Model(&file).
				Where("id =?", file.Id).
				Update("name", strings.Trim(file.Name, ".bmp")+".jpg").
				Error
			if err != nil {
				return nil, util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
			}
		}

		//如果是jpg图片
		if strings.HasSuffix(fileHeader.Filename, ".jpg") ||
			strings.HasSuffix(fileHeader.Filename, ".jpeg") {
			//将jpg图片裁剪、压缩、保存
			resCode, errDetail = jpgResize(destination, 900, 80)
			if resCode != util.Success {
				return nil, resCode, errDetail
			}
		}
	}

	return result, util.Success, nil
}
