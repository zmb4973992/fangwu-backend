package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type FavoriteGet struct {
	Id int64 `json:"-"`
}

type FavoriteCreate struct {
	Creator      int64  `json:"-"`
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

type FavoriteDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type FavoriteGetList struct {
	list
	UserId       int64  `json:"-"`
	BusinessType string `json:"business_type" binding:"required"`
}

type FavoriteResult struct {
	Id int64 `json:"id,omitempty"`
}

func (f *FavoriteGet) Get() (result any, resCode int, errDetail *util.ErrDetail) {
	//获取收藏记录
	var favorite model.Favorite
	err := global.Db.
		Where("id = ?", f.Id).
		First(&favorite).Error
	if err != nil {
		return nil, util.ErrorFailToGetFavorite, util.GetErrDetail(err)
	}

	//如果收藏记录的业务类型为出租
	var forRent model.ForRent
	if favorite.BusinessType == forRent.TableName() {
		var forRent ForRentGet
		forRent.Id = favorite.BusinessId
		result, resCode, errDetail = forRent.Get()
		if resCode != util.Success {
			return nil, resCode, errDetail
		}
		return result, resCode, errDetail
	}

	//如果收藏记录的业务类型为求租
	var seekHouse model.SeekHouse
	if favorite.BusinessType == seekHouse.TableName() {
		var seekHouse SeekHouseGet
		seekHouse.Id = favorite.BusinessId
		result, resCode, errDetail = seekHouse.Get()
		if resCode != util.Success {
			return nil, resCode, errDetail
		}
		return result, resCode, errDetail
	}

	return nil, util.ErrorInvalidBusinessType, nil
}

func (f *FavoriteCreate) Create() (result *FavoriteResult, resCode int, errDetail *util.ErrDetail) {
	var tmpRes FavoriteResult

	//检查是否已收藏
	var favorite model.Favorite
	global.Db.
		Where("creator =?", f.Creator).
		Where("business_type = ?", f.BusinessType).
		Where("business_id = ?", f.BusinessId).
		Limit(1).
		Find(&favorite)

	//如果已收藏，则返回收藏记录
	if favorite.Id != 0 {
		tmpRes.Id = favorite.Id
		return &tmpRes, util.Success, nil
	}

	//创建收藏记录
	favorite.Creator = &f.Creator
	favorite.LastModifier = &f.Creator
	favorite.BusinessType = f.BusinessType
	favorite.BusinessId = f.BusinessId

	err := global.Db.Create(&favorite).Error
	if err != nil {
		return nil, util.ErrorFailToCreateFavorite, util.GetErrDetail(err)
	}

	tmpRes.Id = favorite.Id
	return &tmpRes, util.Success, nil
}

func (f *FavoriteDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	//删除收藏记录
	err := global.Db.
		Where("creator =?", f.Deleter).
		Where("id =?", f.Id).
		Delete(&model.Favorite{}).Error
	if err != nil {
		return util.ErrorFailToDeleteFavorite, util.GetErrDetail(err)
	}

	return util.Success, nil
}

func (f *FavoriteGetList) GetList() (results any, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.Favorite{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("creator = ?", f.UserId).
		Where("business_type = ?", f.BusinessType)

	//找到所有业务id
	var businessIds []int64
	db.Select("business_id").Find(&businessIds)

	//如果收藏记录的业务类型为出租
	var forRent model.ForRent
	if f.BusinessType == forRent.TableName() {
		var param ForRentGetList
		param.Page = f.Page
		param.PageSize = f.PageSize
		param.OrderBy = f.OrderBy
		param.Desc = f.Desc
		param.Ids = businessIds
		results, paging, resCode, errDetail = param.GetList()
		if resCode != util.Success {
			return nil, nil, resCode, errDetail
		}
		return results, paging, resCode, errDetail
	}

	//如果收藏记录的业务类型为求租
	var seekHouse model.SeekHouse
	if f.BusinessType == seekHouse.TableName() {
		var param SeekHouseGetList
		param.Page = f.Page
		param.PageSize = f.PageSize
		param.OrderBy = f.OrderBy
		param.Desc = f.Desc
		param.Ids = businessIds
		results, paging, resCode, errDetail = param.GetList()
		if resCode != util.Success {
			return nil, nil, resCode, errDetail
		}
		return results, paging, resCode, errDetail
	}

	return nil, nil, util.ErrorInvalidBusinessType, nil
}
