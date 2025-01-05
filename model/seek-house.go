package model

type SeekHouse struct {
	Base
	Delete
	Level1AdminDiv    *int    `json:"level_1_admin_div" gorm:"index;column:level_1_admin_div;"`
	Level2AdminDiv    *int    `json:"level_2_admin_div" gorm:"index;column:level_2_admin_div;"`
	Level3AdminDiv    *int    `json:"level_3_admin_div" gorm:"index;column:level_3_admin_div;"`
	Level4AdminDiv    *int    `json:"level_4_admin_div" gorm:"index;column:level_4_admin_div;"`
	MaxBudget         float64 `json:"max_budget" gorm:"index;"`                 //预算上限
	MinBudget         float64 `json:"min_budget" gorm:"index;"`                 //预算下限
	RentType          int64   `json:"rent_type" gorm:"index;"`                  //租赁类型，如整租、合租等
	Description       string  `json:"description"`                              //描述
	GenderRestriction int64   `json:"gender_restriction" gorm:"index;"`         //性别限制，男、女、男女不限等
	MobilePhone       *string `json:"mobile_phone" gorm:"index;"`               //手机号
	WeChatId          *string `json:"wechat_id" gorm:"column:wechat_id;index;"` //微信id
	Community         string  `json:"community" gorm:"index;"`                  //小区
	Area              *int    `json:"area" gorm:"index;"`                       //面积
}

func (s SeekHouse) TableName() string {
	return "seek_house"
}
