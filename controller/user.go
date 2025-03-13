package controller

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models/entity"
	"agricultural_vision/pkg/alioss"
	"agricultural_vision/response"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"agricultural_vision/logic"
	"agricultural_vision/middleware"
	"agricultural_vision/models"
	"agricultural_vision/pkg/gomail"
)

// 用户注册
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数绑定
	var p response.SignUpDTO
	err := c.ShouldBindJSON(&p)
	if err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}

	//2.业务处理
	err = logic.SingUp(&p)
	//如果出现错误
	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		//如果是邮箱已存在的错误
		if errors.Is(err, response.ErrorEmailExist) {
			response.ResponseError(c, http.StatusBadRequest, response.CodeEmailExist)
			return
		}
		//如果是邮箱验证码错误
		if errors.Is(err, response.ErrorInvalidEmailCode) {
			response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidEmailCode)
			return
		}
		//如果是其他错误，返回服务端繁忙错误信息
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	//3.返回成功响应
	response.ResponseSuccess(c, nil)
	return
}

// 用户登录
func LoginHandler(c *gin.Context) {
	//1.获取请求参数以及参数校验
	p := new(response.LoginDTO)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}

	//2.业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.String("name", p.Email), zap.Error(err))
		if errors.Is(err, response.ErrorEmailNotExist) { //如果是邮箱未注册错误
			response.ResponseError(c, http.StatusBadRequest, response.CodeEmailNotExist)
			return
		} else if errors.Is(err, response.ErrorInvalidPassword) { //如果是密码不正确错误
			response.ResponseError(c, http.StatusUnauthorized, response.CodeInvalidPassword)
			return
		} else { //否则返回服务端繁忙错误
			response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
			return
		}
	}

	//3.登陆成功，直接将token返回给用户
	response.ResponseSuccess(c, token)
	return
}

// 发送邮箱验证码
func VerifyEmailHandler(c *gin.Context) {
	// 参数绑定
	sendVerificationCodeParam := new(response.SendVerificationCodeDTO)
	if err := c.ShouldBindJSON(&sendVerificationCodeParam); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}

	// 发送邮箱验证码校验邮箱
	if err := gomail.SendVerificationCode(sendVerificationCodeParam.Email); err != nil {
		zap.L().Error("发送邮箱验证码失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, nil)
	return
}

// 修改密码
func ChangePasswordHandler(c *gin.Context) {
	// 1.获取请求参数以及参数校验
	p := new(response.ChangePasswordDTO)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}

	// 2.业务逻辑处理
	err := logic.ChangePassword(p)
	if err != nil {
		zap.L().Error("修改密码失败", zap.Error(err))
		// 如果是邮箱验证码错误
		if errors.Is(err, response.ErrorInvalidEmailCode) {
			response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidEmailCode)
			return
		}
		// 如果是邮箱未注册错误
		if errors.Is(err, response.ErrorEmailNotExist) {
			response.ResponseError(c, http.StatusBadRequest, response.CodeEmailNotExist)
			return
		}
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, nil)
	return
}

// 查询个人信息
func GetUserInfoHandler(c *gin.Context) {
	// 1.获取用户id
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	// 查询个人信息
	data, err := logic.GetUserInfo(userID)
	if err != nil {
		zap.L().Error("查询个人信息失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, data)
	return
}

// 修改个人信息
func UpdateUserInfoHandler(c *gin.Context) {
	// 1.获取请求参数以及参数校验
	p := new(response.UpdateUserInfoDTO)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("参数校验失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}

	// 2.获取用户id
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	err = logic.UpdateUserInfo(p, userID)
	if err != nil {
		zap.L().Error("修改个人信息失败", zap.Error(err))
		// 如果邮箱已注册错误
		if errors.Is(err, response.ErrorEmailExist) {
			response.ResponseError(c, http.StatusBadRequest, response.CodeEmailExist)
			return
		}
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, nil)
	return
}

// 修改头像
func UpdateUserAvatarHandler(c *gin.Context) {
	// 获取上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		zap.L().Error("获取上传文件失败", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, response.CodeInvalidParam)
		return
	}
	defer file.Close()

	// 限制文件大小（5MB）
	if header.Size > 5*1024*1024 {
		zap.L().Error("文件大小超出5MB", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, "文件大小超出5MB")
		return
	}

	// 获取文件扩展名ext
	ext := filepath.Ext(header.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		zap.L().Error("文件格式不支持", zap.Error(err))
		response.ResponseError(c, http.StatusBadRequest, "文件格式不支持")
		return
	}

	// 生成唯一文件名
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 上传到 OSS
	fileURL, err := alioss.UploadFile(file, newFileName)
	if err != nil {
		zap.L().Error("上传文件失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	// 获取用户id
	userID, err := middleware.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取userID失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	// 将头像地址更新到数据库
	err = mysql.DB.Model(&entity.User{}).Where("id = ?", userID).Update("avatar", fileURL).Error
	if err != nil {
		zap.L().Error("更新头像失败", zap.Error(err))
		response.ResponseError(c, http.StatusInternalServerError, response.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, http.StatusOK)
}
