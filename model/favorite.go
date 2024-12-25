package model

type Favorite struct {
	Base
	BusinessType string `gorm:"index;"` //业务类型，即表名，如：for_rent、seek_house、user
	BusinessId   int64  `gorm:"index;"` //业务id
}

func (f Favorite) TableName() string {
	return "favorite"
}
