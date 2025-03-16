package response

// 评论
type CommentResponse struct {
	ID           int64             `json:"id"`
	Content      string            `json:"content"`           // 内容
	LikeCount    int64             `json:"like_count"`        // 点赞数
	RepliesCount int64             `json:"replies_count"`     // 二级评论数
	Author       UserBriefResponse `json:"author"`            // 作者
	Replies      []CommentResponse `json:"replies,omitempty"` // 二级评论默认折叠
	CreatedAt    string            `json:"created_at"`        // 发布时间
}

// 分页查询评论响应体
type CommentListResponse struct {
	Data  []*CommentResponse `json:"comments"`
	Total int64              `json:"total"`
}
