package service

// 检查业务类型是否合法（business_type必须为数据库的表名）
// func businessTypeIsValid(businessType string) bool {
// 	var count int64
// 	global.Db.Table("information_schema.tables").
// 		Where("table_schema = 'public'").
// 		Where("table_name = ?", businessType).
// 		Count(&count)

// 	return count > 0
// }
