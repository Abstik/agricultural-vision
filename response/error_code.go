package response

import "errors"

// 封装常见的错误信息，向上层进行返回
// 在controller层中对错误信息进行判断，向前端返回最终错误信息
var (
	ErrorUserExist        = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("用户不存在")
	ErrorInvalidPassword  = errors.New("密码错误")
	ErrorInvalidEmailCode = errors.New("邮箱验证码错误")
	ErrorInvalidID        = errors.New("无效的ID")
)
