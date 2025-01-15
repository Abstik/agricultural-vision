package models

//封装自定义的状态码和信息

type ResCode int64

const (
	CodeSuccess          string = "success"
	CodeInvalidParam     string = "请求参数错误"
	CodeUserExist        string = "用户名已存在"
	CodeUserNotExist     string = "用户名不存在"
	CodeInvalidPassword  string = "用户名或密码错误"
	CodeInvalidEmailCode string = "邮箱验证码错误"
	CodeServerBusy       string = "服务繁忙"
	CodeErrorInvalidID   string = "无效的ID"

	CodeNeedLogin     string = "邮箱验证码错误"
	CodeInvalidAToken string = "服务繁忙"
)

// 定义状态码及其对应的信息的map
//var codeMsgMap = map[ResCode]string{
//	CodeSuccess:          "success",
//	CodeInvalidParam:     "请求参数错误",
//	CodeUserExist:        "用户名已存在",
//	CodeUserNotExist:     "用户名不存在",
//	CodeInvalidPassword:  "用户名或密码错误",
//	CodeInvalidEmailCode: "邮箱验证码错误",
//	CodeServerBusy:       "服务繁忙",
//
//	CodeNeedLogin:     "需要登录",
//	CodeInvalidAToken: "无效的token",
//}

// 在codeMsgMap中根据键名获取值
//func (resCode ResCode) GetMsg() string {
//	msg, ok := codeMsgMap[resCode]
//	if !ok {
//		msg = codeMsgMap[CodeServerBusy]
//	}
//	return msg
//}
