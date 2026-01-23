package model

import (
	"time"
)

type User struct {
	BaseModel
	Username  string     `gorm:"column:username" json:"username"`
	Role      string     `gorm:"column:role" json:"role"`
	Active    bool       `gorm:"column:active" json:"active"`
	LastLogin *time.Time `gorm:"column:last_login" json:"last_login"`
}

func (User) TableName() string {
	return "users"
}
