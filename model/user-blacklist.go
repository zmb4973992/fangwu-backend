package model

//记录用户之间的黑名单
type UserBlacklist struct {
	Base
	Blocker int64 `gorm:"index;"` //拉黑人的user_id
	Blocked int64 `gorm:"index;"` //被拉黑人的user_id
}

func (u UserBlacklist) TableName() string {
	return "user_blacklist"
}
