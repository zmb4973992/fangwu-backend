package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
	"time"
)

type SeekHouseGet struct {
	Id int64 `json:"-"`
}

type SeekHouseGetContact struct {
	Id int64 `json:"-"`
}

type SeekHouseCreate struct {
	Creator                 int64   `json:"-"`
	MaxBudget               float64 `json:"max_budget,omitempty"`
	MinBudget               float64 `json:"min_budget,omitempty"`
	RentType                int64   `json:"rent_type" binding:"required"`
	Description             string  `json:"description" binding:"required"`
	GenderRestriction       int64   `json:"gender_restriction" binding:"required"`
	MobilePhone             string  `json:"mobile_phone,omitempty"`
	WeChatId                string  `json:"wechat_id,omitempty"`
	FileIds                 []int64 `json:"file_ids,omitempty"`
	AdministrativeDivision1 int     `json:"administrative_division_1,omitempty"`
	AdministrativeDivision2 int     `json:"administrative_division_2,omitempty"`
	AdministrativeDivision3 int     `json:"administrative_division_3,omitempty"`
	AdministrativeDivision4 int     `json:"administrative_division_4,omitempty"`
	Community               string  `json:"community,omitempty"`
}

type SeekHouseUpdate struct {
	LastModifier            int64   `json:"-"`
	Id                      int64   `json:"id" binding:"required"`
	MaxBudget               float64 `json:"max_budget,omitempty"`
	MinBudget               float64 `json:"min_budget,omitempty"`
	RentType                int64   `json:"rent_type,omitempty"`
	Description             string  `json:"description,omitempty"`
	GenderRestriction       int64   `json:"gender_restriction,omitempty"`
	MobilePhone             string  `json:"mobile_phone,omitempty"`
	WeChatId                string  `json:"wechat_id,omitempty"`
	FileIds                 []int64 `json:"file_ids,omitempty"`
	AdministrativeDivision1 int     `json:"administrative_division_1,omitempty"`
	AdministrativeDivision2 int     `json:"administrative_division_2,omitempty"`
	AdministrativeDivision3 int     `json:"administrative_division_3,omitempty"`
	AdministrativeDivision4 int     `json:"administrative_division_4,omitempty"`
	Community               string  `json:"community,omitempty"`
}

type SeekHouseDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type SeekHouseGetList struct {
	list
	RentType          int64   `json:"rent_type,omitempty"`
	MaxBudget         float64 `json:"max_budget,omitempty"`
	MinBudget         float64 `json:"min_budget,omitempty"`
	GenderRestriction int64   `json:"gender_restriction,omitempty"`
	Ids               []int64 `json:"-"`
	Keyword           string  `json:"keyword,omitempty"`
	Community         string  `json:"community,omitempty"`
}

type SeekHouseResult struct {
	Creator                 int64                         `json:"creator,omitempty"`
	LastModifier            int64                         `json:"last_modifier,omitempty"`
	IsDeleted               bool                          `json:"is_deleted,omitempty"`
	Id                      int64                         `json:"id,omitempty"`
	RentType                *DictionaryDetailResult       `json:"rent_type,omitempty"`
	MaxBudget               float64                       `json:"max_budget,omitempty"`
	MinBudget               float64                       `json:"min_budget,omitempty"`
	Description             string                        `json:"description,omitempty"`
	GenderRestriction       *DictionaryDetailResult       `json:"gender_restriction,omitempty"`
	MobilePhone             string                        `json:"mobile_phone,omitempty"`
	WeChatId                string                        `json:"wechat_id,omitempty"`
	Files                   []ImageResult                 `json:"files,omitempty"`
	AdministrativeDivision1 *administrativeDivisionResult `json:"administrative_division_1,omitempty"`
	AdministrativeDivision2 *administrativeDivisionResult `json:"administrative_division_2,omitempty"`
	AdministrativeDivision3 *administrativeDivisionResult `json:"administrative_division_3,omitempty"`
	AdministrativeDivision4 *administrativeDivisionResult `json:"administrative_division_4,omitempty"`
	Community               string                        `json:"community,omitempty"`
}

func (s *SeekHouseGet) Get() (result *SeekHouseResult, resCode int, errDetail *util.ErrDetail) {
	//获取求租记录
	var seekHouse model.SeekHouse
	err := global.Db.
		Where("id = ?", s.Id).
		First(&seekHouse).Error
	if err != nil {
		return nil, util.ErrorFailToGetSeekHouse, util.GetErrDetail(err)
	}

	var tmpRes SeekHouseResult

	// 填充结果
	if seekHouse.Creator != nil {
		tmpRes.Creator = *seekHouse.Creator
	}

	if seekHouse.LastModifier != nil {
		tmpRes.LastModifier = *seekHouse.LastModifier
	}

	tmpRes.IsDeleted = seekHouse.IsDeleted
	tmpRes.Id = seekHouse.Id

	tmpRes.MaxBudget = seekHouse.MaxBudget
	tmpRes.MinBudget = seekHouse.MinBudget

	//填充租赁类型
	var rentType dictionaryDetailGet
	rentType.Id = seekHouse.RentType
	tmpRes.RentType, _, _ = rentType.Get()

	tmpRes.Description = seekHouse.Description

	// 填充性别限制
	var genderRestriction dictionaryDetailGet
	genderRestriction.Id = seekHouse.GenderRestriction
	tmpRes.GenderRestriction, _, _ = genderRestriction.Get()

	//获取文件下载地址
	var download imageGetList
	download.businessType = seekHouse.TableName()
	download.businessId = seekHouse.Id
	tmpRes.Files, _, _, _ = download.GetList()

	//获取行政区划
	if seekHouse.AdministrativeDivision1 != nil {
		var administrativeDivision1 administrativeDivisionGet
		administrativeDivision1.Code = *seekHouse.AdministrativeDivision1
		tmpRes.AdministrativeDivision1, _, _ = administrativeDivision1.Get()
	}
	if seekHouse.AdministrativeDivision2 != nil {
		var administrativeDivision2 administrativeDivisionGet
		administrativeDivision2.Code = *seekHouse.AdministrativeDivision2
		tmpRes.AdministrativeDivision2, _, _ = administrativeDivision2.Get()
	}
	if seekHouse.AdministrativeDivision3 != nil {
		var administrativeDivision3 administrativeDivisionGet
		administrativeDivision3.Code = *seekHouse.AdministrativeDivision3
		tmpRes.AdministrativeDivision3, _, _ = administrativeDivision3.Get()
	}
	if seekHouse.AdministrativeDivision4 != nil {
		var administrativeDivision4 administrativeDivisionGet
		administrativeDivision4.Code = *seekHouse.AdministrativeDivision4
		tmpRes.AdministrativeDivision4, _, _ = administrativeDivision4.Get()
	}

	//小区
	tmpRes.Community = seekHouse.Community

	return &tmpRes, util.Success, nil
}

func (s *SeekHouseGetContact) GetContact() (result *SeekHouseResult, resCode int, errDetail *util.ErrDetail) {
	//获取求租记录
	var seekHouse model.SeekHouse
	err := global.Db.
		Where("id = ?", s.Id).
		First(&seekHouse).Error
	if err != nil {
		return nil, util.ErrorFailToGetSeekHouse, util.GetErrDetail(err)
	}

	var tmpRes SeekHouseResult

	if seekHouse.MobilePhone != nil {
		tmpRes.MobilePhone = *seekHouse.MobilePhone
	}

	if seekHouse.WeChatId != nil {
		tmpRes.WeChatId = *seekHouse.WeChatId
	}

	return &tmpRes, util.Success, nil
}

func (s *SeekHouseCreate) Create() (result *SeekHouseResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var seekHouse model.SeekHouse
	seekHouse.Creator = &s.Creator
	seekHouse.LastModifier = &s.Creator
	seekHouse.IsDeleted = false
	seekHouse.MaxBudget = s.MaxBudget
	seekHouse.MinBudget = s.MinBudget
	seekHouse.RentType = s.RentType
	seekHouse.Description = s.Description
	seekHouse.GenderRestriction = s.GenderRestriction
	if s.MobilePhone != "" {
		seekHouse.MobilePhone = model.StringToPointer(s.MobilePhone)
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "mobile_phone"
		contactInfoBlacklist.Value = s.MobilePhone
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorMobilePhoneIsInBlacklist, nil
		}
	}
	if s.WeChatId != "" {
		seekHouse.WeChatId = model.StringToPointer(s.WeChatId)
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "wechat_id"
		contactInfoBlacklist.Value = s.MobilePhone
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorMobilePhoneIsInBlacklist, nil
		}
	}

	//行政区划
	if s.AdministrativeDivision1 > 0 {
		seekHouse.AdministrativeDivision1 = &s.AdministrativeDivision1
	}
	if s.AdministrativeDivision2 > 0 {
		seekHouse.AdministrativeDivision2 = &s.AdministrativeDivision2
	}
	if s.AdministrativeDivision3 > 0 {
		seekHouse.AdministrativeDivision3 = &s.AdministrativeDivision3
	}
	if s.AdministrativeDivision4 > 0 {
		seekHouse.AdministrativeDivision4 = &s.AdministrativeDivision4
	}

	//小区
	seekHouse.Community = s.Community

	err := tx.Create(&seekHouse).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateSeekHouse, util.GetErrDetail(err)
	}

	//批量确认上传
	var file FileBatchConfirm
	file.FileIds = s.FileIds
	file.UserId = s.Creator
	file.BusinessType = seekHouse.TableName()
	file.BusinessId = seekHouse.Id
	resCode, errDetail = file.BatchConfirm()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes SeekHouseResult
	tmpRes.Id = seekHouse.Id

	return &tmpRes, util.Success, nil
}

func (s *SeekHouseUpdate) Update() (result *SeekHouseResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	seekHouse := make(map[string]any)

	if s.LastModifier > 0 {
		seekHouse["last_modifier"] = s.LastModifier
	}

	if s.MaxBudget > 0 {
		seekHouse["max_budget"] = s.MaxBudget
	}

	if s.MinBudget > 0 {
		seekHouse["min_budget"] = s.MinBudget
	}

	if s.RentType > 0 {
		seekHouse["rent_type"] = s.RentType
	}

	if s.Description != "" {
		seekHouse["description"] = s.Description
	}

	if s.GenderRestriction > 0 {
		seekHouse["gender_restriction"] = s.GenderRestriction
	}

	if s.MobilePhone != "" {
		seekHouse["mobile_phone"] = s.MobilePhone
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "mobile_phone"
		contactInfoBlacklist.Value = s.MobilePhone
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorMobilePhoneIsInBlacklist, nil
		}
	}

	if s.WeChatId != "" {
		seekHouse["wechat_id"] = s.WeChatId
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "wechat_id"
		contactInfoBlacklist.Value = s.MobilePhone
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorMobilePhoneIsInBlacklist, nil
		}
	}

	//行政区划
	if s.AdministrativeDivision1 > 0 {
		seekHouse["administrative_division_1"] = s.AdministrativeDivision1
	}
	if s.AdministrativeDivision2 > 0 {
		seekHouse["administrative_division_2"] = s.AdministrativeDivision2
	}
	if s.AdministrativeDivision3 > 0 {
		seekHouse["administrative_division_3"] = s.AdministrativeDivision3
	}
	if s.AdministrativeDivision4 > 0 {
		seekHouse["administrative_division_4"] = s.AdministrativeDivision4
	}

	//小区
	if s.Community != "" {
		seekHouse["community"] = s.Community
	}

	err := tx.Model(&model.SeekHouse{}).
		Where("id =?", s.Id).
		Where("creator = ?", s.LastModifier).
		Where("is_deleted = ?", false).
		Updates(seekHouse).Error
	if err != nil {
		return nil, util.ErrorFailToUpdateSeekHouse, util.GetErrDetail(err)
	}

	//批量确认上传
	var uploadBatchConfirm FileBatchConfirm
	uploadBatchConfirm.FileIds = s.FileIds
	uploadBatchConfirm.UserId = s.LastModifier
	uploadBatchConfirm.BusinessType = "seek_house"
	uploadBatchConfirm.BusinessId = s.Id
	resCode, errDetail = uploadBatchConfirm.BatchConfirm()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes SeekHouseResult
	tmpRes.Id = s.Id

	return &tmpRes, util.Success, nil
}

// 删除出租记录：将记录转存到归档表备查，删除相关图片、减小空间占用
func (s *SeekHouseDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//获取求租记录
	var seekHouse model.SeekHouse
	err := tx.Where("id = ?", s.Id).
		Where("creator = ?", s.Deleter).
		Where("is_deleted = ?", false).
		First(&seekHouse).Error
	if err != nil {
		tx.Rollback()
		return util.Success, nil
	}

	//软删除求租记录
	err = tx.Model(&seekHouse).
		Updates(map[string]any{
			"is_deleted": true,
			"deleter":    s.Deleter,
			"deleted_at": time.Now(),
		}).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteSeekHouse, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()
	return util.Success, nil
}

func (s *SeekHouseGetList) GetList() (results []SeekHouseResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.SeekHouse{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("is_deleted = ?", false)

	if s.RentType > 0 {
		db = db.Where("rent_type = ?", s.RentType)
	}
	if s.GenderRestriction > 0 {
		db = db.Where("gender_restriction = ?", s.GenderRestriction)
	}

	//下面两个条件没写错，是为了取数据范围的交集
	//如果有max_budget，那就取数据库中最大预算小于等于max_budget的记录
	if s.MaxBudget > 0 {
		db = db.Where("max_budget >= ?", s.MaxBudget)
	}
	//如果有min_budget，那就取数据库中最大预算大于等于min_budget的记录
	if s.MinBudget > 0 {
		db = db.Where("max_budget >= ?", s.MinBudget)
	}

	if s.Keyword != "" {
		db = db.Where("description like ?", "%"+s.Keyword+"%")
	}

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if s.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if s.Desc {
			db = db.Order("id desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.SeekHouse
		ok := util.FieldIsInModel(db, tmp.TableName(), s.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if s.Desc {
			db = db.Order(s.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(s.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if s.PageSize > 0 && s.PageSize <= maxPageSize {
		pageSize = s.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if s.Page > 0 {
		page = s.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var seekHouses []model.SeekHouse
	db.Find(&seekHouses)

	//将结果转换为Result
	for _, seekHouse := range seekHouses {
		var result SeekHouseResult
		if seekHouse.Creator != nil {
			result.Creator = *seekHouse.Creator
		}
		if seekHouse.LastModifier != nil {
			result.LastModifier = *seekHouse.LastModifier
		}
		result.Id = seekHouse.Id

		//获取租赁类型
		var rentType dictionaryDetailGet
		rentType.Id = seekHouse.RentType
		result.RentType, _, _ = rentType.Get()

		result.MaxBudget = seekHouse.MaxBudget
		result.MinBudget = seekHouse.MinBudget
		result.Description = seekHouse.Description

		//获取性别限制
		var genderRestriction dictionaryDetailGet
		genderRestriction.Id = seekHouse.GenderRestriction
		result.GenderRestriction, _, _ = genderRestriction.Get()

		if seekHouse.MobilePhone != nil {
			moblelePhone := *seekHouse.MobilePhone
			result.MobilePhone = moblelePhone[:(len(*seekHouse.MobilePhone)-2)] + "**"
		}

		if seekHouse.WeChatId != nil {
			wechatId := *seekHouse.WeChatId
			result.WeChatId = wechatId[:(len(*seekHouse.WeChatId)-2)] + "**"
		}

		//获取文件下载地址
		var download imageGetList
		download.businessType = seekHouse.TableName()
		download.businessId = seekHouse.Id
		result.Files, _, _, _ = download.GetList()

		//获取行政区划
		if seekHouse.AdministrativeDivision1 != nil {
			var administrativeDivision administrativeDivisionGet
			administrativeDivision.Code = *seekHouse.AdministrativeDivision1
			result.AdministrativeDivision1, _, _ = administrativeDivision.Get()
		}
		if seekHouse.AdministrativeDivision2 != nil {
			var administrativeDivision administrativeDivisionGet
			administrativeDivision.Code = *seekHouse.AdministrativeDivision2
			result.AdministrativeDivision2, _, _ = administrativeDivision.Get()
		}
		if seekHouse.AdministrativeDivision3 != nil {
			var administrativeDivision administrativeDivisionGet
			administrativeDivision.Code = *seekHouse.AdministrativeDivision3
			result.AdministrativeDivision3, _, _ = administrativeDivision.Get()
		}
		if seekHouse.AdministrativeDivision4 != nil {
			var administrativeDivision administrativeDivisionGet
			administrativeDivision.Code = *seekHouse.AdministrativeDivision4
			result.AdministrativeDivision4, _, _ = administrativeDivision.Get()
		}

		//小区
		result.Community = seekHouse.Community

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
