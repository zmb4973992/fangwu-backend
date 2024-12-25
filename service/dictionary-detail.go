package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type dictionaryDetailGet struct {
	Id int64 `json:"id" binding:"required"`
}

type dictionaryDetailGetByName struct {
	DictionaryTypeName   string
	DictionaryDetailName string
}

type DictionaryDetailGetList struct {
	DictionaryTypeId int64 `json:"dictionary_type_id" binding:"required"`
}

type DictionaryDetailResult struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (d *dictionaryDetailGet) Get() (result *DictionaryDetailResult, resCode int, errDetail *util.ErrDetail) {
	var dictionaryDetail model.DictionaryDetail
	err := global.Db.Where("id = ?", d.Id).
		First(&dictionaryDetail).Error
	if err != nil {
		return nil, util.ErrorFailToGetDictionaryDetail, util.GetErrDetail(err)
	}

	var tmpRes DictionaryDetailResult
	tmpRes.Id = dictionaryDetail.Id
	tmpRes.Name = dictionaryDetail.Name

	return &tmpRes, util.Success, nil
}

func (d *dictionaryDetailGetByName) GetByName() (result *DictionaryDetailResult, resCode int, errDetail *util.ErrDetail) {
	var dictionaryType model.DictionaryType
	err := global.Db.Where("name = ?", d.DictionaryTypeName).
		First(&dictionaryType).Error
	if err != nil {
		return nil, util.ErrorFailToGetDictionaryType, util.GetErrDetail(err)
	}

	var dictionaryDetail model.DictionaryDetail
	err = global.Db.Where("dictionary_type_id = ?", dictionaryType.Id).
		Where("name = ?", d.DictionaryDetailName).
		First(&dictionaryDetail).Error
	if err != nil {
		return nil, util.ErrorFailToGetDictionaryDetail, util.GetErrDetail(err)
	}

	var tmpRes DictionaryDetailResult
	tmpRes.Id = dictionaryDetail.Id
	tmpRes.Name = dictionaryDetail.Name

	return &tmpRes, util.Success, nil
}

func (d *DictionaryDetailGetList) GetList() (results []DictionaryDetailResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	var dictionaryDetails []model.DictionaryDetail
	err := global.Db.Where("dictionary_type_id =?", d.DictionaryTypeId).
		Find(&dictionaryDetails).Error
	if err != nil {
		return nil, nil, util.ErrorFailToGetDictionaryDetail, util.GetErrDetail(err)
	}

	for _, v := range dictionaryDetails {
		var result DictionaryDetailResult
		result.Id = v.Id
		result.Name = v.Name
		results = append(results, result)
	}

	//分页信息
	var tmpPaging response.Paging
	tmpPaging.Page = 1
	tmpPaging.PageSize = 0
	tmpPaging.NumberOfRecords = len(results)
	tmpPaging.NumberOfPages = 1

	return results, &tmpPaging, util.Success, nil
}
