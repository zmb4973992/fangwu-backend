package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type CommentCreate struct {
	Creator      int64  `json:"-"`
	ParentId     int64  `json:"parent_id,omitempty"`
	Content      string `json:"content" binding:"required"`
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
	OnlyToPoster *bool  `json:"only_to_poster" binding:"required"`
}

type CommentUpdate struct {
	LastModifier int64  `json:"-"`
	Id           int64  `json:"id" binding:"required"`
	Content      string `json:"content,omitempty"`
	OnlyToPoster *bool  `json:"only_to_poster,omitempty"`
}

type CommentDelete struct {
	Id      int64 `json:"id" binding:"required"`
	Deleter int64 `json:"-"`
}

type CommentGetList struct {
	list
	BusinessType string `json:"business_type" binding:"required"`
	BusinessId   int64  `json:"business_id" binding:"required"`
}

type CommentResult struct {
	Creator      int64  `json:"creator,omitempty"`
	LastModifier int64  `json:"last_modifier,omitempty"`
	Id           int64  `json:"id,omitempty"`
	ParentId     int64  `json:"parent_id,omitempty"`
	Content      string `json:"content,omitempty"`
	BusinessType string `json:"business_type,omitempty"`
	BusinessId   int64  `json:"business_id,omitempty"`
	OnlyToPoster bool   `json:"only_to_poster,omitempty"`
}

func (c *CommentCreate) Create() (result *CommentResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var comment model.Comment
	comment.Creator = &c.Creator
	comment.LastModifier = &c.Creator

	if c.ParentId > 0 {
		comment.ParentId = &c.ParentId
	}

	comment.IsHidden = false
	comment.Content = c.Content
	comment.BusinessType = c.BusinessType
	comment.BusinessId = c.BusinessId
	comment.OnlyToPoster = *c.OnlyToPoster

	//创建评论
	err := tx.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateComment, util.GetErrDetail(err)
	}

	//获取“消息类型-评论”的字典详情id
	var param dictionaryDetailGetByValue
	param.DictionaryTypeValue = "消息类型"
	param.DictionaryDetailValue = "评论"
	dictionaryDetail, resCode, errDetail := param.GetByValue()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//获取receiver
	var receiver int64
	switch c.BusinessType {
	case "for_rent":
		tx.Model(&model.ForRent{}).
			Where("id = ?", c.BusinessId).
			Select("creator").
			Limit(1).
			Find(&receiver)
	case "seek_house":
		tx.Model(&model.SeekHouse{}).
			Where("id =?", c.BusinessId).
			Select("creator").
			Limit(1).
			Find(&receiver)
	}

	//创建消息
	var notification NotificationCreate
	notification.Type = dictionaryDetail.Id
	notification.BusinessType = c.BusinessType
	notification.BusinessId = c.BusinessId
	notification.Sender = c.Creator
	notification.Receiver = receiver
	notification.Content = c.Content
	_, resCode, errDetail = notification.Create()
	if resCode != util.Success {
		tx.Rollback()
		return nil, resCode, errDetail
	}

	//提交事务
	tx.Commit()

	var tmpRes CommentResult
	tmpRes.Id = comment.Id

	return &tmpRes, util.Success, nil
}

func (c *CommentUpdate) Update() (result *CommentResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	house := make(map[string]any)

	if c.LastModifier > 0 {
		house["last_modifier"] = c.LastModifier
	}

	if c.Content != "" {
		house["content"] = c.Content
	}

	if c.OnlyToPoster != nil {
		house["only_to_poster"] = c.OnlyToPoster
	}

	err := tx.Model(&model.Comment{}).
		Where("id =?", c.Id).
		Where("creator = ?", c.LastModifier).
		Updates(house).Error
	if err != nil {
		return nil, util.ErrorFailToUpdateComment, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()

	var tmpRes CommentResult
	tmpRes.Id = c.Id

	return &tmpRes, util.Success, nil
}

// 删除出租记录：将记录转存到归档表备查，删除相关图片、减小空间占用
func (c *CommentDelete) Delete() (resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//获取出租记录
	var comment model.Comment
	err := tx.Where("id = ?", c.Id).
		Where("creator = ?", c.Deleter).
		First(&comment).Error
	if err != nil {
		tx.Rollback()
		return util.Success, nil
	}

	//存入归档表
	var archivedComment model.ArchivedComment
	archivedComment.Archive.Delete(c.Id, "用户删除")
	archivedComment.Comment = comment
	err = tx.Create(&archivedComment).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteComment, util.GetErrDetail(err)
	}

	//删除评论记录
	err = tx.Delete(&comment).Error
	if err != nil {
		tx.Rollback()
		return util.ErrorFailToDeleteForRent, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()
	return util.Success, nil
}

func (c *CommentGetList) GetList() (results []CommentResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.Comment{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("business_type = ?", c.BusinessType).
		Where("business_id = ?", c.BusinessId)

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
		var tmp model.Comment
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
	var comments []model.Comment
	db.Find(&comments)

	//将结果转换为Result
	for _, comment := range comments {
		var result CommentResult
		if comment.Creator != nil {
			result.Creator = *comment.Creator
		}
		if comment.LastModifier != nil {
			result.LastModifier = *comment.LastModifier
		}
		result.Id = comment.Id

		if comment.ParentId != nil {
			result.ParentId = *comment.ParentId
		}

		result.Content = comment.Content
		result.OnlyToPoster = comment.OnlyToPoster

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
