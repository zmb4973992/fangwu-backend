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

type administrativeDivisionGetList struct {
	ParentCode int64 `json:"parent_code" binding:"required"`
}

type administrativeDivisionResult struct {
	Code int    `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *administrativeDivisionGet) Get() (result *administrativeDivisionResult, resCode int, errDetail *util.ErrDetail) {
	var administrativeDivision model.AdministrativeDivision
	err := global.Db.Where("code = ?", a.Code).
		First(&administrativeDivision).Error
	if err != nil {
		return nil, util.ErrorFailToGetAdministrativeDivision, util.GetErrDetail(err)
	}

	var tmpRes administrativeDivisionResult
	tmpRes.Code = administrativeDivision.Code
	tmpRes.Name = administrativeDivision.Name

	return &tmpRes, util.Success, nil
}

func (a *administrativeDivisionGetList) GetList() (results []administrativeDivisionResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	var administrativeDivisions []model.AdministrativeDivision
	err := global.Db.Where("parent_code =?", a.ParentCode).
		Find(&administrativeDivisions).Error
	if err != nil {
		return nil, nil, util.ErrorFailToGetAdministrativeDivision, util.GetErrDetail(err)
	}

	for _, v := range administrativeDivisions {
		var result administrativeDivisionResult
		result.Code = v.Code
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
