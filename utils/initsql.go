package utils

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models/entity"
)

func InitSqlTable() (err error) {
	err = mysql.DB.AutoMigrate(
		&entity.User{},

		&entity.News{},
		&entity.Proverb{},
		&entity.CropCategory{},
		&entity.CropDetail{},
		&entity.Video{},
		&entity.Poetry{},

		&entity.Post{},
		&entity.Community{},
		&entity.Comment{},
	)
	community1 := &entity.Community{
		CommunityName: "java",
		Introduction:  "java是一门面向对象的语言",
	}
	community2 := &entity.Community{
		CommunityName: "go",
		Introduction:  "go是一种简洁高效的语言",
	}
	community3 := &entity.Community{
		CommunityName: "python",
		Introduction:  "python是一门动态语言",
	}
	mysql.DB.Create(community1)
	mysql.DB.Create(community2)
	mysql.DB.Create(community3)

	return
}
