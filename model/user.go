package model

type User struct {
	Base
	Username               string `gorm:"index;"`
	Password               string
	IsValid                *bool   `gorm:"index;"` //是否有效
	EmailAddress           *string `gorm:"index;"` //邮箱地址
	MobilePhone            *string `gorm:"index;"` //手机号
	RegisterIp             *string `gorm:"index;"` //注册ip3
	TimesForViewingContact int     `gorm:"index;"` //每日允许查看联系人的次数
}

func (u User) TableName() string {
	return "user"
}
