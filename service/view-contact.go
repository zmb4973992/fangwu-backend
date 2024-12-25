package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
	"time"
)

type ViewContactGet struct {
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

type ViewContactCreate struct {
	Creator      int64  `json:"-"`
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

type ViewContactDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type ViewContactCount struct {
	Creator int64 `json:"-"`
}

type ViewContactResult struct {
	Id int64 `json:"id,omitempty"`
}

func (v *ViewContactGet) Get() (result any, resCode int, errDetail *util.ErrDetail) {
	//如果业务类型为出租
	var forRent model.ForRent
	if v.BusinessType == forRent.TableName() {
		var forRent ForRentGetContact
		forRent.Id = v.BusinessId
		result, resCode, errDetail = forRent.GetContact()
		if resCode != util.Success {
			return nil, resCode, errDetail
		}
		return result, resCode, errDetail
	}

	//如果业务类型为求租
	var seekHouse model.SeekHouse
	if v.BusinessType == seekHouse.TableName() {
		var seekHouse SeekHouseGetContact
		seekHouse.Id = v.BusinessId
		result, resCode, errDetail = seekHouse.GetContact()
		if resCode != util.Success {
			return nil, resCode, errDetail
		}
		return result, resCode, errDetail
	}

	return nil, util.ErrorInvalidBusinessType, nil
}

func (v *ViewContactCreate) Create() (result *ViewContactResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var viewContact model.ViewContact

	//创建记录
	viewContact.Creator = &v.Creator
	viewContact.LastModifier = &v.Creator

	viewContact.BusinessType = v.BusinessType
	viewContact.BusinessId = v.BusinessId
	viewContact.Date = time.Now()
	err := tx.FirstOrCreate(&viewContact, viewContact).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateViewContact, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()

	var tmpRes ViewContactResult
	tmpRes.Id = viewContact.Id

	return &tmpRes, util.Success, nil
}

// 删除出租记录：将记录转存到归档表备查，删除相关图片、减小空间占用
func (v *ViewContactDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//获取记录
	var ViewContact model.ViewContact
	err := tx.Where("id = ?", v.Id).
		Where("creator = ?", v.Deleter).
		First(&ViewContact).Error
	if err != nil {
		tx.Rollback()
		return util.Success, nil
	}

	//删除评论记录
	err = tx.Delete(&ViewContact).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteViewContact, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()
	return util.Success, nil
}

func (v *ViewContactCount) Count() (result any, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.ViewContact{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("creator = ?", v.Creator).
		Where("date = ?", time.Now().Format("2006-01-02"))

	// count
	var count int64
	db.Count(&count)

	return map[string]int64{"count": count}, util.Success, nil
}
