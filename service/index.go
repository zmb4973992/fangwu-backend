package service

//getList用的入参
type list struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	OrderBy  string `json:"order_by,omitempty"` //排序字段
	Desc     bool   `json:"desc,omitempty"`     //是否为降序（从大到小）
}
