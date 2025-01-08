package util

import (
	"runtime"
	"strconv"
)

// 自定义的错误代码
const (
	Success             = iota
	ErrorUsernameExists = iota + 1000
	ErrorFailToEncrypt
	ErrorRequestTooFast
	ErrrorNeedLogin
	ErrorRateLimitDoesNotWork
	ErrorFailToUpload
	ErrorInvalidRequest
	ErrorFileTooLarge
	ErrorFailToOepenFile
	ErrorFailToOpenFile
	ErrorFailToCreateFile
	ErrorFailToCopyFile
	ErrorFailToCreateFileRecord
	ErrorFailToCreateForRent
	ErrorFailToCreateUser
	ErrorFailToUpdateUser
	ErrorFailToCreateSeekHouse
	ErrorFailToMoveFile
	ErrorFailToSaveFile
	ErrorFailToDeleteFile
	ErrorAccessTokenNotFound
	ErrorInvalidAccessToken
	ErrorFailToParseAccessToken
	ErrorAccessTokenExpired
	ErrorFailToParseMultipartForm
	ErrorFailToGenerateToken
	ErrorFailToSignToken
	ErrorFailToDecodePng
	ErrorFailToDecodeBmp
	ErrorFailToEncodeJpg
	ErrorNonJpgFile
	ErrorInvalidBusinessType
	ErrorInvalidUriParams
	ErrorInvalidJsonParams
	ErrorFailToUpdateFileRecord
	ErrorFailToDeleteUser
	ErrorFailToDeleteForRent
	ErrorFailToDeleteSeekHouse
	ErrorFailToDeleteFileRecord
	ErrorFailToGetFileRecord
	ErrorFailToGetForRent
	ErrorFailToGetSeekHouse
	ErrorFailToGetUser
	ErrorInvalidUsernameOrPassword
	ErrorFileNotFound
	ErrorWrongCaptcha
	ErrorSortingFieldDoesNotExist
	ErrorFailToGetDictionaryDetail
	ErrorUnsupportedFileType
	ErrorFailToUpdateForRent
	ErrorFailToUpdateSeekHouse
	ErrorFailToGetDictionaryType
	ErrorFailToGetComplaint
	ErrorFailToCreateComplaint
	ErrorFailToCreateComment
	ErrorFailToUpdateComment
	ErrorFailToDeleteComment
	ErrorFailToCreateUserBlacklist
	ErrorFailToDeleteUserBlacklist
	ErrorFailToCreateNotification
	ErrorFailToCreateCaptcha
	ErrorMissingCaptcha
	ErrorFailToGetFavorite
	ErrorFailToCreateFavorite
	ErrorFailToDeleteFavorite
	ErrorFailToDeleteComplaint
	ErrorFailToGetFootprint
	ErrorFailToCreateFootprint
	ErrorFailToDeleteFootprint
	ErrorFailToUpdateFootprint
	ErrorTooFrequentRegistration
	ErrorMobilePhoneIsInBlacklist
	ErrorWechatIdIsInBlacklist
	ErrorFailToGetAdminDiv
	ErrorInvalidLimitedBy
	ErrorInvalidTimeUnit
	ErrorInvalidId
	ErrorFailToUpdateNotification
	ErrorInvalidDate
	ErrorFailToCreateViewContact
	ErrorFailToDeleteViewContact
	ErrorFailToGetViewContact
)

// Message 自定义错误的message
var Message = map[int]string{
	Success:                        "成功",
	ErrorUsernameExists:            "用户名已存在",
	ErrorFailToEncrypt:             "加密失败",
	ErrorRequestTooFast:            "请求过快，请稍后再试",
	ErrrorNeedLogin:                "请登录后再试",
	ErrorRateLimitDoesNotWork:      "限流功能发生故障",
	ErrorFailToUpload:              "上传失败",
	ErrorInvalidRequest:            "请求路径无效",
	ErrorFileTooLarge:              "文件过大",
	ErrorFailToOpenFile:            "打开文件失败",
	ErrorFailToCreateFile:          "创建文件失败",
	ErrorFailToCopyFile:            "复制文件失败",
	ErrorFailToCreateForRent:       "添加出租记录失败",
	ErrorFailToCreateUser:          "添加用户失败",
	ErrorFailToUpdateUser:          "修改用户失败",
	ErrorFailToCreateSeekHouse:     "添加求租记录失败",
	ErrorFailToMoveFile:            "移动文件失败",
	ErrorFailToSaveFile:            "保存文件失败",
	ErrorFailToDeleteFile:          "删除文件失败",
	ErrorAccessTokenNotFound:       "请先登录",
	ErrorInvalidAccessToken:        "access_token无效",
	ErrorFailToParseAccessToken:    "解析access_token失败",
	ErrorAccessTokenExpired:        "access_token已过期",
	ErrorFailToParseMultipartForm:  "解析multipart/form-data失败",
	ErrorFailToGenerateToken:       "生成token失败",
	ErrorFailToSignToken:           "token签名失败",
	ErrorFailToDecodePng:           "解码png失败",
	ErrorFailToDecodeBmp:           "解码bmp失败",
	ErrorFailToEncodeJpg:           "编码jpg失败",
	ErrorNonJpgFile:                "非jpg文件",
	ErrorInvalidBusinessType:       "业务类型无效",
	ErrorFailToGetForRent:          "获取出租记录失败",
	ErrorFailToGetSeekHouse:        "获取求租记录失败",
	ErrorFailToGetUser:             "获取用户失败",
	ErrorFailToGetFileRecord:       "获取文件记录失败",
	ErrorInvalidUriParams:          "URI参数无效",
	ErrorInvalidJsonParams:         "JSON参数无效",
	ErrorFailToDeleteUser:          "删除用户失败",
	ErrorFailToDeleteForRent:       "删除出租记录失败",
	ErrorFailToDeleteSeekHouse:     "删除求租记录失败",
	ErrorFailToDeleteFileRecord:    "删除文件记录失败",
	ErrorFailToUpdateFileRecord:    "修改文件记录失败",
	ErrorInvalidUsernameOrPassword: "用户名或密码错误",
	ErrorFileNotFound:              "文件未找到",
	ErrorWrongCaptcha:              "验证码错误",
	ErrorSortingFieldDoesNotExist:  "排序字段不存在",
	ErrorFailToGetDictionaryDetail: "获取字典详情失败",
	ErrorUnsupportedFileType:       "不支持的文件类型",
	ErrorFailToUpdateForRent:       "修改出租记录失败",
	ErrorFailToUpdateSeekHouse:     "修改求租记录失败",
	ErrorFailToGetDictionaryType:   "获取字典类型失败",
	ErrorFailToGetComplaint:        "获取投诉失败",
	ErrorFailToCreateComplaint:     "创建投诉失败",
	ErrorFailToCreateComment:       "创建评论失败",
	ErrorFailToUpdateComment:       "修改评论失败",
	ErrorFailToDeleteComment:       "删除评论失败",
	ErrorFailToCreateUserBlacklist: "创建用户黑名单失败",
	ErrorFailToDeleteUserBlacklist: "删除用户黑名单失败",
	ErrorFailToCreateNotification:  "创建通知失败",
	ErrorFailToCreateCaptcha:       "创建验证码失败",
	ErrorMissingCaptcha:            "缺少验证码",
	ErrorFailToGetFavorite:         "获取收藏失败",
	ErrorFailToCreateFavorite:      "创建收藏失败",
	ErrorFailToDeleteFavorite:      "删除收藏失败",
	ErrorFailToDeleteComplaint:     "删除投诉失败",
	ErrorFailToGetFootprint:        "获取足迹失败",
	ErrorFailToCreateFootprint:     "创建足迹失败",
	ErrorFailToDeleteFootprint:     "删除足迹失败",
	ErrorFailToUpdateFootprint:     "修改足迹失败",
	ErrorTooFrequentRegistration:   "注册过于频繁",
	ErrorMobilePhoneIsInBlacklist:  "手机号在黑名单中",
	ErrorWechatIdIsInBlacklist:     "微信号在黑名单中",
	ErrorFailToGetAdminDiv:         "获取行政区划失败",
	ErrorInvalidLimitedBy:          "无效的limitedBy参数",
	ErrorInvalidTimeUnit:           "无效的timeUnit参数",
	ErrorInvalidId:                 "无效的id",
	ErrorFailToUpdateNotification:  "更新通知失败",
	ErrorInvalidDate:               "日期无效",
	ErrorFailToCreateViewContact:   "创建浏览联系方式记录失败",
	ErrorFailToDeleteViewContact:   "删除浏览联系方式记录失败",
	ErrorFailToGetViewContact:      "获取浏览联系方式记录失败",
}

type ErrDetail struct {
	FileName     string `json:"file_name"`
	FunctionName string `json:"function_name"`
	Line         int    `json:"line"`
	Description  string `json:"description"`
}

func GetErrDetail(err error) *ErrDetail {
	if err == nil {
		return nil
	}

	discription := err.Error()

	// 获取调用栈信息
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return &ErrDetail{
			FileName:    "runtime.Caller出现问题，无法获取代码位置",
			Description: discription,
		}
	}

	functionName := runtime.FuncForPC(pc).Name()
	return &ErrDetail{
		FileName:     file,
		FunctionName: functionName,
		Line:         line,
		Description:  discription,
	}
}

func GetResMessage(code int) string {
	message, ok := Message[code]
	if !ok {
		return "当前错误代码为：" + strconv.Itoa(code) +
			"。由于错误代码未定义返回信息，导致获取错误信息失败，" +
			"请检查后端util/error里的Message变量。"
	}
	return message
}
