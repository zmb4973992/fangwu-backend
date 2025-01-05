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
	Area              int     `json:"area,omitempty"`                        //面积
	Bedroom           int     `json:"bedroom,omitempty"`                     //卧室数量
	LivingRoom        int     `json:"living_room,omitempty"`                 //客厅数量
	Bathroom          int     `json:"bathroom,omitempty"`                    //卫生间数量
	Kitchen           int     `json:"kitchen,omitempty"`                     //厨房数量
	Floor             int     `json:"floor,omitempty"`                       //楼层
	TotalFloor        int     `json:"total_floor,omitempty"`                 //总楼层
	Orientation       int64   `json:"orientation,omitempty"`                 //朝向
	Tenant            int     `json:"tenant,omitempty"`                      //合租户数
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
	Area              int     `json:"area,omitempty"`               //面积
	Bedroom           int     `json:"bedroom,omitempty"`            //卧室数量
	LivingRoom        int     `json:"living_room,omitempty"`        //客厅数量
	Bathroom          int     `json:"bathroom,omitempty"`           //卫生间数量
	Kitchen           int     `json:"kitchen,omitempty"`            //厨房数量
	Floor             int     `json:"floor,omitempty"`              //楼层
	TotalFloor        int     `json:"total_floor,omitempty"`        //总楼层
	Orientation       int64   `json:"orientation,omitempty"`        //朝向
	Tenant            int     `json:"tenant,omitempty"`             //合租户数
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
	Level2AdminDiv     int     `json:"level_2_admin_div,omitempty"`    //二级行政区划
	Level3AdminDiv     int     `json:"level_3_admin_div,omitempty"`    //三级行政区划
	Level4AdminDiv     int     `json:"level_4_admin_div,omitempty"`    //四级行政区划
}

type ForRentResult struct {
	Creator           int64                   `json:"creator,omitempty"`
	LastModifier      int64                   `json:"last_modifier,omitempty"`
	IsDeleted         bool                    `json:"is_deleted,omitempty"`
	Id                int64                   `json:"id,omitempty"`
	RentType          *DictionaryDetailResult `json:"rent_type,omitempty"`
	Price             float64                 `json:"price,omitempty"`              //租赁类型，如整租、合租等
	Description       string                  `json:"description,omitempty"`        //描述
	GenderRestriction *DictionaryDetailResult `json:"gender_restriction,omitempty"` //性别限制，男、女、男女不限等
	MobilePhone       string                  `json:"mobile_phone,omitempty"`       //手机号
	WeChatId          string                  `json:"wechat_id,omitempty"`          //微信id
	Files             []ImageResult           `json:"files,omitempty"`
	Level1AdminDiv    *AdminDivResult         `json:"level_1_admin_div,omitempty"`
	Level2AdminDiv    *AdminDivResult         `json:"level_2_admin_div,omitempty"`
	Level3AdminDiv    *AdminDivResult         `json:"level_3_admin_div,omitempty"`
	Level4AdminDiv    *AdminDivResult         `json:"level_4_admin_div,omitempty"`
	Community         string                  `json:"community,omitempty"`   //小区
	Area              int                     `json:"area,omitempty"`        //面积
	Bedroom           int                     `json:"bedroom,omitempty"`     //卧室数量
	LivingRoom        int                     `json:"living_room,omitempty"` //客厅数量
	Bathroom          int                     `json:"bathroom,omitempty"`    //卫生间数量
	Kitchen           int                     `json:"kitchen,omitempty"`     //厨房数量
	Floor             int                     `json:"floor,omitempty"`       //楼层
	TotalFloor        int                     `json:"total_floor,omitempty"` //总楼层
	Orientation       *DictionaryDetailResult `json:"orientation,omitempty"` //朝向
	Tenant            int                     `json:"tenant,omitempty"`      //合租户数

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
	if forRent.Creator != nil {
		tmpRes.Creator = *forRent.Creator
	}
	if forRent.LastModifier != nil {
		tmpRes.LastModifier = *forRent.LastModifier
	}
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

	//获取文件信息
	var download imageGetList
	download.businessType = forRent.TableName()
	download.businessId = forRent.Id
	tmpRes.Files, _, _, _ = download.GetList()

	//获取行政区划
	if forRent.Level1AdminDiv != nil {
		var level1AdminDiv adminDivGet
		level1AdminDiv.Code = *forRent.Level1AdminDiv
		tmpRes.Level1AdminDiv, _, _ = level1AdminDiv.Get()
	}
	if forRent.Level2AdminDiv != nil {
		var level2AdminDiv adminDivGet
		level2AdminDiv.Code = *forRent.Level2AdminDiv
		tmpRes.Level2AdminDiv, _, _ = level2AdminDiv.Get()
	}
	if forRent.Level3AdminDiv != nil {
		var level3AdminDiv adminDivGet
		level3AdminDiv.Code = *forRent.Level3AdminDiv
		tmpRes.Level3AdminDiv, _, _ = level3AdminDiv.Get()
	}
	if forRent.Level4AdminDiv != nil {
		var level4AdminDiv adminDivGet
		level4AdminDiv.Code = *forRent.Level4AdminDiv
		tmpRes.Level4AdminDiv, _, _ = level4AdminDiv.Get()
	}

	//小区
	tmpRes.Community = forRent.Community

	//建筑面积
	if forRent.Area != nil {
		tmpRes.Area = *forRent.Area
	}

	//户型
	if forRent.Bedroom != nil {
		tmpRes.Bedroom = *forRent.Bedroom
	}
	if forRent.LivingRoom != nil {
		tmpRes.LivingRoom = *forRent.LivingRoom
	}
	if forRent.Bathroom != nil {
		tmpRes.Bathroom = *forRent.Bathroom
	}
	if forRent.Kitchen != nil {
		tmpRes.Kitchen = *forRent.Kitchen
	}

	//楼层
	if forRent.Floor != nil {
		tmpRes.Floor = *forRent.Floor
	}
	if forRent.TotalFloor != nil {
		tmpRes.TotalFloor = *forRent.TotalFloor
	}

	//朝向
	if forRent.Orientation != nil {
		var orientation dictionaryDetailGet
		orientation.Id = *forRent.Orientation
		tmpRes.Orientation, _, _ = orientation.Get()
	}

	//合租户数
	if forRent.Tenant != nil {
		tmpRes.Tenant = *forRent.Tenant
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
	if forRent.WechatId != nil {
		tmpRes.WeChatId = *forRent.WechatId
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
		forRent.WechatId = model.StringToPointer(f.WeChatId)
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

	//建筑面积
	if f.Area > 0 {
		forRent.Area = &f.Area
	}

	//户型
	if f.Bedroom > 0 {
		forRent.Bedroom = &f.Bedroom
	}
	if f.LivingRoom > 0 {
		forRent.LivingRoom = &f.LivingRoom
	}
	if f.Bathroom > 0 {
		forRent.Bathroom = &f.Bathroom
	}
	if f.Kitchen > 0 {
		forRent.Kitchen = &f.Kitchen
	}

	//楼层
	if f.Floor > 0 {
		forRent.Floor = &f.Floor
	}
	if f.TotalFloor > 0 {
		forRent.TotalFloor = &f.TotalFloor
	}

	//朝向
	if f.Orientation > 0 {
		forRent.Orientation = &f.Orientation
	}

	//合租户数
	if f.Tenant > 0 {
		forRent.Tenant = &f.Tenant
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

	//建筑面积
	if f.Area > 0 {
		forRent["area"] = f.Area
	}

	//户型
	if f.Bedroom > 0 {
		forRent["bedroom"] = f.Bedroom
	}
	if f.LivingRoom > 0 {
		forRent["living_room"] = f.LivingRoom
	}
	if f.Bathroom > 0 {
		forRent["bathroom"] = f.Bathroom
	}
	if f.Kitchen > 0 {
		forRent["kitchen"] = f.Kitchen
	}

	//楼层
	if f.Floor > 0 {
		forRent["floor"] = f.Floor
	}
	if f.TotalFloor > 0 {
		forRent["total_floor"] = f.TotalFloor
	}

	//朝向
	if f.Orientation > 0 {
		forRent["orientation"] = f.Orientation
	}

	//合租户数
	if f.Tenant > 0 {
		forRent["tenant"] = f.Tenant
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
	if f.Level2AdminDiv > 0 {
		db = db.Where("level_2_admin_div = ?", f.Level2AdminDiv)
	}
	if f.Level3AdminDiv > 0 {
		db = db.Where("level_3_admin_div = ?", f.Level3AdminDiv)
	}
	if f.Level4AdminDiv > 0 {
		db = db.Where("level_4_admin_div = ?", f.Level4AdminDiv)
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

		//获取文件下载地址
		var download imageGetList
		download.businessType = forRent.TableName()
		download.businessId = forRent.Id
		result.Files, _, _, _ = download.GetList()

		//获取行政区划
		if forRent.Level1AdminDiv != nil {
			var adminDiv adminDivGet
			adminDiv.Code = *forRent.Level1AdminDiv
			result.Level1AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level2AdminDiv != nil {
			var adminDiv adminDivGet
			adminDiv.Code = *forRent.Level2AdminDiv
			result.Level2AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level3AdminDiv != nil {
			var adminDiv adminDivGet
			adminDiv.Code = *forRent.Level3AdminDiv
			result.Level3AdminDiv, _, _ = adminDiv.Get()
		}
		if forRent.Level4AdminDiv != nil {
			var adminDiv adminDivGet
			adminDiv.Code = *forRent.Level4AdminDiv
			result.Level4AdminDiv, _, _ = adminDiv.Get()
		}
		//小区
		result.Community = forRent.Community

		//建筑面积
		if forRent.Area != nil {
			result.Area = *forRent.Area
		}

		//户型
		if forRent.Bedroom != nil {
			result.Bedroom = *forRent.Bedroom
		}
		if forRent.LivingRoom != nil {
			result.LivingRoom = *forRent.LivingRoom
		}
		if forRent.Bathroom != nil {
			result.Bathroom = *forRent.Bathroom
		}
		if forRent.Kitchen != nil {
			result.Kitchen = *forRent.Kitchen
		}

		//楼层
		if forRent.Floor != nil {
			result.Floor = *forRent.Floor
		}
		if forRent.TotalFloor != nil {
			result.TotalFloor = *forRent.TotalFloor
		}

		//朝向
		if forRent.Orientation != nil {
			var orientation dictionaryDetailGet
			orientation.Id = *forRent.Orientation
			result.Orientation, _, _ = orientation.Get()
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
