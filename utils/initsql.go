package utils

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
)

func InitSqlTable() (err error) {
	err = mysql.DB.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.Proverb{},
		&models.CropCategory{},
		&models.CropDetail{},
		&models.Video{},
	)
	return
}
