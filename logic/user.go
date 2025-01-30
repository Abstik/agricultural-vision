package logic

import (
	"time"

	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
	"agricultural_vision/pkg/gomail"
	auth "agricultural_vision/pkg/jwt"
	"agricultural_vision/pkg/md5"
)

// 用户注册
func SingUp(p *models.SignUpParam) error {
	// 判断邮箱是否已注册
	flag, err := mysql.CheckEmailExist(p.Email)
	//如果数据库查询出错
	if err != nil {
		return err
	}
	// 如果邮箱已注册
	if flag {
		return models.ErrorEmailExist
	}

	user := models.User{
		Username:    p.Username,
		Password:    md5.EncryptPassword(p.Password),
		Email:       p.Email,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	// 保存进数据库
	err = mysql.InsertUser(&user)
	return err
}

// 用户登录
func Login(p *models.LoginParam) (string, error) {
	//可以从user中拿到UserID
	user, err := mysql.Login(p.Email, md5.EncryptPassword(p.Password))
	if err != nil {
		return "", err
	}

	//生成JWT
	token, err := auth.GenToken(user.UserID, user.Username)
	return token, err
}

// 修改密码
func ChangePassword(p *models.ChangePasswordParam) error {
	// 验证邮箱是否已注册
	flag, err := mysql.CheckEmailExist(p.Email)
	// 如果数据库查询出错
	if err != nil {
		return err
	}
	// 如果邮箱未注册
	if !flag {
		return models.ErrorEmailNotExist
	}

	// 验证邮箱验证码是否正确
	if err = gomail.VerifyVerificationCode(p.Email, p.Code); err != nil {
		return models.ErrorInvalidEmailCode
	}

	// 修改密码
	// 先对密码明文进行加密
	p.Password = md5.EncryptPassword(p.Password)
	user := models.User{
		Password: p.Password,
		Email:    p.Email,
	}

	// 再更新数据库
	return mysql.UpdatePassword(&user)
}

// 获取用户信息
func GetUserInfo(id int64) (*models.User, error) {
	return mysql.GetUserInfo(id)
}

// 更新用户信息
func UpdateUserInfo(p *models.UpdateUserInfoParam, id int64) error {
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
		return models.ErrorEmailExist
	}

	// 2. 更新用户信息
	newUser := models.User{
		Username:    p.Username,
		Password:    md5.EncryptPassword(p.Password),
		Email:       p.Email,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	return mysql.UpdateUserByID(&newUser)
}
