package model

type ForRent struct {
	Base
	Delete
	AdminDiv
	RentType          int64   `json:"rent_type" gorm:"index;"`                  //租赁类型，如整租、合租等
	Price             float64 `json:"price" gorm:"index;"`                      //价格
	BuildingArea      *int    `json:"building_area" gorm:"index;"`              //建筑面积
	Description       string  `json:"description"`                              //描述
	GenderRestriction int64   `json:"gender_restriction" gorm:"index;"`         //性别限制，男、女、男女不限等
	MobilePhone       *string `json:"mobile_phone" gorm:"index;"`               //手机号
	WeChatId          *string `json:"wechat_id" gorm:"column:wechat_id;index;"` //微信id
	Community         string  `json:"community" gorm:"index;"`                  //小区
}

func (f ForRent) TableName() string {
	return "for_rent"
}
