package request

// 发布评论
type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"` // 内容
	PostID   int64  `json:"post_id" binding:"required"` // 帖子ID
	ParentID *int64 `json:"parent_id,omitempty"`        // 父评论ID（nil表示一级评论）
}

// 查询帖子评论
type CommentListRequest struct {
	PostID int64  `json:"post_id" binding:"required"` // 帖子ID
	Page   int64  `json:"page" form:"page"`           //查询第几页的数据
	Size   int64  `json:"size" form:"size"`           //每页数据条数
	Order  string `json:"order" form:"order"`         //排序方式
}
