package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
)

type DictionaryTypeGet struct {
	Name string `json:"dictionary_type_name" binding:"required"`
}

type dictionaryTypeResult struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (d *DictionaryTypeGet) Get() (result *dictionaryTypeResult, resCode int, errDetail *util.ErrDetail) {
	err := global.Db.Model(&model.DictionaryType{}).
		Where("name = ?", d.Name).
		First(&result).Error
	if err != nil {
		return nil, util.ErrorFailToGetDictionaryType, util.GetErrDetail(err)
	}

	return result, util.Success, nil
}
