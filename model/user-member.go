package model

import "time"

type UserMember struct {
	Base
	UserId    int64     `gorm:"index;"`
	MemberId  int64     `gorm:"index;"`
	CityCode  int       `gorm:"index;"`
	ExpiredAt time.Time `gorm:"index;"`
}

func (u UserMember) TableName() string {
	return "user_member"
}
