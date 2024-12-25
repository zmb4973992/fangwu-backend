package model

type Footprint struct {
	Base
	Delete
	BusinessType string `json:"business_type" gorm:"index;"` //业务类型，如出租、求租等
	BusinessId   int64  `json:"business_id" gorm:"index;"`   //业务id，如出租id、求租id等
}

func (f Footprint) TableName() string {
	return "footprint"
}
