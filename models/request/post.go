package request

// 发布帖子
type CreatePostRequest struct {
	Content     string `json:"content" binding:"required"`      // 内容
	Image       string `json:"image,omitempty"`                 // 图片（可选）
	CommunityID int64  `json:"community_id" binding:"required"` // 归属社区
}

// 批量查询帖子
type PostListRequest struct {
	Page  int64  `json:"page" form:"page"`   //查询第几页的数据
	Size  int64  `json:"size" form:"size"`   //每页数据条数
	Order string `json:"order" form:"order"` //排序方式
}

// 批量查询帖子（根据社区id查询）
type CommunityPostListRequest struct {
	PostListRequest       //嵌入获取帖子列表的参数的结构体
	CommunityID     int64 `json:"community_id" form:"community_id" binding:"required"` //社区id
}
