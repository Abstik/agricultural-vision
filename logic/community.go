package logic

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
	"agricultural_vision/models/entity"
)

func GetCommunityList() ([]*response.CommunityResponse, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*entity.Community, error) {
	return mysql.GetCommunityDetailById(id)
}
