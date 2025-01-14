package utils

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
)

func InitSqlTable() {
	mysql.DB.AutoMigrate(&models.User{})
}
