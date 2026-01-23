package model

// AllModels 返回所有需要迁移的模型
func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Authority{},
		// 在这里添加其他模型
	}
}
