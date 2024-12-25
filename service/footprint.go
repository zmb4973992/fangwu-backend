package service

type FootprintGet struct {
	Id int64 `json:"-"`
}

type FootprintCreate struct {
	Creator      int64  `json:"-"`
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

type FootprintDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type FootprintGetList struct {
	list
	UserId int64 `json:"-"`
}

type FootprintResult struct {
	Id          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

// func (f *FootprintGet) Get() (result *FootprintResult, resCode int, errDetail *util.ErrDetail) {
// 	//获取足迹记录
// 	var footprint model.Footprint
// 	err := global.Db.
// 		Where("id = ?", f.Id).
// 		First(&footprint).Error
// 	if err != nil {
// 		return nil, util.ErrorFailToGetFootprint, util.GetErrDetail(err)
// 	}

// 	//如果足迹记录的业务类型为出租
// 	var forRent model.ForRent
// 	if footprint.BusinessType == forRent.TableName() {
// 		var forRent ForRentGet
// 		forRent.Id = footprint.BusinessId
// 		result, resCode, errDetail = forRent.Get()
// 		if resCode != util.Success {
// 			return nil, resCode, errDetail
// 		}
// 		return result, resCode, errDetail
// 	}

// 	//如果足迹记录的业务类型为求租
// 	var seekHouse model.SeekHouse
// 	if footprint.BusinessType == seekHouse.TableName() {
// 		var seekHouse SeekHouseGet
// 		seekHouse.Id = footprint.BusinessId
// 		result, resCode, errDetail = seekHouse.Get()
// 		if resCode != util.Success {
// 			return nil, resCode, errDetail
// 		}
// 		return result, resCode, errDetail
// 	}

// 	return nil, util.ErrorInvalidBusinessType, nil
// }

// func (f *FootprintCreate) Create() (result *FootprintResult, resCode int, errDetail *util.ErrDetail) {
// 	var tmpRes FootprintResult

// 	//检查是否已存在足迹记录
// 	var footprint model.Footprint
// 	global.Db.
// 		Where("creator =?", f.Creator).
// 		Where("business_type = ?", f.BusinessType).
// 		Where("business_id = ?", f.BusinessId).
// 		Limit(1).
// 		Find(&footprint)

// 	//如果已存在，则更新足迹记录
// 	if footprint.Id != 0 {
// 		err := global.Db.Model(&footprint).
// 			Where("creator =?", f.Creator).
// 			Where("business_type =?", f.BusinessType).
// 			Where("business_id =?", f.BusinessId).
// 			Updates(map[string]any{
// 				"last_modifier":    f.Creator,
// 				"last_modified_at": time.Now(),
// 			}).Error
// 		if err != nil {
// 			return nil, util.ErrorFailToUpdateFootprint, util.GetErrDetail(err)
// 		}
// 		tmpRes.Id = footprint.Id
// 		return &tmpRes, util.Success, nil
// 	}

// 	//创建足迹记录
// 	footprint.Creator = &f.Creator
// 	footprint.LastModifier = &f.Creator
// 	footprint.BusinessType = f.BusinessType
// 	footprint.BusinessId = f.BusinessId

// 	err := global.Db.Create(&footprint).Error
// 	if err != nil {
// 		return nil, util.ErrorFailToCreateFootprint, util.GetErrDetail(err)
// 	}

// 	tmpRes.Id = footprint.Id
// 	return &tmpRes, util.Success, nil
// }

// func (f *FootprintDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
// 	//删除足迹记录
// 	err := global.Db.
// 		Where("creator =?", f.Deleter).
// 		Where("id =?", f.Id).
// 		Delete(&model.Footprint{}).Error
// 	if err != nil {
// 		return util.ErrorFailToDeleteFootprint, util.GetErrDetail(err)
// 	}

// 	return util.Success, nil
// }

// func (f *FootprintGetList) GetList() (results any, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
// 	db := global.Db.Model(&model.Footprint{})
// 	// 顺序：where -> count -> Order -> limit -> offset -> data

// 	// where
// 	db = db.Where("creator = ?", f.UserId).
// 		Where("business_type = ?", f.BusinessType)

// 	//找到所有业务id
// 	var businessIds []int64
// 	db.Select("business_id").Find(&businessIds)

// 	//如果收藏记录的业务类型为出租
// 	var forRent model.ForRent
// 	if f.BusinessType == forRent.TableName() {
// 		var param ForRentGetList
// 		param.Page = f.Page
// 		param.PageSize = f.PageSize
// 		param.OrderBy = f.OrderBy
// 		param.Desc = f.Desc
// 		param.Ids = businessIds
// 		results, paging, resCode, errDetail = param.GetList()
// 		if resCode != util.Success {
// 			return nil, nil, resCode, errDetail
// 		}
// 		return results, paging, resCode, errDetail
// 	}

// 	//如果收藏记录的业务类型为求租
// 	var seekHouse model.SeekHouse
// 	if f.BusinessType == seekHouse.TableName() {
// 		var param SeekHouseGetList
// 		param.Page = f.Page
// 		param.PageSize = f.PageSize
// 		param.OrderBy = f.OrderBy
// 		param.Desc = f.Desc
// 		param.Ids = businessIds
// 		results, paging, resCode, errDetail = param.GetList()
// 		if resCode != util.Success {
// 			return nil, nil, resCode, errDetail
// 		}
// 		return results, paging, resCode, errDetail
// 	}

// 	return nil, nil, util.ErrorInvalidBusinessType, nil
// }
