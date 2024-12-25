package model

import (
	"fangwu-backend/global"
	"os"
	"time"
)

type AdministrativeDivision struct {
	Base
	Code         int    `gorm:"index;"` //行政区划代码
	ParentCode   int    `gorm:"index;"` //上级行政区划代码
	Name         string `gorm:"index;"` //名称
	Level        int    `gorm:"index;"` //层级
	PinyinPrefix string `gorm:"index;"` //拼音首字母
}

func (a AdministrativeDivision) TableName() string {
	return "administrative_division"
}

func executeSqlForAdministrativeDivisions() {
	//检查是否有行政区划的数据
	var count int64
	global.Db.Model(&AdministrativeDivision{}).
		Count(&count)

	//如果有数据就不初始化了
	if count > 0 {
		return
	}

	//读取行政区划数据的sql文件
	sqlStatemnent, err := os.ReadFile("./config/administrative-division.sql")
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}

	//执行sql文件
	global.Db.Exec(string(sqlStatemnent))

	//完善行政区划记录的创建时间和更新时间
	err = global.Db.Model(&AdministrativeDivision{}).
		Where("id > 0").
		Updates(map[string]any{
			"created_at":       time.Now(),
			"last_modified_at": time.Now()}).
		Error
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}
}
