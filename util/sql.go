package util

import (
	"fmt"

	"gorm.io/gorm"
)

func FieldIsInModel(db *gorm.DB, model string, field string) bool {
	var count int64

	sqlStatement := fmt.Sprintf("select column_name from information_schema.columns where table_schema = 'public' and table_name = '%s' and column_name = '%s'",
		model, field)
	fmt.Println(sqlStatement)
	db.Raw(sqlStatement).Count(&count)
	fmt.Println(count)
	return count > 0
}

func GetNumberOfPages(numberOfRecords int, pageSize int) (numberOfPages int) {
	//如果记录数不合法
	if numberOfRecords < 0 {
		return 0
	}

	//如果记录数为零（无记录）：
	if numberOfRecords == 0 {
		return 0
	}

	//如果单页条数不合法：
	if pageSize < 0 {
		return 0
	}

	//如果单页条数为零（不分页）：
	if pageSize == 0 {
		return 1
	}

	//计算页数：
	numberOfPages = numberOfRecords / pageSize
	if numberOfRecords%pageSize != 0 {
		numberOfPages++
	}
	return numberOfPages
}
