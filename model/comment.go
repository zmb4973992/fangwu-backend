package model

type Comment struct {
	Base
	ParentId     *int64 `gorm:"index;"` //父级id
	Content      string //内容
	BusinessType string `gorm:"index;"` //业务类型，即表名，如：for_rent、seek_house、user
	BusinessId   int64  `gorm:"index;"` //业务id
	OnlyToPoster bool   //是否仅对发布者可见
	IsHidden     bool   //是否隐藏
}

func (c Comment) TableName() string {
	return "comment"
}
