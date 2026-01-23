package model

// 角色
type Authority struct {
	// BaseModel
	AuthorityID   uint         `gorm:"primaryKey" json:"authority_id"`
	AuthorityName string       `json:"authority_name"`
	ParentID      uint         `json:"parent_id"`
	Children      []*Authority `gorm:"-" json:"children"` // 用于存放子角色
	TenantID      uint        `json:"tenant_id"`
}

func (Authority) TableName() string {
	return "sys_authorities"
}
