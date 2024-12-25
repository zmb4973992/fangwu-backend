package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type UserBlackListVerify struct {
	Blocker int64 `json:"blocker" bind:"required"`
	Blocked int64 `json:"blocked" bind:"required"`
}

type UserBlacklistCreate struct {
	Blocker int64 `json:"-"`
	Blocked int64 `json:"blocked" bind:"required"`
}

type UserBlacklistDelete struct {
	Blocker int64 `json:"-"`
	Blocked int64 `json:"blocked" bind:"required"`
}

type UserBlacklistGetList struct {
	list
	Blocker int64 `json:"-"`
}

type UserBlacklistResult struct {
	Id      int64 `json:"id,omitempty"`
	Blocked int64 `json:"blocked,omitempty"`
}

func (u *UserBlackListVerify) Verify() (isBlocked bool, resCode int, errDetail *util.ErrDetail) {
	var count int64
	global.Db.Model(&model.UserBlacklist{}).
		Where("blocker = ? and blocked = ?", u.Blocker, u.Blocked).
		Count(&count)
	if count > 0 {
		return true, util.Success, nil
	}

	return false, util.Success, nil
}

func (u *UserBlacklistCreate) Create() (result *UserBlacklistResult, resCode int, errDetail *util.ErrDetail) {
	var count int64
	global.Db.Model(&model.UserBlacklist{}).
		Where("blocker = ? and blocked = ?", u.Blocker, u.Blocked).
		Count(&count)
	if count > 0 {
		return nil, util.Success, nil
	}

	var userBlacklist model.UserBlacklist
	userBlacklist.Creator = &u.Blocker
	userBlacklist.LastModifier = &u.Blocker
	userBlacklist.Blocker = u.Blocker
	userBlacklist.Blocked = u.Blocked

	err := global.Db.Create(&userBlacklist).Error
	if err != nil {
		return nil, util.ErrorFailToCreateUserBlacklist, util.GetErrDetail(err)
	}

	var userBlacklistResult UserBlacklistResult
	userBlacklistResult.Id = userBlacklist.Id

	return &userBlacklistResult, util.Success, nil
}

func (u *UserBlacklistDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	err := global.Db.
		Where("blocker =? and blocked =?", u.Blocker, u.Blocked).
		Delete(&model.UserBlacklist{}).Error
	if err != nil {
		return util.ErrorFailToDeleteUserBlacklist, util.GetErrDetail(err)
	}
	return util.Success, nil
}

func (u *UserBlacklistGetList) GetList() (results []UserBlacklistResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.UserBlacklist{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("blocker = ?", u.Blocker)

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if u.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if u.Desc {
			db = db.Order("id desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.UserBlacklist
		ok := util.FieldIsInModel(db, tmp.TableName(), u.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if u.Desc {
			db = db.Order(u.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(u.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if u.PageSize > 0 && u.PageSize <= maxPageSize {
		pageSize = u.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if u.Page > 0 {
		page = u.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var userBlacklists []model.UserBlacklist
	db.Find(&userBlacklists)

	//将结果转换为Result
	for _, userBlacklist := range userBlacklists {
		var result UserBlacklistResult
		result.Blocked = userBlacklist.Blocked

		results = append(results, result)
	}

	//分页信息
	var tmpPaging response.Paging
	tmpPaging.Page = page
	tmpPaging.PageSize = pageSize
	tmpPaging.NumberOfRecords = int(count)
	tmpPaging.NumberOfPages = util.GetNumberOfPages(int(count), pageSize)

	return results, &tmpPaging, util.Success, nil
}
