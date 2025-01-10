package model

import (
	"fangwu-backend/global"
	"os"
	"time"
)

type Community struct {
	Base
	Name         string `gorm:"index;"`                                             //名称
	Code         string `gorm:"index;"`                                             //代码
	ParentCode   int    `gorm:"index;"`                                             //上级行政区划代码
	PinyinPrefix string `gorm:"index;"`                                             //拼音首字母
	SimilarName1 string `json:"similar_name_1" gorm:"index;column:similar_name_1;"` //相似名称1
	SimilarName2 string `json:"similar_name_2" gorm:"index;column:similar_name_2;"` //相似名称2
	SimilarName3 string `json:"similar_name_3" gorm:"index;column:similar_name_3;"` //相似名称3
}

func (c Community) TableName() string {
	return "community"
}

func executeSqlForCommunity() {
	//检查是否有小区的数据
	var count int64
	global.Db.Model(&Community{}).
		Count(&count)

	//如果有数据就不初始化了
	if count > 0 {
		return
	}

	//读取小区数据的sql文件
	sqlStatemnent, err := os.ReadFile("./config/community.sql")
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}

	//执行sql文件
	global.Db.Exec(string(sqlStatemnent))

	//完善小区记录的创建时间和更新时间
	err = global.Db.Model(&Community{}).
		Where("id > 0").
		Updates(map[string]any{
			"created_at":       time.Now(),
			"last_modified_at": time.Now()}).
		Error
	if err != nil {
		global.SugaredLogger.Panicln(err)
	}
}
