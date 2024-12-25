package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
)

type FileDelete struct {
	UserId int64 `json:"-"`

	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

func (f *FileDelete) Archive() (resCode int, errDetail *util.ErrDetail) {
	//找到指定业务id的文件记录
	var files []model.File
	global.Db.Where("creator = ?", f.UserId).
		Where("business_type =?", f.BusinessType).
		Where("business_id = ?", f.BusinessId).
		Find(&files)

	for _, file := range files {
		// 更新文件记录，将业务类型改为 archived_xxxxxx
		err := global.Db.Model(model.File{}).
			Where("id = ?", file.Id).
			Update("business_type", "archived_"+f.BusinessType).Error
		if err != nil {
			return util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
		}
	}

	return util.Success, nil
}
