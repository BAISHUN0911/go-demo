package request

// *uint 表示可选，如果没传就是 nil
type SysAuthoritiesReq struct {
	TenantID *uint `json:"tenantId"` // 可选参数
}
