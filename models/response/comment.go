package response

// 评论
type CommentResponse struct {
	ID           int64             `json:"id"`
	Content      string            `json:"content"`           // 内容
	Author       UserBriefResponse `json:"author"`            // 作者
	LikeCount    int64             `json:"like_count"`        // 点赞数
	Liked        bool              `json:"liked"`             // 当前用户是否已点赞
	RepliesCount int64             `json:"replies_count"`     // 二级评论数
	Replies      []CommentResponse `json:"replies,omitempty"` // 二级评论默认折叠
	CreatedAt    string            `json:"created_at"`        // 发布时间
}

// 分页查询评论响应体
type CommentListResponse struct {
	Data  []*CommentResponse `json:"comments"`
	Total int64              `json:"total"`
}
