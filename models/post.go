package models

// 帖子
type Post struct {
	BaseModel
	Content     string `json:"content" gorm:"type:text;not_null;comment:帖子内容" binding:"required"`
	AuthorID    int64  `json:"author_id" gorm:"not_null;comment:帖子作者ID" binding:"required"`
	CommunityID int64  `json:"community_id" gorm:"not_null;comment:帖子所属社区ID" binding:"required"`
}
