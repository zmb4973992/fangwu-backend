package model

type ViewContact struct {
	Base
	BusinessType string `gorm:"index;"` //业务类型，即表名，如：for_rent、seek_house
	BusinessId   int64  `gorm:"index;"` //业务id
}

func (v ViewContact) TableName() string {
	return "view_contact"
}
