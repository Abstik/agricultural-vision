package models

//封装自定义的状态码和信息

type ResCode int64

const (
	CodeSuccess          string = "success"
	CodeInvalidParam     string = "请求参数错误"
	CodeEmailExist       string = "此邮箱已注册"
	CodeEmailNotExist    string = "此邮箱未注册"
	CodeInvalidPassword  string = "邮箱或密码错误"
	CodeInvalidEmailCode string = "邮箱验证码错误"
	CodeNeedLogin        string = "用户需要登录"
	CodeInvalidAToken    string = "无效的token"
	CodeServerBusy       string = "服务繁忙"
)
