package models

import "time"

// 帖子
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorID    int64     `json:"author_id" db:"author_id"`                          // 帖子作者的id（用户id）
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 帖子所属社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 帖子创建时间
}
