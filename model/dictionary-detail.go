package model

import "fangwu-backend/global"

type DictionaryDetail struct {
	Base
	DictionaryTypeId int64   `gorm:"index;"` //字典类型的Id
	Name             string  `gorm:"index;"` //名称
	Sort             *int    `gorm:"index;"` //用于排序的值
	IsValid          *bool   `gorm:"index;"` //是否启用
	ParentId         *int64  `gorm:"index;"` //父级Id
	Remarks          *string //备注
}

func (d DictionaryDetail) TableName() string {
	return "dictionary_detail"
}

type dictionaryDetailFormat struct {
	DictionaryTypeName    string
	DictionaryDetailNames []string
}

var rawDictionaryDetails = []dictionaryDetailFormat{
	{
		DictionaryTypeName:    "户型",
		DictionaryDetailNames: []string{"一室", "二室", "三室", "四室及以上"},
	}, {
		DictionaryTypeName:    "租赁类型",
		DictionaryDetailNames: []string{"整租", "合租", "整租/合租均可", "转租"},
	}, {
		DictionaryTypeName:    "性别限制",
		DictionaryDetailNames: []string{"男女不限", "限男生", "限女生"},
	}, {
		DictionaryTypeName:    "拉黑原因",
		DictionaryDetailNames: []string{"中介/代理", "虚假信息", "淫秽色情", "其他"},
	}, {
		DictionaryTypeName:    "投诉原因",
		DictionaryDetailNames: []string{"房屋已出租", "虚假信息", "中介房源", "淫秽色情", "其他"},
	},
	{
		DictionaryTypeName:    "投诉状态",
		DictionaryDetailNames: []string{"未处理", "已处理"},
	},
	{
		DictionaryTypeName:    "消息类型",
		DictionaryDetailNames: []string{"系统", "用户", "评论"},
	},
	{
		DictionaryTypeName:    "朝向",
		DictionaryDetailNames: []string{"东", "南", "西", "北", "东南", "东北", "西南", "西北", "南北", "东西"},
	},
}

func initDictionaryDetail() {
	var dictionaryDetails []DictionaryDetail
	for _, rawDictionaryDetail := range rawDictionaryDetails {
		//先找到字典类型的记录
		var dictionaryTypeInfo DictionaryType
		err = global.Db.
			Where("name = ?", rawDictionaryDetail.DictionaryTypeName).
			First(&dictionaryTypeInfo).Error
		if err != nil {
			global.SugaredLogger.Panicln(err)
		}

		//再根据字典类型的Id，将字典详情的名称转换成字典详情的记录
		for _, dictionaryDetailName := range rawDictionaryDetail.DictionaryDetailNames {
			dictionaryDetails = append(dictionaryDetails, DictionaryDetail{
				DictionaryTypeId: dictionaryTypeInfo.Id,
				Name:             dictionaryDetailName,
				Remarks:          &dictionaryTypeInfo.Name,
			})
		}
	}

	//将字典详情的记录插入到数据库中
	for _, dictionaryDetail := range dictionaryDetails {
		err = global.Db.
			Where("name = ?", dictionaryDetail.Name).
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
