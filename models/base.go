package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `gorm:"primaryKey" json:"id,string"`
	CreatedAt time.Time      `gorm:"autoCreateTime"` // 自动填充创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // 自动填充更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`          // 软删除支持
}
