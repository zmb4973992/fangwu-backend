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
	list
	ParentCode  int64 `json:"parent_code,omitempty"`
	Level       int   `json:"level,omitempty"`
	GrandpaCode int64 `json:"grandpa_code,omitempty"`
}

type AdminDivResult struct {
	Code         int    `json:"code,omitempty"`
	Name         string `json:"name,omitempty"`
	PinyinPrefix string `json:"pinyin_prefix,omitempty"`
	ParentCode   int    `json:"parent_code,omitempty"`
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
	tmpRes.PinyinPrefix = adminDiv.PinyinPrefix
	tmpRes.ParentCode = adminDiv.ParentCode

	return &tmpRes, util.Success, nil
}

func (a *AdminDivGetList) GetList() (results []AdminDivResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.AdminDiv{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	//where
	if a.ParentCode > 0 {
		db = db.Where("parent_code = ?", a.ParentCode)
	}
	if a.Level > 0 {
		db = db.Where("level = ?", a.Level)
	}
	if a.GrandpaCode > 0 {
		db = db.Where("parent_code in (select code from administrative_division where parent_code = ?)", a.GrandpaCode)
	}

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if a.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if a.Desc {
			db = db.Order("code desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.AdminDiv
		ok := util.FieldIsInModel(db, tmp.TableName(), a.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if a.Desc {
			db = db.Order(a.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(a.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if a.PageSize > 0 && a.PageSize <= maxPageSize {
		pageSize = a.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if a.Page > 0 {
		page = a.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var adminDivs []model.AdminDiv
	db.Find(&adminDivs)

	for _, v := range adminDivs {
		var result AdminDivResult
		result.Code = v.Code
		result.Name = v.Name
		result.ParentCode = v.ParentCode
		result.PinyinPrefix = v.PinyinPrefix
		results = append(results, result)
	}

	//分页信息
	var tmpPaging response.Paging
	tmpPaging.Page = 1
	tmpPaging.PageSize = 0
	tmpPaging.TotalRecords = len(results)
	tmpPaging.TotalPages = 1

	return results, &tmpPaging, util.Success, nil
}
