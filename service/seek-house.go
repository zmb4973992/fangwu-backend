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
	Creator           int64   `json:"-"`
	MaxBudget         float64 `json:"max_budget,omitempty"`
	MinBudget         float64 `json:"min_budget,omitempty"`
	RentType          int64   `json:"rent_type" binding:"required"`
	Description       string  `json:"description" binding:"required"`
	GenderRestriction int64   `json:"gender_restriction" binding:"required"`
	MobilePhone       string  `json:"mobile_phone,omitempty"`
	WeChatId          string  `json:"wechat_id,omitempty"`
	FileIds           []int64 `json:"file_ids,omitempty"`
	Level1AdminDiv    int     `json:"level_1_admin_div,omitempty"`
	Level2AdminDiv    int     `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    int     `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    int     `json:"level_4_admin_div,omitempty"`
	Community         string  `json:"community,omitempty"`
	Area              int     `json:"area,omitempty"`
	Name              string  `json:"name,omitempty"`
	Gender            int64   `json:"gender,omitempty"`
}

type SeekHouseUpdate struct {
	LastModifier      int64   `json:"-"`
	Id                int64   `json:"id" binding:"required"`
	MaxBudget         float64 `json:"max_budget,omitempty"`
	MinBudget         float64 `json:"min_budget,omitempty"`
	RentType          int64   `json:"rent_type,omitempty"`
	Description       string  `json:"description,omitempty"`
	GenderRestriction int64   `json:"gender_restriction,omitempty"`
	MobilePhone       string  `json:"mobile_phone,omitempty"`
	WeChatId          string  `json:"wechat_id,omitempty"`
	FileIds           []int64 `json:"file_ids,omitempty"`
	Level1AdminDiv    int     `json:"level_1_admin_div,omitempty"`
	Level2AdminDiv    int     `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    int     `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    int     `json:"level_4_admin_div,omitempty"`
	Community         string  `json:"community,omitempty"`
	Area              int     `json:"area,omitempty"`
	Name              string  `json:"name,omitempty"`
	Gender            int64   `json:"gender,omitempty"`
}

type SeekHouseDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type SeekHouseGetList struct {
	list
	Creator           int64   `json:"-"`
	CreatedByMyself   *bool   `json:"created_by_myself,omitempty"`
	RentType          int64   `json:"rent_type,omitempty"`
	MaxBudget         float64 `json:"max_budget,omitempty"`
	MinBudget         float64 `json:"min_budget,omitempty"`
	GenderRestriction int64   `json:"gender_restriction,omitempty"`
	Ids               []int64 `json:"-"`
	Keyword           string  `json:"keyword,omitempty"`
	Community         string  `json:"community,omitempty"`
	Level2AdminDiv    int     `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    int     `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    int     `json:"level_4_admin_div,omitempty"`
}

type SeekHouseResult struct {
	Creator           int64                   `json:"creator,omitempty"`
	LastModifier      int64                   `json:"last_modifier,omitempty"`
	IsDeleted         bool                    `json:"is_deleted,omitempty"`
	Id                int64                   `json:"id,omitempty"`
	RentType          *DictionaryDetailResult `json:"rent_type,omitempty"`
	MaxBudget         float64                 `json:"max_budget,omitempty"`
	MinBudget         float64                 `json:"min_budget,omitempty"`
	Description       string                  `json:"description,omitempty"`
	GenderRestriction *DictionaryDetailResult `json:"gender_restriction,omitempty"`
	MobilePhone       string                  `json:"mobile_phone,omitempty"`
	WeChatId          string                  `json:"wechat_id,omitempty"`
	Files             []ImageResult           `json:"files,omitempty"`
	Level1AdminDiv    *AdminDivResult         `json:"level_1_admin_div,omitempty"`
	Level2AdminDiv    *AdminDivResult         `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    *AdminDivResult         `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    *AdminDivResult         `json:"level_4_admin_div,omitempty"`
	Community         string                  `json:"community,omitempty"`
	Area              int                     `json:"area,omitempty"`
	Name              string                  `json:"name,omitempty"`
	Gender            *DictionaryDetailResult `json:"gender,omitempty"`
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

	//获取文件信息
	var download imageGetList
	download.businessType = seekHouse.TableName()
	download.businessId = seekHouse.Id
	tmpRes.Files, _, _, _ = download.GetList()

	//获取行政区划
	if seekHouse.Level1AdminDiv != nil {
		var level1AdminDiv adminDivGetByCode
		level1AdminDiv.Code = *seekHouse.Level1AdminDiv
		tmpRes.Level1AdminDiv, _, _ = level1AdminDiv.Get()
	}
	if seekHouse.Level2AdminDiv != nil {
		var level2AdminDiv adminDivGetByCode
		level2AdminDiv.Code = *seekHouse.Level2AdminDiv
		tmpRes.Level2AdminDiv, _, _ = level2AdminDiv.Get()
	}
	if seekHouse.Level3AdminDiv != nil {
		var level3AdminDiv adminDivGetByCode
		level3AdminDiv.Code = *seekHouse.Level3AdminDiv
		tmpRes.Level3AdminDiv, _, _ = level3AdminDiv.Get()
	}
	if seekHouse.Level4AdminDiv != nil {
		var level4AdminDiv adminDivGetByCode
		level4AdminDiv.Code = *seekHouse.Level4AdminDiv
		tmpRes.Level4AdminDiv, _, _ = level4AdminDiv.Get()
	}

	//小区
	tmpRes.Community = seekHouse.Community

	//建筑面积
	if seekHouse.Area != nil {
		tmpRes.Area = *seekHouse.Area
	}

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
	if seekHouse.Name != nil {
		tmpRes.Name = *seekHouse.Name
	}
	if seekHouse.Gender != nil {
		var gender dictionaryDetailGet
		gender.Id = *seekHouse.Gender
		tmpRes.Gender, _, _ = gender.Get()
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
	if s.Level1AdminDiv > 0 {
		seekHouse.Level1AdminDiv = &s.Level1AdminDiv
	}
	if s.Level2AdminDiv > 0 {
		seekHouse.Level2AdminDiv = &s.Level2AdminDiv
	}
	if s.Level3AdminDiv > 0 {
		seekHouse.Level3AdminDiv = &s.Level3AdminDiv
	}
	if s.Level4AdminDiv > 0 {
		seekHouse.Level4AdminDiv = &s.Level4AdminDiv
	}

	//小区
	seekHouse.Community = s.Community

	//建筑面积
	if s.Area > 0 {
		seekHouse.Area = &s.Area
	}

	//姓名
	if s.Name != "" {
		seekHouse.Name = &s.Name
	}

	//性别
	if s.Gender > 0 {
		seekHouse.Gender = &s.Gender
	}

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
	if s.Level1AdminDiv > 0 {
		seekHouse["level_1_admin_div"] = s.Level1AdminDiv
	}
	if s.Level2AdminDiv > 0 {
		seekHouse["level_2_admin_div"] = s.Level2AdminDiv
	}
	if s.Level3AdminDiv > 0 {
		seekHouse["level_3_admin_div"] = s.Level3AdminDiv
	}
	if s.Level4AdminDiv > 0 {
		seekHouse["level_4_admin_div"] = s.Level4AdminDiv
	}

	//小区
	if s.Community != "" {
		seekHouse["community"] = s.Community
	}

	//建筑面积
	if s.Area > 0 {
		seekHouse["area"] = s.Area
	}

	//姓名
	if s.Name != "" {
		seekHouse["name"] = s.Name
	}

	//性别
	if s.Gender > 0 {
		seekHouse["gender"] = s.Gender
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

	if s.Creator > 0 {
		db = db.Where("creator = ?", s.Creator)
	}
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
	if s.Community != "" {
		db = db.Where("community LIKE ?", "%"+s.Community+"%")
	}
	if s.Level2AdminDiv > 0 {
		db = db.Where("level_2_admin_div = ?", s.Level2AdminDiv)
	}
	if s.Level3AdminDiv > 0 {
		db = db.Where("level_3_admin_div = ?", s.Level3AdminDiv)
	}
	if s.Level4AdminDiv > 0 {
		db = db.Where("level_4_admin_div = ?", s.Level4AdminDiv)
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

		// if seekHouse.MobilePhone != nil {
		// 	moblelePhone := *seekHouse.MobilePhone
		// 	result.MobilePhone = moblelePhone[:(len(*seekHouse.MobilePhone)-2)] + "**"
		// }

		// if seekHouse.WeChatId != nil {
		// 	wechatId := *seekHouse.WeChatId
		// 	result.WeChatId = wechatId[:(len(*seekHouse.WeChatId)-2)] + "**"
		// }

		//获取文件下载地址
		var download imageGetList
		download.businessType = seekHouse.TableName()
		download.businessId = seekHouse.Id
		result.Files, _, _, _ = download.GetList()

		//获取行政区划
		if seekHouse.Level1AdminDiv != nil {
			var level1AdminDiv adminDivGetByCode
			level1AdminDiv.Code = *seekHouse.Level1AdminDiv
			result.Level1AdminDiv, _, _ = level1AdminDiv.Get()
		}
		if seekHouse.Level2AdminDiv != nil {
			var level2AdminDiv adminDivGetByCode
			level2AdminDiv.Code = *seekHouse.Level2AdminDiv
			result.Level2AdminDiv, _, _ = level2AdminDiv.Get()
		}
		if seekHouse.Level3AdminDiv != nil {
			var level3AdminDiv adminDivGetByCode
			level3AdminDiv.Code = *seekHouse.Level3AdminDiv
			result.Level3AdminDiv, _, _ = level3AdminDiv.Get()
		}
		if seekHouse.Level4AdminDiv != nil {
			var level4AdminDiv adminDivGetByCode
			level4AdminDiv.Code = *seekHouse.Level4AdminDiv
			result.Level4AdminDiv, _, _ = level4AdminDiv.Get()
		}

		//小区
		result.Community = seekHouse.Community

		//建筑面积
		if seekHouse.Area != nil {
			result.Area = *seekHouse.Area
		}

		results = append(results, result)
	}

	//分页信息
	var tmpPaging response.Paging
	tmpPaging.Page = page
	tmpPaging.PageSize = pageSize
	tmpPaging.TotalRecords = int(count)
	tmpPaging.TotalPages = util.GetNumberOfPages(int(count), pageSize)

	return results, &tmpPaging, util.Success, nil
}
