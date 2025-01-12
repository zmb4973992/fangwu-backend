package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
	"os"
	"path/filepath"
	"strconv"
)

type uploadConfirm struct {
	UserId       int64          `json:"-"`
	BusinessType string         `json:"business_type" binding:"required"`
	BusinessId   int64          `json:"business_id" binding:"required"`
	Files        []UploadResult `json:"files,omitempty"`
}

func (u *uploadConfirm) Confirm() (resCode int, errDetail *util.ErrDetail) {
	//清除之前的关联
	err := global.Db.Model(&model.File{}).
		Where("creator = ?", u.UserId).
		Where("business_id = ?", u.BusinessId).
		Updates(map[string]any{
			"business_type": nil,
			"business_id":   nil}).
		Error
	if err != nil {
		return util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
	}

	//添加新的关联
	for _, f := range u.Files {
		// 检查文件记录是否存在
		var file model.File
		err := global.Db.
			Where("id = ?", f.Id).
			Where("creator = ?", u.UserId).
			First(&file).Error
		if err != nil {
			continue
		}

		// 将文件从临时文件夹剪切到正式文件夹中
		oldDest := filepath.Join(global.Config.Upload.TmpStoragePath,
			strconv.FormatInt(f.Id, 10)+filepath.Ext(file.Name))
		newDest := filepath.Join(global.Config.Upload.StoragePath,
			strconv.FormatInt(f.Id, 10)+filepath.Ext(file.Name))
		_ = os.Rename(oldDest, newDest)

		// 更新文件表
		err = global.Db.Model(&file).
			Where("creator =?", u.UserId).
			Where("id = ?", f.Id).
			Updates(map[string]any{
				"business_type": u.BusinessType,
				"business_id":   u.BusinessId,
				"sort":          f.Sort,
			}).
			Error
		if err != nil {
			return util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
		}
	}

	//删除未使用的文件
	var unusedFiles []model.File
	global.Db.
		Where("creator = ?", u.UserId).
		Where("business_type IS NULL").
		Where("business_id IS NULL").
		Find(&unusedFiles)

	for _, unusedFile := range unusedFiles {
		//删除文件
		dst := filepath.Join(global.Config.Upload.StoragePath,
			strconv.FormatInt(unusedFile.Id, 10)+
				filepath.Ext(unusedFile.Name))
		err = os.Remove(dst)
		//如果删除失败，则跳过该文件
		if err != nil {
			continue
		}

		//如果删除成功，则同时删除文件记录
		err = global.Db.
			Where("id = ?", unusedFile.Id).
			Delete(&unusedFile).Error
		if err != nil {
			return util.ErrorFailToUpdateFileRecord, util.GetErrDetail(err)
		}
	}

	return util.Success, nil
}
