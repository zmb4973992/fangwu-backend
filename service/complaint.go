package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type ComplaintGet struct {
	Id int64 `json:"id"`
}

type ComplaintCreate struct {
	Creator      int64   `json:"-"`
	BusinessType string  `json:"business_type" binding:"required"`
	BusinessId   int64   `json:"business_id" binding:"required"`
	Reason       int64   `json:"reason" binding:"required"`
	Description  string  `json:"description,omitempty"`
	FileIds      []int64 `json:"file_ids,omitempty"`
}

type ComplaintUpdate struct {
	LastModifier int64  `json:"-"`
	Id           int64  `json:"id" binding:"required"`
	Status       int64  `json:"status,omitempty"`
	Response     string `json:"response,omitempty"`
}

type ComplaintDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type ComplaintGetList struct {
	list
	UserId int64 `json:"-"`
}

type ComplaintResult struct {
	Creator      int64                   `json:"creator,omitempty"`
	LastModifier int64                   `json:"last_modifier,omitempty"`
	Id           int64                   `json:"id,omitempty"`
	Reason       *DictionaryDetailResult `json:"reason,omitempty"`
	Description  string                  `json:"description,omitempty"`
	Status       *DictionaryDetailResult `json:"status,omitempty"`
	Response     string                  `json:"response,omitempty"`
	Files        []ImageResult           `json:"files,omitempty"`
}

func (c *ComplaintGet) Get() (result *ComplaintResult, resCode int, errDetail *util.ErrDetail) {
	//获取投诉记录
	var complaint model.Complaint
	err := global.Db.
		Where("id = ?", c.Id).
		First(&complaint).Error
	if err != nil {
		return nil, util.ErrorFailToGetComplaint, util.GetErrDetail(err)
	}

	var tmpRes ComplaintResult

	// 填充创建者
	if complaint.Creator != nil {
		tmpRes.Creator = *complaint.Creator
	}

	//填充最后修改者
	if complaint.LastModifier != nil {
		tmpRes.LastModifier = *complaint.LastModifier
	}

	//填充id
	tmpRes.Id = complaint.Id

	//填充投诉原因
	var reason dictionaryDetailGet
	reason.Id = complaint.Reason
	tmpRes.Reason, _, _ = reason.Get()

	//填充描述
	if complaint.Description != nil {
		tmpRes.Description = *complaint.Description
	}

	//填充状态
	var status dictionaryDetailGet
	status.Id = complaint.Status
	tmpRes.Status, _, _ = status.Get()

	//填充回复
	if complaint.Response != nil {
		tmpRes.Response = *complaint.Response
	}

	//获取文件下载地址
	var download imageGetList
	download.businessType = complaint.TableName()
	download.businessId = complaint.Id
	tmpRes.Files, _, _, _ = download.GetList()

	return &tmpRes, util.Success, nil
}

func (c *ComplaintCreate) Create() (result *ComplaintResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var complaint model.Complaint
	complaint.Creator = &c.Creator
	complaint.LastModifier = &c.Creator
	complaint.BusinessType = c.BusinessType
	complaint.BusinessId = c.BusinessId
	complaint.Reporter = c.Creator
	complaint.Reason = c.Reason
	complaint.Description = &c.Description

	//获取待处理的字典详情id
	var param dictionaryDetailGetByValue
	param.DictionaryTypeValue = "投诉状态"
	param.DictionaryDetailValue = "未处理"
	dictionaryDetail, resCode, errDetail := param.GetByValue()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}
	complaint.Status = dictionaryDetail.Id

	err := tx.Create(&complaint).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateComplaint, util.GetErrDetail(err)
	}

	//批量确认上传
	var file FileBatchConfirm
	file.FileIds = c.FileIds
	file.UserId = c.Creator
	file.BusinessType = complaint.TableName()
	file.BusinessId = complaint.Id
	resCode, errDetail = file.BatchConfirm()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes ComplaintResult
	tmpRes.Id = complaint.Id

	return &tmpRes, util.Success, nil
}

func (c *ComplaintUpdate) Update() (result *ComplaintResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	complaint := make(map[string]any)

	if c.LastModifier > 0 {
		complaint["last_modifier"] = c.LastModifier
	}

	complaint["status"] = c.Status
	complaint["response"] = c.Response

	err := tx.Model(&model.Complaint{}).
		Where("id = ?", c.Id).
		Where("creator = ?", c.LastModifier).
		Updates(complaint).Error
	if err != nil {
		return nil, util.ErrorFailToUpdateComment, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()

	var tmpRes ComplaintResult
	tmpRes.Id = c.Id

	return &tmpRes, util.Success, nil
}

func (c *ComplaintDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//获取投诉记录
	var complaint model.Complaint
	err := tx.Where("id = ?", c.Id).
		Where("creator = ?", c.Deleter).
		First(&complaint).Error
	if err != nil {
		tx.Rollback()
		return util.Success, nil
	}

	//存入归档表
	var archivedComplaint model.ArchivedComplaint
	archivedComplaint.Archive.Delete(c.Id, "用户删除")
	archivedComplaint.Complaint = complaint
	err = tx.Create(&archivedComplaint).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteComplaint, util.GetErrDetail(err)
	}

	//删除投诉记录
	err = tx.Delete(&complaint).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteComplaint, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()
	return util.Success, nil
}

func (c *ComplaintGetList) GetList() (results []ComplaintResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.Complaint{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("creator = ?", c.UserId)

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if c.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if c.Desc {
			db = db.Order("id desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.Complaint
		ok := util.FieldIsInModel(db, tmp.TableName(), c.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if c.Desc {
			db = db.Order(c.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(c.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if c.PageSize > 0 && c.PageSize <= maxPageSize {
		pageSize = c.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if c.Page > 0 {
		page = c.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var complaints []model.Complaint
	db.Find(&complaints)

	//将结果转换为Result
	for _, complaint := range complaints {
		var result ComplaintResult
		if complaint.Creator != nil {
			result.Creator = *complaint.Creator
		}
		if complaint.LastModifier != nil {
			result.LastModifier = *complaint.LastModifier
		}
		result.Id = complaint.Id

		//获取投诉原因
		var reason dictionaryDetailGet
		reason.Id = complaint.Reason
		result.Reason, _, _ = reason.Get()

		//获取投诉详情
		if complaint.Description != nil {
			result.Description = *complaint.Description
		}

		//获取投诉状态
		var status dictionaryDetailGet
		status.Id = complaint.Status
		result.Status, _, _ = status.Get()

		//获取投诉回复
		if complaint.Response != nil {
			result.Response = *complaint.Response
		}

		//获取文件下载地址
		var download imageGetList
		download.businessType = complaint.TableName()
		download.businessId = complaint.Id
		result.Files, _, _, _ = download.GetList()

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
