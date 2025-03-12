package mysql

import (
	"agricultural_vision/models"
)

// 查询社区列表
func GetCommunityList() ([]*models.CommunityResponse, error) {
	var communities []*models.CommunityResponse

	result := DB.Model(&models.Community{}).
		Select("community_id, community_name").
		Order("created_at DESC").
		Find(&communities)

	/*// 如果未查询到结果
	// 处理空结果集
	if result.RowsAffected == 0 {
		return nil, models.ErrorNoResult
	}*/

	return communities, result.Error
}

// 根据ID获取社区详情
func GetCommunityDetailById(id int64) (*models.CommunityDetailResponse, error) {
	var communityDetail models.CommunityDetailResponse

	result := DB.Select(
		"community_id",
		"community_name",
		"introduction",
		"created_at",
	).Where("community_id = ?", id).First(&communityDetail)

	/*// 如果未查询到结果
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrorNoResult
	}*/

	return &communityDetail, result.Error
}
