package logic

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
	"agricultural_vision/pkg/jwt"
	"agricultural_vision/pkg/snowflake"
	"time"
)

// 用户注册
func SingUp(p *models.SignUpParam) error {
	//1.判断用户存不存在
	err := mysql.CheckUserExist(p.Username)
	if err != nil {
		//数据库查询出错
		return err
	}

	//2.生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := models.User{
		UserID:      userID,
		Username:    p.Username,
		Email:       p.Email,
		Password:    p.Password,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	//3.保存进数据库
	err = mysql.InsertUser(&user)
	return err
}

// 用户登录
func Login(p *models.LoginParam) (string, error) {
	//可以从user中拿到UserID
	user, err := mysql.Login(p.Email, p.Password)
	if err != nil {
		return "", err
	}

	//生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	return token, err
}
