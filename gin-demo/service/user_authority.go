package service

import (
	"gin-demo/db"
	"gin-demo/model"
)

const SystemRootRoleID = 1

func GetRoleTree(tenantID *uint) ([]*model.Authority, error) {
	// 用于存储所有角色
	var roles []model.Authority

	query := db.DB.Model(&model.Authority{})
	if tenantID != nil {
		query = query.Where("tenant_id = ?", *tenantID)
	}
	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}

	// 构建 map 方便查找子角色
	roleMap := make(map[uint]*model.Authority)
	for i := range roles {
		role := &roles[i]
		role.Children = []*model.Authority{} // 初始化子角色切片
		roleMap[role.AuthorityID] = role
	}
	var tree []*model.Authority

	// 构建树结构
	for _, role := range roleMap {
		if tenantID != nil {
			if role.ParentID == SystemRootRoleID {
				tree = append(tree, role)
			}
		} else {
			if role.ParentID == 0 {
				tree = append(tree, role)
			}
		}

		if parent, ok := roleMap[role.ParentID]; ok {
			parent.Children = append(parent.Children, role)
		}
	}

	return tree, nil
}
