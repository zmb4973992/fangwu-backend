package model

type Complaint struct {
	Base
	BusinessType string  `gorm:"index;"` //业务类型，如房源，用户等
	BusinessId   int64   `gorm:"index;"` //业务对象id
	Reporter     int64   `gorm:"index;"` //举报人
	Reason       int64   `gorm:"index;"` //原因
	Description  *string //描述
	Status       int64   `gorm:"index;"` //状态，如未处理、已处理等
	Response     *string //回复
}

func (c Complaint) TableName() string {
	return "complaint"
}
