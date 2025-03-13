package request

// 发布评论
type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"` // 内容
	PostID   int64  `json:"post_id" binding:"required"` // 帖子ID
	ParentID *int64 `json:"parent_id,omitempty"`        // 父评论ID（nil表示一级评论）
}
