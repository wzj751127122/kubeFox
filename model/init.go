package model

import (
	"database/sql"
	"strconv"

	adapter "github.com/casbin/gorm-adapter/v3"

	uuid "github.com/satori/go.uuid"
)

// 初始化顺序
const (
	SysUserOrder = iota
	MenuAuthorityOrder
	SysBaseMenuOrder
	SysAuthorityOrder
	SysApisInitOrder
	CasbinInitOrder
	OperatorationOrder
	WorkFlowOrder
)

var (
	AdminDefaultAuth      uint = 111
	AdminDefaultAuthStr        = strconv.Itoa(int(AdminDefaultAuth))
	UserDefaultAuth       uint = 222
	UserDefaultAuthStr         = strconv.Itoa(int(UserDefaultAuth))
	UserSubDefaultAuth    uint = 2221
	UserSubDefaultAuthStr      = strconv.Itoa(int(UserSubDefaultAuth))
)
const (
	LoginURL    = "/api/user/login"
	LogoutURL   = "/api/user/logout"
	WebShellURL = "/api/k8s/pod/webshell"
)

// SysUserEntities 用户初始化数据
var (
	SysUserEntities = []*SysUser{
		{
			UUID:        uuid.NewV4(),
			UserName:    "admin",
			Password:    "$2a$14$EhdHTX14t7uGT/Xo8FM/cusNrPKtNE1Ktf7mhbPc2my5Qupts26T2",
			NickName:    "admin",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: AdminDefaultAuth,
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
		{
			UUID:        uuid.NewV4(),
			UserName:    "mylove",
			Password:    "$2a$14$Nav0R8zHMMXrGjcl8fNXv.aW968ihONyUsGBtuNcIpSAuSut.MUKe",
			NickName:    "zw",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: AdminDefaultAuth,
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
	}
)

// SysBaseMenuEntities 菜单初始化数据
var (
	SysBaseMenuEntities = []SysBaseMenu{
		// 根菜单
		{MenuLevel: 0, Hidden: false, Disabled: true, ParentId: "0", Path: "dashboard", Name: "仪表盘", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "cmdb", Name: "资产中心", Sort: 3, Meta: Meta{Title: "资产中心", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "kubernetes", Name: "容器管理", Sort: 4, Meta: Meta{Title: "容器管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "devops", Name: "应用发布", Sort: 5, Meta: Meta{Title: "应用发布", Icon: "compass"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "setting", Name: "系统设置", Sort: 6, Meta: Meta{Title: "系统设置", Icon: "setting"}},
		//子菜单 ParentId对应跟菜单顺序 且不需要icon
		// 资产中心子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "2", Path: "host", Name: "主机管理", Sort: 0, Meta: Meta{Title: "主机管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "2", Path: "secret", Name: "授权管理", Sort: 1, Meta: Meta{Title: "授权管理"}},
		// 容器管理子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "cluster", Name: "集群管理", Sort: 0, Meta: Meta{Title: "集群管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "deployment", Name: "工作负载", Sort: 1, Meta: Meta{Title: "工作负载"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "service", Name: "服务发现", Sort: 2, Meta: Meta{Title: "服务发现"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "node", Name: "节点管理", Sort: 3, Meta: Meta{Title: "节点管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "config", Name: "配置中心", Sort: 4, Meta: Meta{Title: "配置中心"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "events", Name: "事件中心", Sort: 5, Meta: Meta{Title: "事件中心"}},
		// 系统设置子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "authority", Name: "角色管理", Sort: 1, Meta: Meta{Title: "角色管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "user", Name: "用户管理", Sort: 2, Meta: Meta{Title: "用户管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "operation", Name: "操作历史", Sort: 3, Meta: Meta{Title: "操作历史"}},
	}
)

// SysAuthorityEntities 角色初始化数据
var (
	SysAuthorityEntities = []SysAuthority{
		{
			AuthorityId:   AdminDefaultAuth,
			AuthorityName: "管理员",
			DefaultRouter: "dashboard",
			ParentId:      0,
		},
		{
			AuthorityId:   UserDefaultAuth,
			AuthorityName: "普通用户",
			DefaultRouter: "dashboard",
			ParentId:      0,
		},
		{
			AuthorityId:   UserSubDefaultAuth,
			AuthorityName: "普通用户子角色",
			DefaultRouter: "dashboard",
			ParentId:      222,
		},
	}
)

var CasbinApi = buildCasbinRule(SysApis)

// buildCasbinRule 构建角色casbin api
func buildCasbinRule(apis []SysApi) []adapter.CasbinRule {
	var out []adapter.CasbinRule
	// 管理员角色添加所有api
	for _, api := range apis {
		rule := adapter.CasbinRule{
			Ptype: "p",
			V0:    AdminDefaultAuthStr,
			V1:    api.Path,
			V2:    api.Method,
		}
		out = append(out, rule)
	}
	otherRule := []adapter.CasbinRule{
		// admin添加所有接口
		{Ptype: "p", V0: UserDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
		{Ptype: "p", V0: UserDefaultAuthStr, V1: "/api/user/loginout", V2: "GET"},
		{Ptype: "p", V0: UserDefaultAuthStr, V1: "/api/user/getinfo", V2: "GET"},
		{Ptype: "p", V0: UserDefaultAuthStr, V1: "/api/user/:id/change_pwd", V2: "POST"},

		{Ptype: "p", V0: UserSubDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
		{Ptype: "p", V0: UserSubDefaultAuthStr, V1: "/api/user/loginout", V2: "GET"},
		{Ptype: "p", V0: UserSubDefaultAuthStr, V1: "/api/user/getinfo", V2: "GET"},
		{Ptype: "p", V0: UserSubDefaultAuthStr, V1: "/api/user/:id/change_pwd", V2: "POST"},
	}
	allRules := append(append(out, otherRule...))
	return allRules
}

var SysApis = []SysApi{
	// api接口
	{Path: "/api/sysApi/getAPiList", Description: "获取系统API列表", ApiGroup: "系统", Method: "GET"},

	// 用户相关接口
	{Path: "/api/user/login", Description: "用户登录", ApiGroup: "用户", Method: "POST"},
	{Path: "/api/user/loginout", Description: "用户退出", ApiGroup: "用户", Method: "GET"},
	{Path: "/api/user/getinfo", Description: "获取用户信息", ApiGroup: "用户", Method: "GET"},
	{Path: "/api/user/:id/set_auth", Description: "设置用户权限", ApiGroup: "用户", Method: "PUT"},
	{Path: "/api/user/:id/delete_user", Description: "删除用户", ApiGroup: "用户", Method: "DELETE"},
	{Path: "/api/user/:id/change_pwd", Description: "修改密码", ApiGroup: "用户", Method: "POST"},
	{Path: "/api/user/:id/reset_pwd", Description: "重置密码", ApiGroup: "用户", Method: "PUT"},
	// 操作审计接口
	{Path: "/api/operation/get_operations", Description: "查询操作记录列表", ApiGroup: "操作审计", Method: "GET"},
	{Path: "/api/operation/:id/delete_operation", Description: "删除单条记录", ApiGroup: "操作审计", Method: "DELETE"},
	{Path: "/api/operation/delete_operations", Description: "批量删除记录", ApiGroup: "操作审计", Method: "POST"},
	// Other
	{Path: "/api/swagger/*any", Description: "swagger文档", ApiGroup: "Other", Method: "GET"},
	// 菜单接口
	{Path: "/api/menu/:authID/getMenuByAuthID", Description: "根据角色获取菜单", ApiGroup: "菜单", Method: "GET"},
	{Path: "/api/menu/getBaseMenuTree", Description: "获取菜单总树", ApiGroup: "菜单", Method: "GET"},
	{Path: "/api/menu/add_base_menu", Description: "添加菜单", ApiGroup: "菜单", Method: "POST"},
	{Path: "/api/menu/add_menu_authority", Description: "添加角色", ApiGroup: "菜单", Method: "POST"},
	// 权限RBAC接口
	{Path: "/api/authority/getPolicyPathByAuthorityId", Description: "获取角色api权限", ApiGroup: "权限", Method: "GET"},
	{Path: "/api/authority/updateCasbinByAuthority", Description: "更改角色api权限", ApiGroup: "用户", Method: "POST"},
	{Path: "/api/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "权限", Method: "GET"},
	// K8S相关接口
	{Path: "/api/k8s/deployment/create", Description: "创建deployment", ApiGroup: "Kubernetes", Method: "POST"},
	{Path: "/api/k8s/deployment/del", Description: "删除deployment", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/deployment/update", Description: "更新deployment", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/deployment/list", Description: "查询deployment列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/deployment/detail", Description: "查询deployment详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/deployment/restart", Description: "重启deployment", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/deployment/scale", Description: "deployment扩缩容", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/deployment/numnp", Description: "查询deployment数量信息", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/list", Description: "查询pod列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/detail", Description: "查询pod详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/del", Description: "删除pod", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/pod/update", Description: "更新pod", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/pod/container", Description: "获取Pod内容器名", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/log", Description: "获取容器日志", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/numnp", Description: "查询pod数量信息", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/pod/webshell", Description: "web终端", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/daemonset/del", Description: "删除daemonset", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/daemonset/update", Description: "更新daemonset", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/daemonset/list", Description: "查询daemonset列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/daemonset/detail", Description: "查询daemonset详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/statefulset/del", Description: "删除statefulset", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/statefulset/update", Description: "更新statefulset", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/statefulset/list", Description: "查询statefulset列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/statefulset/detail", Description: "查询statefulset详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/node/list", Description: "查询node列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/node/detail", Description: "查询node详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/namespace/create", Description: "创建namespace", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/namespace/del", Description: "删除namespace", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/namespace/list", Description: "查询namespace列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/namespace/detail", Description: "查询namespace详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/persistentvolume/del", Description: "删除persistentvolume", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/persistentvolume/list", Description: "查询persistentvolume列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/persistentvolume/detail", Description: "查询persistentvolume详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/service/create", Description: "创建service", ApiGroup: "Kubernetes", Method: "POST"},
	{Path: "/api/k8s/service/del", Description: "删除service", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/service/update", Description: "更新service", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/service/list", Description: "查询service列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/service/detail", Description: "查询service详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/service/numnp", Description: "查询service数量信息", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/ingress/create", Description: "创建ingress", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/ingress/del", Description: "删除ingress", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/ingress/update", Description: "更新ingress", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/ingress/list", Description: "查询ingress列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/ingress/detail", Description: "查询ingress详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/ingress/numnp", Description: "查询ingress数量信息", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/configmap/del", Description: "删除configmap", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/configmap/update", Description: "更新configmap", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/configmap/list", Description: "查询configmap列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/configmap/detail", Description: "查询configmap详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/persistentvolumeclaim/del", Description: "删除persistentvolumeclaim", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/persistentvolumeclaim/update", Description: "更新persistentvolumeclaim", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/persistentvolumeclaim/list", Description: "查询persistentvolumeclaim列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/persistentvolumeclaim/detail", Description: "查询persistentvolumeclaim详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/secret/del", Description: "删除secret", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/secret/update", Description: "更新secret", ApiGroup: "Kubernetes", Method: "PUT"},
	{Path: "/api/k8s/secret/list", Description: "查询secret列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/secret/detail", Description: "查询secret详情", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/workflow/create", Description: "创建workflow", ApiGroup: "Kubernetes", Method: "POST"},
	{Path: "/api/k8s/workflow/del", Description: "删除workflow", ApiGroup: "Kubernetes", Method: "DELETE"},
	{Path: "/api/k8s/workflow/list", Description: "查询workflow列表", ApiGroup: "Kubernetes", Method: "GET"},
	{Path: "/api/k8s/workflow/id", Description: "查看workflow", ApiGroup: "Kubernetes", Method: "GET"},
}
