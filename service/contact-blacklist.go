package service

import (
	"fangwu-backend/global"
	"fangwu-backend/model"
	"fangwu-backend/util"
)

// value: mobile_phone/wechat_id
type ContactBlackListVerify struct {
	Type  string `json:"type" bind:"required"`
	Value string `json:"value" bind:"required"`
}

// 待添加功能
type ContactBlacklistCreate struct {
	Type        string `json:"type" bind:"required"`
	Value       string `json:"value" bind:"required"`
	Reason      string `json:"reason" bind:"required"`
	Description string `json:"description,omitempty"`
}

type ContactBlacklistGetList struct {
	Type  string `json:"type" bind:"required"`
	Value string `json:"value" bind:"required"`
}

type ContactBlacklistResult struct {
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *ContactBlackListVerify) Verify() (isBlocked bool, resCode int, errDetail *util.ErrDetail) {
	var count int64
	global.Db.Model(&model.ContactBlacklist{}).
		Where("type = ? and value = ?", c.Type, c.Value).
		Count(&count)
	return count > 0, util.Success, nil
}

func (c *ContactBlacklistGetList) GetList() (results []ContactBlacklistResult, resCode int, errDetail *util.ErrDetail) {
	global.Db.Model(&model.ContactBlacklist{}).
		Where("type = ? and value like ?", c.Type, "%"+c.Value+"%").
		Find(&results)

	return results, util.Success, nil
}
