package response

import "time"

// 评论
type CommentResponse struct {
	ID        int64             `json:"id"`
	Content   string            `json:"content"`           // 内容
	LikeCount int64             `json:"like_count"`        // 点赞数
	Author    UserBriefResponse `json:"author"`            // 作者
	Replies   []CommentResponse `json:"replies,omitempty"` // 子评论默认折叠
	CreatedAt time.Time         `json:"created_at"`        // 发布时间
}
