package models

import "errors"

// 封装常见的错误信息，向上层进行返回
// 在controller层中对错误信息进行判断，向前端返回最终错误信息
var (
	ErrorUserExist        = errors.New(CodeUserExist)
	ErrorUserNotExist     = errors.New(CodeUserNotExist)
	ErrorInvalidPassword  = errors.New(CodeInvalidPassword)
	ErrorInvalidEmailCode = errors.New(CodeInvalidEmailCode)
	ErrorInvalidID        = errors.New(CodeErrorInvalidID)
)
