package model

// CasbinInfo Casbin info structure
type CasbinInfo struct {
	Path   string `form:"path"  json:"path"`      // 路径
	Method string ` form:"method"  json:"method"` // 方法
}

// UpdateCasbinInput 通过角色id更改接口权限
type UpdateCasbinInput struct {
	AuthorityId uint         `form:"authorityId" json:"authorityId"` // 权限id
	CasbinInfo  []CasbinInfo `json:"casbinInfos"`
}

// CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint `form:"authorityId" json:"authorityId"` // 权限id
}
