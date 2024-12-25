package model

import "time"

type ViewContact struct {
	Base
	BusinessType string    `gorm:"index;"`           //业务类型，即表名，如：for_rent、seek_house
	BusinessId   int64     `gorm:"index;"`           //业务id
	Date         time.Time `gorm:"index;type:date;"` //日期
}

func (v ViewContact) TableName() string {
	return "view_contact"
}
