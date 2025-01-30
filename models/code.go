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

	CodeNeedLogin     string = "需要登录"
	CodeInvalidAToken string = "无效的token"
)
