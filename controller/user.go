package controller

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/logic"
	"agricultural_vision/models"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户注册
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数绑定
	var p models.SignUpParam
	err := c.ShouldBindJSON(&p)
	if err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2.业务处理
	err = logic.SingUp(&p)
	//如果出现错误
	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		//如果是用户已存在的错误
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		//如果是其他错误，返回服务端繁忙错误信息
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回成功响应
	ResponseSuccess(c, nil)
	return
}

// 用户登录
func LoginHandler(c *gin.Context) {
	//1.获取请求参数以及参数校验
	p := new(models.LoginParam)
	err := c.ShouldBindJSON(p)
	if err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2.业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.String("email", p.Email), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) { //如果是用户不存在错误
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) { //如果是密码不正确错误
			ResponseError(c, CodeInvalidPassword)
			return
		} else { //否则返回服务端繁忙错误
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	//3.登陆成功，直接将token返回给用户
	ResponseSuccess(c, token)
	return
}