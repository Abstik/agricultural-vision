package models

import (
	"time"
)

type User struct {
	Id          int64     `gorm:"primaryKey" json:"-"`
	Username    string    `gorm:"type:varchar(64);not null" json:"username"`
	Email       string    `gorm:"type:varchar(64);not null;unique" json:"email"`
	Password    string    `gorm:"type:varchar(64);not null" json:"-"`
	CreatedTime time.Time `gorm:"type:datetime" json:"-"`
	UpdatedTime time.Time `gorm:"type:datetime" json:"-"`
}
