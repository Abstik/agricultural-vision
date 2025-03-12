package logic

import (
	"agricultural_vision/dao/mysql"
	"agricultural_vision/models"
)

func GetCommunityList() ([]*models.CommunityResponse, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetailResponse, error) {
	return mysql.GetCommunityDetailById(id)
}
