package model

import "time"

type UserMember struct {
	Base
	MemberId     int64     `gorm:"index;"`
	AdminDivCode int       `gorm:"index;"`
	ExpiredAt    time.Time `gorm:"type:timestamp(3)"`
}

func (u UserMember) TableName() string {
	return "user_member"
}
