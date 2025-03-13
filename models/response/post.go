package response

import (
	"agricultural_vision/models/entity"
)

// 帖子列表
type PostBriefResponse struct {
	ID           int64             `json:"id"`
	Content      string            `json:"content"`
	Image        string            `json:"image"`
	LikeCount    int64             `json:"like_count"`    // 点赞数
	CommentCount int64             `json:"comment_count"` // 评论数
	Author       UserBriefResponse `json:"author"`
}

// 帖子详情
type PostDetailResponse struct {
	Author            UserBriefResponse  `json:"author"`
	*entity.Post      `json:"post"`      // 嵌入帖子结构体
	*entity.Community `json:"community"` // 嵌入社区信息
	VoteNum           int64              `json:"vote_num"` // 赞成票数
}
