package response

import (
	"fangwu-backend/util"
)

type common struct {
	Data      any             `json:"data"`
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	ErrDetail *util.ErrDetail `json:"err_detail,omitempty"`
}

type Paging struct {
	Page         int `json:"page"`
	PageSize     int `json:"page_size"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

func GenerateSingle(data any, resCode int, errDetail *util.ErrDetail) common {
	if resCode == util.Success {
		errDetail = nil
	}

	return common{
		Data:      data,
		Code:      resCode,
		Message:   util.GetResMessage(resCode),
		ErrDetail: errDetail,
	}
}

func GenerateList(list any, paging *Paging, resCode int, errDetail *util.ErrDetail) common {
	if resCode == util.Success {
		errDetail = nil
	}

	return common{
		Data: map[string]any{
			"list":   list,
			"paging": paging,
		},
		Code:      resCode,
		Message:   util.GetResMessage(resCode),
		ErrDetail: errDetail,
	}
}
