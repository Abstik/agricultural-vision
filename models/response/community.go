package response

// 社区简略信息
type CommunityBriefResponse struct {
	CommunityID   int64  `json:"id" db:"community_id"`
	CommunityName string `json:"name" db:"community_name"`
}

// 社区详情
type CommunityResponse struct {
	ID            int64               `json:"id"`
	CommunityName string              `json:"community_name"`
	Introduction  string              `json:"introduction"`
	Posts         []PostBriefResponse `json:"posts"`
}
