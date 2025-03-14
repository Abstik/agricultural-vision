package mysql

import (
	"agricultural_vision/models/entity"
	"agricultural_vision/models/response"
)

// 查询社区列表
func GetCommunityList() ([]*response.CommunityBriefResponse, error) {
	var communities []*response.CommunityBriefResponse

	result := DB.Model(&response.CommunityResponse{}).
		Select("community_id", "community_name").
		Order("created_at DESC"). // 根据创建时间倒序排序
		Find(&communities)

	/*// 如果未查询到结果
	// 处理空结果集
	if result.RowsAffected == 0 {
		return nil, models.ErrorNoResult
	}*/

	return communities, result.Error
}

// 根据ID获取社区详情
func GetCommunityById(id int64) (*entity.Community, error) {
	var community entity.Community

	result := DB.Where("community_id = ?", id).First(&community)

	/*// 如果未查询到结果
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrorNoResult
	}*/

	return &community, result.Error
}
