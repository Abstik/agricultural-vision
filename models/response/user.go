package response

// 用户详情
type UserResponse struct {
	ID       int64             `json:"id"`
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Avatar   string            `json:"avatar"`
	Posts    []PostResponse    `json:"posts"`    // 该用户的帖子
	Comments []CommentResponse `json:"comments"` // 该用户的评论
}

// 用户简略信息（帖子和评论中展示）
type UserBriefResponse struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"` // 用户名
	Avatar   string `json:"avatar" gorm:"column:avatar"`     // 头像
}
