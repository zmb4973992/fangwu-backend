package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type administrativeDivisionGet struct {
	Code int `json:"code" binding:"required"`
}

type AdministrativeDivisionGetList struct {
	ParentCode int64 `json:"parent_code" binding:"required"`
}

type AdministrativeDivisionResult struct {
	Code         int    `json:"code,omitempty"`
	Name         string `json:"name,omitempty"`
	PinyinPrefix string `json:"pinyin_prefix,omitempty"`
}

func (a *administrativeDivisionGet) Get() (result *AdministrativeDivisionResult, resCode int, errDetail *util.ErrDetail) {
	var administrativeDivision model.AdministrativeDivision
	err := global.Db.Where("code = ?", a.Code).
		First(&administrativeDivision).Error
	if err != nil {
		return nil, util.ErrorFailToGetAdministrativeDivision, util.GetErrDetail(err)
	}

	var tmpRes AdministrativeDivisionResult
	tmpRes.Code = administrativeDivision.Code
	tmpRes.Name = administrativeDivision.Name

	return &tmpRes, util.Success, nil
}

func (a *AdministrativeDivisionGetList) GetList() (results []AdministrativeDivisionResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	var administrativeDivisions []model.AdministrativeDivision
	err := global.Db.Where("parent_code = ?", a.ParentCode).
		Find(&administrativeDivisions).Error
	if err != nil {
		return nil, nil, util.ErrorFailToGetAdministrativeDivision, util.GetErrDetail(err)
	}

	for _, v := range administrativeDivisions {
		var result AdministrativeDivisionResult
		result.Code = v.Code
		result.Name = v.Name
		result.PinyinPrefix = v.PinyinPrefix
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
