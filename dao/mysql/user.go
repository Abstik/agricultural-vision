package mysql

import (
	"errors"

	"gorm.io/gorm"

	"agricultural_vision/models"
)

// 查询邮箱是否已注册
func CheckEmailExist(email string) (bool, error) {
	var count int64
	// 使用GORM进行查询，查找符合条件的用户数量
	err := DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	// 如果邮箱已存在
	if count > 0 {
		return true, nil
	}
	// 否则返回邮箱未注册
	return false, nil
}

// 新增用户
func InsertUser(user *models.User) error {
	return DB.Create(user).Error
}

// 用户登录
func Login(email, password string) (*models.User, error) {
	// 新建用户结构体，用来保存查询到的用户信息
	user := new(models.User)

	// 根据邮箱查询用户
	err := DB.Where("email = ?", email).First(user).Error
	// 如果查询不到用户，返回用户不存在错误
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, models.ErrorEmailNotExist
	}

	// 判断密码是否正确
	// 如果密码不正确，返回密码不正确错误
	if password != user.Password {
		return user, models.ErrorInvalidPassword
	}

	return user, nil
}

// 根据用户ID更新用户信息
func UpdateUserByID(user *models.User) error {
	err := DB.Model(&models.User{}).Where("id = ?", user.UserID).Updates(user).Error
	return err
}

// 根据用户ID获取用户信息
func GetUserInfo(id int64) (*models.User, error) {
	user := new(models.User)
	err := DB.Where("id = ?", id).First(user).Error
	return user, err
}

// 根据邮箱更新用户密码
func UpdatePassword(user *models.User) error {
	// 忽略零值动态更新
	err := DB.Model(&models.User{}).Where("email = ?", user.Email).Update("password", user.Password).Error
	return err
}
