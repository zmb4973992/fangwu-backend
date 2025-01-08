package model

type ForRent struct {
	Base
	Delete
	Level1AdminDiv    *int    `json:"level_1_admin_div" gorm:"index;column:level_1_admin_div;"`
	Level2AdminDiv    *int    `json:"level_2_admin_div" gorm:"index;column:level_2_admin_div;"`
	Level3AdminDiv    *int    `json:"level_3_admin_div" gorm:"index;column:level_3_admin_div;"`
	Level4AdminDiv    *int    `json:"level_4_admin_div" gorm:"index;column:level_4_admin_div;"`
	RentType          int64   `json:"rent_type" gorm:"index;"`                  //租赁类型，如整租、合租等
	Price             float64 `json:"price" gorm:"index;"`                      //价格
	Area              *int    `json:"area" gorm:"index;"`                       //面积
	Description       string  `json:"description"`                              //描述
	GenderRestriction int64   `json:"gender_restriction" gorm:"index;"`         //性别限制，男、女、男女不限等
	MobilePhone       *string `json:"mobile_phone" gorm:"index;"`               //手机号
	WechatId          *string `json:"wechat_id" gorm:"column:wechat_id;index;"` //微信id
	Community         string  `json:"community" gorm:"index;"`                  //小区
	Bedroom           *int    `json:"bedroom" gorm:"index;"`                    //卧室数量
	LivingRoom        *int    `json:"living_room" gorm:"index;"`                //客厅数量
	Bathroom          *int    `json:"bathroom" gorm:"index;"`                   //卫生间数量
	Kitchen           *int    `json:"kitchen" gorm:"index;"`                    //厨房数量
	Floor             *int    `json:"floor" gorm:"index;"`                      //楼层
	TotalFloor        *int    `json:"total_floor" gorm:"index;"`                //总楼层
	Orientation       *int64  `json:"orientation" gorm:"index;"`                //朝向
	Tenant            *int    `json:"tenant" gorm:"index;"`                     //合租户数
	Name              *string `json:"name"`                                     //姓名
	Gender            *int64  `json:"gender" gorm:"index;"`                     //性别
}

func (f ForRent) TableName() string {
	return "for_rent"
}
