package model

type Notification struct {
	Base
	Type         int64   `gorm:"index;"` //通知类型（评论、回复我的、系统通知等）
	Receiver     int64   `gorm:"index;"` //接收者
	BusinessType *string `gorm:"index;"` //业务类型，即表名，如：for_rent、seek_house、user
	BusinessId   *int64  `gorm:"index;"` //业务id
	Content      string  //内容
	IsRead       bool    `gorm:"index;"` //是否已读
}

func (n Notification) TableName() string {
	return "notification"
}
