package models

import (
	"time"
)

type User struct {
	UserID      int64  `gorm:"primaryKey"`
	Username    string `gorm:"not null;uniqueIndex;type:varchar(64)"`
	Email       string `gorm:"not null;type:varchar(64)"`
	Password    string `gorm:"not null;type:varchar(64)"`
	CreatedTime time.Time
	UpdatedTime time.Time
}
