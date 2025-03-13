package entity

// 帖子
type Post struct {
	BaseModel
	Content string `json:"content" gorm:"type:text;not_null"`
	Image   string `gorm:"type:varchar(512)" json:"image"`

	// 用户关联（BelongsTo关系）
	AuthorID int64 `gorm:"index;not null" json:"author_id" binding:"required"`
	Author   User  `gorm:"foreignKey:AuthorID" json:"author"` // 实现预加载用户信息

	// 社区关联（BelongsTo关系）
	CommunityID int64     `gorm:"index;not null" json:"community_id" binding:"required"`
	Community   Community `gorm:"foreignKey:CommunityID" json:"-"` // 实现预加载社区信息

	// 评论关联（HasMany关系）
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"` // 定义关联关系
}
