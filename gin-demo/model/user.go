package model

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Role     string `gorm:"column:role" json:"role"`
	Active   bool   `gorm:"column:active" json:"active"`
}

func (User) TableName() string {
	return "users"
}
