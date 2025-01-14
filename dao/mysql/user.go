package mysql

import (
	"agricultural_vision/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
)

// 对密码进行加密的盐
const secret = "agricultural_vision"

// 根据用户名查询用户是否存在
func CheckUserExist(username string) error {
	var count int64
	// 使用GORM进行查询，查找符合条件的用户数量
	err := DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return err
	}
	// 如果用户已存在，返回错误
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// 新增用户
func InsertUser(user *models.User) error {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 使用 GORM 插入数据
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 对密码进行加密
func encryptPassword(oPassword string) string {
	h := md5.New()          // 创建一个 MD5 哈希对象
	h.Write([]byte(secret)) // 向哈希对象中写入 `secret` 的字节数据
	//把 secret 的字节数据写入到 MD5 哈希的内部状态，开始计算哈希值。 相当于让 secret 成为一个固定的输入。
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
	// h.Sum([]byte(oPassword))：将 oPassword 的字节作为已有哈希值的“附加值”，生成最终的哈希
	//hex.EncodeToString：将计算出的 MD5 哈希值（16 字节）转换成一个可读的十六进制字符串，便于存储或显示。
}

// 用户登录
func Login(email, password string) (*models.User, error) {
	// 新建用户结构体，用来保存查询到的用户信息
	user := new(models.User)

	// 使用 GORM 查询用户
	err := DB.Where("email = ?", email).First(user).Error
	// 如果查询不到用户，返回用户不存在错误
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, ErrorUserNotExist
	}

	// 判断密码是否正确
	password = encryptPassword(password)
	// 如果密码不正确，返回密码不正确错误
	if password != user.Password {
		return user, ErrorInvalidPassword
	}

	return user, nil
}

// 根据用户id查询用户
func GetUserById(uid int64) (*models.User, error) {
	user := new(models.User)

	// 使用 GORM 根据 user_id 查询用户
	err := DB.Where("user_id = ?", uid).First(user).Error
	return user, err
}
