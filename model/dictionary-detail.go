package model

import "fangwu-backend/global"

type DictionaryDetail struct {
	Base
	DictionaryTypeId int64   `gorm:"index;"` //字典类型的Id
	Value            string  `gorm:"index;"` //值
	Sort             *int    `gorm:"index;"` //用于排序的值
	IsValid          *bool   `gorm:"index;"` //是否启用
	ParentId         *int64  `gorm:"index;"` //父级Id
	Remarks          *string //备注
}

func (d DictionaryDetail) TableName() string {
	return "dictionary_detail"
}

type dictionaryDetailFormat struct {
	DictionaryTypeValue    string
	DictionaryDetailValues []string
}

var rawDictionaryDetails = []dictionaryDetailFormat{
	{
		DictionaryTypeValue:    "户型",
		DictionaryDetailValues: []string{"一室", "二室", "三室", "四室及以上"},
	}, {
		DictionaryTypeValue:    "租赁类型",
		DictionaryDetailValues: []string{"整租", "合租"},
	}, {
		DictionaryTypeValue:    "性别限制",
		DictionaryDetailValues: []string{"男女不限", "限男生", "限女生"},
	}, {
		DictionaryTypeValue:    "拉黑原因",
		DictionaryDetailValues: []string{"中介/代理", "虚假信息", "淫秽色情", "其他"},
	}, {
		DictionaryTypeValue:    "投诉原因",
		DictionaryDetailValues: []string{"房屋已出租", "虚假信息", "中介房源", "淫秽色情", "其他"},
	},
	{
		DictionaryTypeValue:    "投诉状态",
		DictionaryDetailValues: []string{"未处理", "已处理"},
	},
	{
		DictionaryTypeValue:    "消息类型",
		DictionaryDetailValues: []string{"系统", "用户", "评论"},
	},
	{
		DictionaryTypeValue:    "朝向",
		DictionaryDetailValues: []string{"东", "南", "西", "北", "东南", "东北", "西南", "西北", "南北", "东西"},
	},
	{
		DictionaryTypeValue:    "性别",
		DictionaryDetailValues: []string{"先生", "女士"},
	},
}

func initDictionaryDetail() {
	var dictionaryDetails []DictionaryDetail
	for _, rawDictionaryDetail := range rawDictionaryDetails {
		//先找到字典类型的记录
		var dictionaryTypeInfo DictionaryType
		err = global.Db.
			Where("value = ?", rawDictionaryDetail.DictionaryTypeValue).
			First(&dictionaryTypeInfo).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}

		//再根据字典类型的Id，将字典详情的名称转换成字典详情的记录
		for _, dictionaryDetailValue := range rawDictionaryDetail.DictionaryDetailValues {
			dictionaryDetails = append(dictionaryDetails, DictionaryDetail{
				DictionaryTypeId: dictionaryTypeInfo.Id,
				Value:            dictionaryDetailValue,
				Remarks:          &dictionaryTypeInfo.Value,
			})
		}
	}

	//将字典详情的记录插入到数据库中
	for _, dictionaryDetail := range dictionaryDetails {
		err = global.Db.
			Where("value = ?", dictionaryDetail.Value).
			Where("dictionary_type_id = ?", dictionaryDetail.DictionaryTypeId).
			//将IsValid设置为true
			Attrs(&DictionaryDetail{
				IsValid: BoolToPointer(true),
			}).
			FirstOrCreate(&dictionaryDetail).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}
	}
}
