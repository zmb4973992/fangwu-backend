package model

import "time"

type UserMember struct {
	Base
	MemberId                                       int64     `gorm:"index;"`
	CityCode                                       int       `gorm:"index;"`
	RemainingTimesForViewingForRentContactPerDay   int       `gorm:"index;"`
	TotalTimesForViewingForRentContactPerDay       int       `gorm:"index;"`
	RemainingTimesForPublishingSeekHouse           int       `gorm:"index;"`
	TotalTimesForPublishingSeekHouse               int       `gorm:"index;"`
	RemainingTimesForViewingSeekHouseContactPerDay int       `gorm:"index;"`
	TotalTimesForViewingSeekHouseContactPerDay     int       `gorm:"index;"`
	RemainingTimesForPublishingForRent             int       `gorm:"index;"`
	TotalTimesForPublishingForRent                 int       `gorm:"index;"`
	RemainingTopTimesPerMonth                      int       `gorm:"index;"`
	TotalTopTimesPerMonth                          int       `gorm:"index;"`
	RemainingTimesForchangingCity                  int       `gorm:"index;"`
	TotalTimesForchangingCity                      int       `gorm:"index;"`
	ExpiredAt                                      time.Time `gorm:"index;"`
}

func (u UserMember) TableName() string {
	return "user_member"
}
