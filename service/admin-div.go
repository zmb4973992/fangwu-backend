package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type adminDivGet struct {
	Code int `json:"code" binding:"required"`
}

type AdminDivGetList struct {
	ParentCode int64 `json:"parent_code" binding:"required"`
}

type AdminDivResult struct {
	Code         int    `json:"code,omitempty"`
	Name         string `json:"name,omitempty"`
	PinyinPrefix string `json:"pinyin_prefix,omitempty"`
}

func (a *adminDivGet) Get() (result *AdminDivResult, resCode int, errDetail *util.ErrDetail) {
	var adminDiv model.AdminDiv
	err := global.Db.Where("code = ?", a.Code).
		First(&adminDiv).Error
	if err != nil {
		return nil, util.ErrorFailToGetAdminDiv, util.GetErrDetail(err)
	}

	var tmpRes AdminDivResult
	tmpRes.Code = adminDiv.Code
	tmpRes.Name = adminDiv.Name

	return &tmpRes, util.Success, nil
}

func (a *AdminDivGetList) GetList() (results []AdminDivResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	var adminDivs []model.AdminDiv
	err := global.Db.Where("parent_code = ?", a.ParentCode).
		Find(&adminDivs).Error
	if err != nil {
		return nil, nil, util.ErrorFailToGetAdminDiv, util.GetErrDetail(err)
	}

	for _, v := range adminDivs {
		var result AdminDivResult
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
