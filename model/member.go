package model

import "fangwu-backend/global"

type Member struct {
	Base
	Type1                                      string `gorm:"index;"` //会员类型1：普通用户、会员(租客版)、会员(房东版)
	Type2                                      string `gorm:"index;"` //会员类型2：月卡、季卡、年卡
	IsValid                                    bool   `gorm:"index;"` //是否启用
	ValidityDays                               int    `gorm:"index;"` //有效期（天）
	TotalTimesForViewingForRentContactPerDay   int    `gorm:"index;"` //每日允许查看出租信息联系人的总次数
	TotalTimesForPublishingSeekHouse           int    `gorm:"index;"` //允许发布求租信息的总条数
	TotalTimesForPinningSeekHouse              int    `gorm:"index;"` //置顶求租信息的总次数（每次1天）
	TotalTimesForViewingSeekHouseContactPerDay int    `gorm:"index;"` //每日允许查看求租信息联系人的总次数
	TotalTimesForPublishingForRent             int    `gorm:"index;"` //允许发布出租信息的总条数
	TotalTimesForPinningForRent                int    `gorm:"index;"` //置顶出租信息的总次数（每次1天）
	TotalTimesForSwitchingCity                 int    `gorm:"index;"` //切换城市的总次数
}

func (m Member) TableName() string {
	return "member"
}

var rawMembers = []Member{
	{
		Type1:                                    "普通用户",
		Type2:                                    "",
		IsValid:                                  true,
		ValidityDays:                             0,
		TotalTimesForViewingForRentContactPerDay: 3,
		TotalTimesForPublishingSeekHouse:         1,
		TotalTimesForPinningSeekHouse:            0,
		TotalTimesForViewingSeekHouseContactPerDay: 3,
		TotalTimesForPublishingForRent:             1,
		TotalTimesForPinningForRent:                0,
		TotalTimesForSwitchingCity:                 0,
	},
	{
		Type1:                                    "会员(租客版)",
		Type2:                                    "月卡",
		IsValid:                                  true,
		ValidityDays:                             31,
		TotalTimesForViewingForRentContactPerDay: 20,
		TotalTimesForPublishingSeekHouse:         3,
		TotalTimesForPinningSeekHouse:            1,
		TotalTimesForViewingSeekHouseContactPerDay: 3,
		TotalTimesForPublishingForRent:             1,
		TotalTimesForPinningForRent:                0,
		TotalTimesForSwitchingCity:                 0,
	},
	{
		Type1:                                    "会员(租客版)",
		Type2:                                    "季卡",
		IsValid:                                  true,
		ValidityDays:                             91,
		TotalTimesForViewingForRentContactPerDay: 30,
		TotalTimesForPublishingSeekHouse:         5,
		TotalTimesForPinningSeekHouse:            4,
		TotalTimesForViewingSeekHouseContactPerDay: 3,
		TotalTimesForPublishingForRent:             1,
		TotalTimesForPinningForRent:                0,
		TotalTimesForSwitchingCity:                 0,
	},
	{
		Type1:                                    "会员(租客版)",
		Type2:                                    "年卡",
		IsValid:                                  true,
		ValidityDays:                             365,
		TotalTimesForViewingForRentContactPerDay: 50,
		TotalTimesForPublishingSeekHouse:         7,
		TotalTimesForPinningSeekHouse:            20,
		TotalTimesForViewingSeekHouseContactPerDay: 3,
		TotalTimesForPublishingForRent:             1,
		TotalTimesForPinningForRent:                0,
		TotalTimesForSwitchingCity:                 1,
	},
	{
		Type1:                                    "会员(房东版)",
		Type2:                                    "月卡",
		IsValid:                                  true,
		ValidityDays:                             31,
		TotalTimesForViewingForRentContactPerDay: 3,
		TotalTimesForPublishingSeekHouse:         1,
		TotalTimesForPinningSeekHouse:            0,
		TotalTimesForViewingSeekHouseContactPerDay: 20,
		TotalTimesForPublishingForRent:             3,
		TotalTimesForPinningForRent:                1,
		TotalTimesForSwitchingCity:                 0,
	},
	{
		Type1:                                    "会员(房东版)",
		Type2:                                    "季卡",
		IsValid:                                  true,
		ValidityDays:                             91,
		TotalTimesForViewingForRentContactPerDay: 3,
		TotalTimesForPublishingSeekHouse:         1,
		TotalTimesForPinningSeekHouse:            0,
		TotalTimesForViewingSeekHouseContactPerDay: 30,
		TotalTimesForPublishingForRent:             5,
		TotalTimesForPinningForRent:                4,
		TotalTimesForSwitchingCity:                 0,
	},
	{
		Type1:                                    "会员(房东版)",
		Type2:                                    "年卡",
		IsValid:                                  true,
		ValidityDays:                             365,
		TotalTimesForViewingForRentContactPerDay: 3,
		TotalTimesForPublishingSeekHouse:         1,
		TotalTimesForPinningSeekHouse:            0,
		TotalTimesForViewingSeekHouseContactPerDay: 50,
		TotalTimesForPublishingForRent:             7,
		TotalTimesForPinningForRent:                20,
		TotalTimesForSwitchingCity:                 1,
	},
}

func initMember() {
	for _, member := range rawMembers {
		err := global.Db.
			Where("type1 = ?", member.Type1).
			Where("type2 = ?", member.Type2).
			FirstOrCreate(&member).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}
}
