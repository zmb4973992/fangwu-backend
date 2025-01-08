package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
)

type DictionaryTypeGet struct {
	Value string `json:"dictionary_type_value" binding:"required"`
}

type dictionaryTypeResult struct {
	Id    int64  `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

func (d *DictionaryTypeGet) Get() (result *dictionaryTypeResult, resCode int, errDetail *util.ErrDetail) {
	err := global.Db.Model(&model.DictionaryType{}).
		Where("value = ?", d.Value).
		First(&result).Error
	if err != nil {
		return nil, util.ErrorFailToGetDictionaryType, util.GetErrDetail(err)
	}

	return result, util.Success, nil
}
