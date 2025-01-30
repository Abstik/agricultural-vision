package models

import (
	"time"
)

type User struct {
	UserID      int64     `gorm:"primaryKey" json:"user_id"`
	Username    string    `gorm:"not null;type:varchar(64)" json:"username"`
	Email       string    `gorm:"not null;unique;type:varchar(64)" json:"email"`
	Password    string    `gorm:"not null;type:varchar(64)" json:"password"`
	CreatedTime time.Time `gorm:"type:datetime" json:"created_time"`
	UpdatedTime time.Time `gorm:"type:datetime" json:"updated_time"`
}
