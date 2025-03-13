package entity

// 评论
type Comment struct {
	BaseModel
	Content string `gorm:"type:text;not null"`

	// 层级关系控制
	ParentID *int64     `gorm:"index;default:null"` // 父评论ID（null表示一级评论）
	Replies  []*Comment `gorm:"foreignKey:ParentID"`
	//RootID   *int64 `gorm:"index;default:null" json:"root_id"`   // 根评论ID（null表示自身是根评论）

	// 用户关联
	AuthorID int64 `gorm:"index;not null"`
	Author   User  `gorm:"foreignKey:AuthorID"`

	// 帖子关联
	PostID int64 `gorm:"index;not null"`

	// 记录点赞用户（多对多）
	LikedBy []User `gorm:"many2many:user_likes_comments;"`
}
