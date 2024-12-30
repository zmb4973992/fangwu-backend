package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
	"time"
)

type ForRentGet struct {
	Id int64 `json:"-"`
}

type ForRentGetContact struct {
	Id int64 `json:"-"`
}

type ForRentCreate struct {
	Creator           int64   `json:"-"`                                     //用户id
	Price             float64 `json:"price" binding:"required"`              //价格
	RentType          int64   `json:"rent_type" binding:"required"`          //租赁类型，如整租、合租等
	Description       string  `json:"description" binding:"required"`        //描述
	GenderRestriction int64   `json:"gender_restriction" binding:"required"` //男、女、男女不限等
	MobilePhone       string  `json:"mobile_phone,omitempty"`                //手机号
	WeChatId          string  `json:"wechat_id,omitempty"`                   //微信id
	FileIds           []int64 `json:"file_ids,omitempty"`                    //文件id
	Level1AdminDiv    int     `json:"level_1_admin_div,omitempty"`           //一级行政区划（省/自治区/直辖市）
	Level2AdminDiv    int     `json:"level_2_admin_div,omitempty"`           //二级行政区划（市/区/县）
	Level3AdminDiv    int     `json:"level_3_admin_div,omitempty"`           //三级行政区划（乡/镇）
	Level4AdminDiv    int     `json:"level_4_admin_div,omitempty"`           //四级行政区划（村/社区）
	Community         string  `json:"community" binding:"required"`          //小区
	HouseType         int64   `json:"house_type,omitempty"`                  //户型，如一室、二室等
	BuildingArea      int     `json:"building_area,omitempty"`               //建筑面积
}

type ForRentUpdate struct {
	LastModifier      int64   `json:"-"`
	Id                int64   `json:"id" binding:"required"`
	Price             float64 `json:"price,omitempty"`              //价格
	RentType          int64   `json:"rent_type,omitempty"`          //租赁类型，如整租、合租等
	Description       string  `json:"description,omitempty"`        //描述
	GenderRestriction int64   `json:"gender_restriction,omitempty"` //男、女、男女不限等
	MobilePhone       string  `json:"mobile_phone,omitempty"`       //手机号
	WeChatId          string  `json:"wechat_id,omitempty"`          //微信id
	FileIds           []int64 `json:"file_ids,omitempty"`           //文件id
	Level1AdminDiv    int     `json:"level_1_admin_div,omitempty"`  //一级行政区划（省/自治区/直辖市）
	Level2AdminDiv    int     `json:"level_2_admin_div,omitempty"`  //二级行政区划（市/区/县）
	Level3AdminDiv    int     `json:"level_3_admin_div,omitempty"`  //三级行政区划（乡/镇）
	Level4AdminDiv    int     `json:"level_4_admin_div,omitempty"`  //四级行政区划（村/社区）
	Community         string  `json:"community,omitempty"`          //小区
	HouseType         int64   `json:"house_type,omitempty"`         //户型，如一室、二室等
	BuildingArea      int     `json:"building_area,omitempty"`      //建筑面积
}

type ForRentDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type ForRentGetList struct {
	list
	RentType           int64   `json:"rent_type,omitempty"`            //租赁类型，如整租、合租等
	MaxPrice           float64 `json:"max_price,omitempty"`            //最高价格
	MinPrice           float64 `json:"min_price,omitempty"`            //最低价格
	GenderRestriction  int64   `json:"gender_restriction,omitempty"`   //性别限制，男、女、男女不限等
	MobilePhoneInclude string  `json:"mobile_phone_include,omitempty"` //手机号包含
	Ids                []int64 `json:"-"`                              //id列表
	Keyword            string  `json:"keyword,omitempty"`              //关键字
	Community          string  `json:"community,omitempty"`            //小区
	HouseType          int64   `json:"house_type,omitempty"`           //户型，如一室、二室等
}

type ForRentResult struct {
	Creator           int64                         `json:"creator,omitempty"`
	LastModifier      int64                         `json:"last_modifier,omitempty"`
	IsDeleted         bool                          `json:"is_deleted,omitempty"`
	Id                int64                         `json:"id,omitempty"`
	RentType          *DictionaryDetailResult       `json:"rent_type,omitempty"`
	Price             float64                       `json:"price,omitempty"`              //租赁类型，如整租、合租等
	Description       string                        `json:"description,omitempty"`        //描述
	GenderRestriction *DictionaryDetailResult       `json:"gender_restriction,omitempty"` //性别限制，男、女、男女不限等
	MobilePhone       string                        `json:"mobile_phone,omitempty"`       //手机号
	WeChatId          string                        `json:"wechat_id,omitempty"`          //微信id
	Files             []ImageResult                 `json:"files,omitempty"`
	Level1AdminDiv    *AdministrativeDivisionResult `json:"level_1_admin_div,omitempty"`
	Level2AdminDiv    *AdministrativeDivisionResult `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    *AdministrativeDivisionResult `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    *AdministrativeDivisionResult `json:"level_4_admin_div,omitempty"`
	Community         string                        `json:"community,omitempty"`     //小区
	HouseType         *DictionaryDetailResult       `json:"house_type,omitempty"`    //户型，如一室、二室等
	BuildingArea      int                           `json:"building_area,omitempty"` //建筑面积
}

func (f *ForRentGet) Get() (result *ForRentResult, resCode int, errDetail *util.ErrDetail) {
	//获取出租记录
	var forRent model.ForRent
	err := global.Db.
		Where("id = ?", f.Id).
		First(&forRent).Error
	if err != nil {
		return nil, util.ErrorFailToGetForRent, util.GetErrDetail(err)
	}

	var tmpRes ForRentResult

	// 填充结果
	tmpRes.IsDeleted = forRent.IsDeleted
	tmpRes.Id = forRent.Id
	tmpRes.Price = forRent.Price

	//填充租赁类型
	var rentType dictionaryDetailGet
	rentType.Id = forRent.RentType
	tmpRes.RentType, _, _ = rentType.Get()

	tmpRes.Description = forRent.Description

	// 填充性别限制
	var genderRestriction dictionaryDetailGet
	genderRestriction.Id = forRent.GenderRestriction
	tmpRes.GenderRestriction, _, _ = genderRestriction.Get()

	//获取文件下载地址
	var download imageGetList
	download.businessType = forRent.TableName()
	download.businessId = forRent.Id
	tmpRes.Files, _, _, _ = download.GetList()

	//获取行政区划
	if forRent.Level1AdminDiv != nil {
		var level1AdminDiv administrativeDivisionGet
		level1AdminDiv.Code = *forRent.Level1AdminDiv
		tmpRes.Level1AdminDiv, _, _ = level1AdminDiv.Get()
	}
	if forRent.Level2AdminDiv != nil {
		var level2AdminDiv administrativeDivisionGet
		level2AdminDiv.Code = *forRent.Level2AdminDiv
		tmpRes.Level2AdminDiv, _, _ = level2AdminDiv.Get()
	}
	if forRent.Level3AdminDiv != nil {
		var level3AdminDiv administrativeDivisionGet
		level3AdminDiv.Code = *forRent.Level3AdminDiv
		tmpRes.Level3AdminDiv, _, _ = level3AdminDiv.Get()
	}
	if forRent.Level4AdminDiv != nil {
		var level4AdminDiv administrativeDivisionGet
		level4AdminDiv.Code = *forRent.Level4AdminDiv
		tmpRes.Level4AdminDiv, _, _ = level4AdminDiv.Get()
	}

	//小区
	tmpRes.Community = forRent.Community

	//户型
	if forRent.HouseType != nil {
		var houseType dictionaryDetailGet
		houseType.Id = *forRent.HouseType
		tmpRes.HouseType, _, _ = houseType.Get()
	}

	//建筑面积
	if forRent.BuildingArea != nil {
		tmpRes.BuildingArea = *forRent.BuildingArea
	}

	return &tmpRes, util.Success, nil
}

func (f *ForRentGetContact) GetContact() (result *ForRentResult, resCode int, errDetail *util.ErrDetail) {
	//获取出租记录
	var forRent model.ForRent
	err := global.Db.
		Where("id = ?", f.Id).
		First(&forRent).Error
	if err != nil {
		return nil, util.ErrorFailToGetForRent, util.GetErrDetail(err)
	}

	var tmpRes ForRentResult

	//获取联系方式
	if forRent.MobilePhone != nil {
		tmpRes.MobilePhone = *forRent.MobilePhone
	}
	if forRent.WeChatId != nil {
		tmpRes.WeChatId = *forRent.WeChatId
	}

	return &tmpRes, util.Success, nil
}

func (f *ForRentCreate) Create() (result *ForRentResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var forRent model.ForRent
	forRent.Creator = &f.Creator
	forRent.LastModifier = &f.Creator
	forRent.IsDeleted = false
	forRent.Price = f.Price
	forRent.RentType = f.RentType
	forRent.Description = f.Description
	forRent.GenderRestriction = f.GenderRestriction
	if f.MobilePhone != "" {
		forRent.MobilePhone = model.StringToPointer(f.MobilePhone)
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "mobile_phone"
		contactInfoBlacklist.Value = f.MobilePhone
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

	if f.WeChatId != "" {
		forRent.WeChatId = model.StringToPointer(f.WeChatId)
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "wechat_id"
		contactInfoBlacklist.Value = f.WeChatId
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorWechatIdIsInBlacklist, nil
		}
	}

	//行政区划
	if f.Level1AdminDiv > 0 {
		forRent.Level1AdminDiv = &f.Level1AdminDiv
	}
	if f.Level2AdminDiv > 0 {
		forRent.Level2AdminDiv = &f.Level2AdminDiv
	}
	if f.Level3AdminDiv > 0 {
		forRent.Level3AdminDiv = &f.Level3AdminDiv
	}
	if f.Level4AdminDiv > 0 {
		forRent.Level4AdminDiv = &f.Level4AdminDiv
	}

	//小区
	forRent.Community = f.Community

	//户型
	if f.HouseType > 0 {
		forRent.HouseType = &f.HouseType
	}

	//建筑面积
	if f.BuildingArea > 0 {
		forRent.BuildingArea = &f.BuildingArea
	}

	err := tx.Create(&forRent).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateForRent, util.GetErrDetail(err)
	}

	//批量确认上传
	var file FileBatchConfirm
	file.FileIds = f.FileIds
	file.UserId = f.Creator
	file.BusinessType = forRent.TableName()
	file.BusinessId = forRent.Id
	resCode, errDetail = file.BatchConfirm()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes ForRentResult
	tmpRes.Id = forRent.Id

	return &tmpRes, util.Success, nil
}

func (f *ForRentUpdate) Update() (result *ForRentResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	forRent := make(map[string]any)

	if f.LastModifier > 0 {
		forRent["last_modifier"] = f.LastModifier
	}

	if f.Price > 0 {
		forRent["price"] = f.Price
	}

	if f.RentType > 0 {
		forRent["rent_type"] = f.RentType
	}

	if f.Description != "" {
		forRent["description"] = f.Description
	}

	if f.GenderRestriction > 0 {
		forRent["gender_restriction"] = f.GenderRestriction
	}

	if f.MobilePhone != "" {
		forRent["mobile_phone"] = f.MobilePhone
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "mobile_phone"
		contactInfoBlacklist.Value = f.MobilePhone
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

	if f.WeChatId != "" {
		forRent["wechat_id"] = f.WeChatId
		var contactInfoBlacklist ContactBlackListVerify
		contactInfoBlacklist.Type = "wechat_id"
		contactInfoBlacklist.Value = f.WeChatId
		isBlocked, resCode, errerrDetail := contactInfoBlacklist.Verify()
		if resCode != util.Success {
			tx.Rollback()
			return nil, resCode, errerrDetail
		}
		if isBlocked {
			tx.Rollback()
			return nil, util.ErrorWechatIdIsInBlacklist, nil
		}
	}

	//行政区划
	if f.Level1AdminDiv > 0 {
		forRent["level_1_admin_div"] = f.Level1AdminDiv
	}
	if f.Level2AdminDiv > 0 {
		forRent["level_2_admin_div"] = f.Level2AdminDiv
	}
	if f.Level3AdminDiv > 0 {
		forRent["level_3_admin_div"] = f.Level3AdminDiv
	}
	if f.Level4AdminDiv > 0 {
		forRent["level_4_admin_div"] = f.Level4AdminDiv
	}

	//小区
	if f.Community != "" {
		forRent["community"] = f.Community
	}

	//户型
	if f.HouseType > 0 {
		forRent["house_type"] = f.HouseType
	}

	//建筑面积
	if f.BuildingArea > 0 {
		forRent["building_area"] = f.BuildingArea
	}

	err := tx.Model(&model.ForRent{}).
		Where("id =?", f.Id).
		Where("creator = ?", f.LastModifier).
		Where("is_deleted = ?", false).
		Updates(forRent).Error
	if err != nil {
		return nil, util.ErrorFailToUpdateForRent, util.GetErrDetail(err)
	}

	//批量确认上传
	var uploadBatchConfirm FileBatchConfirm
	uploadBatchConfirm.FileIds = f.FileIds
	uploadBatchConfirm.UserId = f.LastModifier
	uploadBatchConfirm.BusinessType = "for_rent"
	uploadBatchConfirm.BusinessId = f.Id
	resCode, errDetail = uploadBatchConfirm.BatchConfirm()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes ForRentResult
	tmpRes.Id = f.Id

	return &tmpRes, util.Success, nil
}

// 删除出租记录：将记录转存到归档表备查，删除相关图片、减小空间占用
func (f *ForRentDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//获取出租记录
	var forRent model.ForRent
	err := tx.Where("id = ?", f.Id).
		Where("creator = ?", f.Deleter).
		Where("is_deleted = ?", false).
		First(&forRent).Error
	if err != nil {
		tx.Rollback()
		return util.Success, nil
	}

	//软删除出租记录
	err = tx.Model(&forRent).
		Updates(map[string]any{
			"is_deleted": true,
			"deleter":    f.Deleter,
			"deleted_at": time.Now(),
		}).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteForRent, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()
	return util.Success, nil
}

func (f *ForRentGetList) GetList() (results []ForRentResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.ForRent{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("is_deleted = ?", false)

	if f.RentType > 0 {
		db = db.Where("rent_type = ?", f.RentType)
	}
	if f.GenderRestriction > 0 {
		db = db.Where("gender_restriction = ?", f.GenderRestriction)
	}
	if f.MaxPrice > 0 {
		db = db.Where("price <= ?", f.MaxPrice)
	}
	if f.MinPrice > 0 {
		db = db.Where("price >= ?", f.MinPrice)
	}
	if f.MobilePhoneInclude != "" {
		db = db.Where("mobile_phone LIKE ?", "%"+f.MobilePhoneInclude+"%")
	}
	if len(f.Ids) > 0 {
		db = db.Where("id in ?", f.Ids)
	}
	if f.Keyword != "" {
		db = db.Where("description LIKE ?", "%"+f.Keyword+"%")
	}
	if f.Community != "" {
		db = db.Where("community LIKE ?", "%"+f.Community+"%")
	}
	if f.HouseType > 0 {
		db = db.Where("house_type = ?", f.HouseType)
	}

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if f.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if f.Desc {
			db = db.Order("id desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.ForRent
		ok := util.FieldIsInModel(db, tmp.TableName(), f.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if f.Desc {
			db = db.Order(f.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(f.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if f.PageSize > 0 && f.PageSize <= maxPageSize {
		pageSize = f.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if f.Page > 0 {
		page = f.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var forRents []model.ForRent
	db.Find(&forRents)

	//将结果转换为Result
	for _, forRent := range forRents {
		var result ForRentResult

		result.Id = forRent.Id

		//获取租赁类型
		var rentType dictionaryDetailGet
		rentType.Id = forRent.RentType
		result.RentType, _, _ = rentType.Get()

		result.Price = forRent.Price
		result.Description = forRent.Description

		//获取性别限制
		var genderRestriction dictionaryDetailGet
		genderRestriction.Id = forRent.GenderRestriction
		result.GenderRestriction, _, _ = genderRestriction.Get()

		// if forRent.MobilePhone != nil {
		// 	moblelePhone := *forRent.MobilePhone
		// 	result.MobilePhone = moblelePhone[:(len(*forRent.MobilePhone)-2)] + "**"
		// }

		// if forRent.WeChatId != nil {
		// 	wechatId := *forRent.WeChatId
		// 	result.WeChatId = wechatId[:(len(*forRent.WeChatId)-2)] + "**"
		// }

		//获取文件下载地址
		var download imageGetList
		download.businessType = forRent.TableName()
		download.businessId = forRent.Id
		result.Files, _, _, _ = download.GetList()

		//获取行政区划
		if forRent.Level1AdminDiv != nil {
			var adminDiv administrativeDivisionGet
			adminDiv.Code = *forRent.Level1AdminDiv
			result.Level1AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level2AdminDiv != nil {
			var adminDiv administrativeDivisionGet
			adminDiv.Code = *forRent.Level2AdminDiv
			result.Level2AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level3AdminDiv != nil {
			var adminDiv administrativeDivisionGet
			adminDiv.Code = *forRent.Level3AdminDiv
			result.Level3AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level4AdminDiv != nil {
			var adminDiv administrativeDivisionGet
			adminDiv.Code = *forRent.Level4AdminDiv
			result.Level4AdminDiv, _, _ = adminDiv.Get()
		}
		//小区
		result.Community = forRent.Community

		//户型
		if forRent.HouseType != nil {
			var houseType dictionaryDetailGet
			houseType.Id = *forRent.HouseType
			result.HouseType, _, _ = houseType.Get()
		}

		//建筑面积
		if forRent.BuildingArea != nil {
			result.BuildingArea = *forRent.BuildingArea
		}

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
