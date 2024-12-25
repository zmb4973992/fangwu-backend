package model

import "time"

type RequestLog struct {
	Base
	IP           *string    `gorm:"index"`  //ip
	Location     *string    `gorm:"index;"` //所在地
	Method       *string    `gorm:"index;"` //请求方式
	Path         *string    `gorm:"index;"` //请求路径
	Remarks      *string    //备注
	ResponseCode *int       `gorm:"index;"`                  //响应码
	StartTime    *time.Time `gorm:"index;type:timestamp(3)"` //发起时间
	TimeElapsed  *int       `gorm:"index;"`                  //用时（毫秒）
	UserAgent    *string    //浏览器标识
}

// TableName 修改数据库的表名
func (r RequestLog) TableName() string {
	return "request_log"
}
