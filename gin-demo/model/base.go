package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 通用基础模型，包含常用字段
type BaseModel struct {
	ID        uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
