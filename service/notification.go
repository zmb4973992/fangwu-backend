package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/response"
	"fangwu-backend/util"
)

type NotificationCreate struct {
	Type         int64  `json:"-"`
	BusinessType string `json:"-"`
	BusinessId   int64  `json:"-"`
	Sender       int64  `json:"-"`
	Receiver     int64  `json:"receiver" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

type NotificationUpdate struct {
	Id           int64 `json:"id" binding:"required"`
	LastModifier int64 `json:"-"`
	IsRead       *bool `json:"is_read"`
}

type NotificationGetList struct {
	list
	Receiver int64 `json:"-"`
}

type NotificationResult struct {
	Id           int64  `json:"id,omitempty"`
	BusinessType string `json:"business_type,omitempty"`
	BusinessId   int64  `json:"business_id,omitempty"`
	Receiver     int64  `json:"receiver,omitempty"`
	Content      string `json:"content,omitempty"`
}

func (n *NotificationCreate) Create() (result *NotificationResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	var Notification model.Notification
	Notification.Creator = &n.Sender
	Notification.LastModifier = &n.Sender

	Notification.Type = n.Type

	if n.BusinessType != "" {
		Notification.BusinessType = &n.BusinessType
	}

	if n.BusinessId > 0 {
		Notification.BusinessId = &n.BusinessId
	}

	Notification.Receiver = n.Receiver
	Notification.Content = n.Content
	Notification.IsRead = false

	err := tx.Create(&Notification).Error
	if err != nil {
		tx.Rollback()
		return nil, util.ErrorFailToCreateNotification, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()

	var tmpRes NotificationResult
	tmpRes.Id = Notification.Id

	return &tmpRes, util.Success, nil
}

func (n *NotificationUpdate) Update() (result *NotificationResult, resCode int, errDetail *util.ErrDetail) {
	// 开启事务
	tx := global.Db.Begin()

	//接收参数
	house := make(map[string]any)

	if n.LastModifier > 0 {
		house["last_modifier"] = n.LastModifier
	}

	if n.IsRead != nil {
		house["is_read"] = *n.IsRead
	}

	err := tx.Model(&model.Notification{}).
		Where("id = ?", n.Id).
		Where("receiver = ?", n.LastModifier).
		Updates(house).Error
	if err != nil {
		return nil, util.ErrorFailToUpdateNotification, util.GetErrDetail(err)
	}

	//提交事务
	tx.Commit()

	return nil, util.Success, nil
}

func (m *NotificationGetList) GetList() (results []NotificationResult, paging *response.Paging, resCode int, errDetail *util.ErrDetail) {
	db := global.Db.Model(&model.Notification{})
	// 顺序：where -> count -> Order -> limit -> offset -> data

	// where
	db = db.Where("receiver =?", m.Receiver).
		Where("is_read = ?", false)

	// count
	var count int64
	db.Count(&count)

	// order
	//如果没有排序字段
	if m.OrderBy == "" {
		//如果要求降序排列，则默认按id降序排列
		if m.Desc {
			db = db.Order("id desc")
		}
	} else { //如果有排序字段
		//先看排序字段是否存在于表中
		var tmp model.Notification
		ok := util.FieldIsInModel(db, tmp.TableName(), m.OrderBy)
		if !ok {
			return nil, nil, util.ErrorSortingFieldDoesNotExist, nil
		}
		//如果要求降序排列
		if m.Desc {
			db = db.Order(m.OrderBy + " desc")
		} else { //如果没有要求排序方式，则默认升序排列
			db = db.Order(m.OrderBy)
		}
	}

	//limit
	pageSize := global.Config.Paging.PageSize
	maxPageSize := global.Config.Paging.MaxPageSize
	if m.PageSize > 0 && m.PageSize <= maxPageSize {
		pageSize = m.PageSize
	}
	db = db.Limit(pageSize)

	//offset
	page := 1
	if m.Page > 0 {
		page = m.Page
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset)

	//原始数据
	var notifications []model.Notification
	db.Find(&notifications)

	//将结果转换为Result
	for _, notification := range notifications {
		var result NotificationResult
		result.Id = notification.Id
		if notification.BusinessType != nil {
			result.BusinessType = *notification.BusinessType
		}
		if notification.BusinessId != nil {
			result.BusinessId = *notification.BusinessId
		}
		result.Content = notification.Content

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
