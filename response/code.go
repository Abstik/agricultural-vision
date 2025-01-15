package response

//封装自定义的状态码和信息

type ResCode int64

const (
	CodeSuccess          ResCode = 1000
	CodeInvalidParam     ResCode = 1001
	CodeUserExist        ResCode = 1002
	CodeUserNotExist     ResCode = 1003
	CodeInvalidPassword  ResCode = 1004
	CodeInvalidEmailCode ResCode = 1005
	CodeServerBusy       ResCode = 1006

	CodeNeedLogin     ResCode = 1007
	CodeInvalidAToken ResCode = 1008
)

// 定义状态码及其对应的信息的map
var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户名已存在",
	CodeUserNotExist:     "用户名不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeInvalidEmailCode: "邮箱验证码错误",
	CodeServerBusy:       "服务繁忙",

	CodeNeedLogin:     "需要登录",
	CodeInvalidAToken: "无效的token",
}

// 在codeMsgMap中根据键名获取值
func (resCode ResCode) GetMsg() string {
	msg, ok := codeMsgMap[resCode]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
