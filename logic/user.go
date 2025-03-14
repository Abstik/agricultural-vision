package logic

import (
	"agricultural_vision/constants"
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models/entity"
	"agricultural_vision/models/request"
	"agricultural_vision/pkg/gomail"
	auth "agricultural_vision/pkg/jwt"
	"agricultural_vision/pkg/md5"
)

// 用户注册
func SingUp(p *request.SignUpRequest) error {
	// 1.判断邮箱是否已注册
	flag, err := mysql.CheckEmailExist(p.Email)
	// 如果数据库查询出错
	if err != nil {
		return err
	}
	// 如果邮箱已注册
	if flag {
		return constants.ErrorEmailExist
	}

	// 2.校验邮箱
	if err = gomail.VerifyVerificationCode(p.Email, p.Code); err != nil {
		return constants.ErrorInvalidEmailCode
	}

	user := entity.User{
		Username: p.Username,
		Password: md5.EncryptPassword(p.Password),
		Email:    p.Email,
	}

	// 3.保存进数据库
	err = mysql.InsertUser(&user)
	return err
}

// 修改密码
func ChangePassword(p *request.ChangePasswordRequest) error {
	// 验证邮箱是否已注册
	flag, err := mysql.CheckEmailExist(p.Email)
	// 如果数据库查询出错
	if err != nil {
		return err
	}
	// 如果邮箱未注册
	if !flag {
		return constants.ErrorEmailNotExist
	}

	// 验证邮箱验证码是否正确
	if err = gomail.VerifyVerificationCode(p.Email, p.Code); err != nil {
		return constants.ErrorInvalidEmailCode
	}

	// 修改密码
	// 先对密码明文进行加密
	p.Password = md5.EncryptPassword(p.Password)
	user := entity.User{
		Password: p.Password,
		Email:    p.Email,
	}

	// 再更新数据库
	return mysql.UpdatePassword(&user)
}

// 用户登录
func Login(p *request.LoginRequest) (string, error) {
	//可以从user中拿到UserID
	user, err := mysql.Login(p.Email, md5.EncryptPassword(p.Password))
	if err != nil {
		return "", err
	}

	//生成JWT
	token, err := auth.GenToken(user.ID, user.Username)
	return token, err
}

// 获取用户信息
func GetUserInfo(id int64) (*entity.User, error) {
	return mysql.GetUserInfo(id)
}

// 更新用户信息
func UpdateUserInfo(p *request.UpdateUserInfoRequest, id int64) error {
	// 1. 邮箱校验
	// 查询用户原本的邮箱
	user, err := mysql.GetUserInfo(id)
	if err != nil {
		return err
	}
	// 判断邮箱是否已注册
	flag, err := mysql.CheckEmailExist(p.Email)
	if err != nil {
		return err
	}
	// 如果新邮箱和原本邮箱不同，且新邮箱已被其他用户注册
	if user.Email != p.Email && flag {
		return constants.ErrorEmailExist
	}

	// 2. 更新用户信息
	newUser := entity.User{
		Username: p.Username,
		Email:    p.Email,
		Avatar:   p.Avatar,
	}

	return mysql.UpdateUserByID(&newUser)
}
