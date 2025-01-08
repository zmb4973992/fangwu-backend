package model

import (
	"fangwu-backend/global"
)

type DictionaryType struct {
	Base
	Value   string  `gorm:"index;"` //值
	Sort    *int    `gorm:"index;"` //排序
	Status  *bool   `gorm:"index;"` //状态
	Remarks *string //备注
}

func (d DictionaryType) TableName() string {
	return "dictionary_type"
}

var rawDictionaryTypes = []DictionaryType{
	{
		//for_rent、seek_house表的字典类型
		Value:   "户型",
		Sort:    IntToPointer(1),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("一室、二室、三室、四室、五室、五室以上等等"),
	},
	{
		//for_rent、seek_house表的字典类型
		Value:   "租赁类型",
		Sort:    IntToPointer(2),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("整租、合租等"),
	},
	{
		//for_rent、seek_house表的字典类型
		Value:   "性别限制",
		Sort:    IntToPointer(3),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("限男、限女、男女不限等"),
	},
	{
		//blacklist表的字典类型
		Value:   "拉黑原因",
		Sort:    IntToPointer(4),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("中介代理、虚假信息等"),
	},
	{
		//comment表的字典类型
		Value:   "评论状态",
		Sort:    IntToPointer(5),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("待审核、审核通过、审核未通过等"),
	},
	{
		//complaint表的字典类型
		Value:   "投诉原因",
		Sort:    IntToPointer(6),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("中介/代理/托管公司、色情、虚假信息/垃圾营销、其他等"),
	},
	{
		//complaint表的字典类型
		Value:   "投诉状态",
		Sort:    IntToPointer(7),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("未处理、已处理等"),
	},
	{
		//message表的字典类型
		Value:   "消息类型",
		Sort:    IntToPointer(8),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("系统消息、用户消息、评论等"),
	},
	{
		//for_rent表的字典类型
		Value:   "朝向",
		Sort:    IntToPointer(9),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("东、南、西、北、东南、东北、西南、西北、南北、东西等"),
	},
	{
		//for_rent表的字典类型
		Value:   "性别",
		Sort:    IntToPointer(10),
		Status:  BoolToPointer(true),
		Remarks: StringToPointer("先生、女士"),
	},
}

func initDictionaryType() {
	for _, dictionary_type := range rawDictionaryTypes {
		err := global.Db.
			Where("value = ?", dictionary_type.Value).
			FirstOrCreate(&dictionary_type).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}
}
