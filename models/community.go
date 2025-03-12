package models

import (
	"gorm.io/gorm"
)

// 社区结构体
type Community struct {
	gorm.Model
	CommunityID   uint   `gorm:"type:int unsigned;not_null;uniqueIndex:idx_community_id"`
	CommunityName string `gorm:"type:varchar(128);not_null;uniqueIndex:idx_community_name"`
	Introduction  string `gorm:"type:varchar(625);not_null"`
}
