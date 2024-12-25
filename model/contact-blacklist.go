package model

//记录联系方式的黑名单
type ContactBlacklist struct {
	Base
	Type        string //手机号、微信号等
	Value       string //手机号的值、微信号的值等
	Reason      string
	Description *string
}

func (c ContactBlacklist) TableName() string {
	return "contact_blacklist"
}
