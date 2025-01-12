package model

import "fangwu-backend/global"

type Member struct {
	Base
	Name                                  string `gorm:"index;"`
	IsValid                               bool   `gorm:"index;"` //是否启用
	TimesForViewingForRentContactPerDay   int    `gorm:"index;"` //每日允许查看出租信息联系人的次数
	TotalTimesForPublishingSeekHouse      int    `gorm:"index;"` //允许发布求租信息的总条数
	TimesForViewingSeekHouseContactPerDay int    `gorm:"index;"` //每日允许查看求租信息联系人的次数
	TotalTimesForPublishingForRent        int    `gorm:"index;"` //允许发布出租信息的总条数
	TopTimesPerMonth                      int    `gorm:"index;"` //每月置顶次数（每次1天）

}

func (m Member) TableName() string {
	return "member"
}

var rawMembers = []Member{
	{
		Name:                                "黄金会员(房客版)",
		IsValid:                             true,
		TimesForViewingForRentContactPerDay: 30,
		TotalTimesForPublishingSeekHouse:    5,
		TopTimesPerMonth:                    3,
	},
	{
		Name:                                  "黄金会员(房东版)",
		IsValid:                               true,
		TimesForViewingSeekHouseContactPerDay: 30,
		TotalTimesForPublishingForRent:        5,
		TopTimesPerMonth:                      3,
	},
}

func initMember() {
	for _, member := range rawMembers {
		err := global.Db.
			Where("name = ?", member.Name).
			FirstOrCreate(&member).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}
}
