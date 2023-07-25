package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AddSysMenusInput struct {
	ParentId string `json:"parentId" comment:"父菜单ID" binding:"required"` // 父菜单ID
	Name     string `json:"name"  comment:"路由name" binding:"required"`   // 路由name
	Path     string `json:"path" comment:"路由path" binding:"required"`    // 路由path
	Disabled bool   `json:"disabled" comment:"是否禁用" binding:"required"`  // 是否在列表隐藏
	Hidden   bool   `json:"hidden" comment:"是否在列表隐藏" binding:"required"` // 是否在列表隐藏
	Sort     int    `json:"sort" comment:"排序标记" binding:"required"`      // 排序标记
	Meta
}

type SysMenusResponse struct {
	Menus []SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu SysBaseMenu `json:"menu"`
}

type AddMenuAuthorityInput struct {
	Menus       []SysBaseMenu `json:"menus"`
	AuthorityId uint          `json:"authorityId"  binding:"required"` // 角色ID
}

type SysBaseMenu struct {
	ID            int                                        `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	MenuLevel     uint                                       `json:"-"`
	ParentId      string                                     `json:"parentId" gorm:"comment:父菜单ID"`    // 父菜单ID
	Path          string                                     `json:"path" gorm:"comment:路由path"`       // 路由path
	Name          string                                     `json:"name" gorm:"comment:路由name"`       // 路由name
	Hidden        bool                                       `json:"hidden" gorm:"comment:是否在列表隐藏"`    // 是否在列表隐藏
	Disabled      bool                                       `json:"disabled" gorm:"comment:是否禁止修改菜单"` // 是否在列表隐藏
	Sort          int                                        `json:"sort" gorm:"comment:排序标记"`         // 排序标记
	Children      []SysBaseMenu                              `json:"children" gorm:"-"`
	Meta          `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	SysAuthoritys []SysAuthority                             `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	CreatedAt     time.Time                                  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time                                  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt                             `gorm:"index" json:"-"`
}

type Meta struct {
	ActiveName string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive  bool   `json:"keepAlive" gorm:"comment:是否缓存"`   // 是否缓存
	Title      string `json:"title" gorm:"comment:菜单名"`        // 菜单名
	Icon       string `json:"icon" gorm:"comment:菜单图标"`        // 菜单图标
	CloseTab   bool   `json:"closeTab" gorm:"comment:自动关闭tab"` // 自动关闭tab
}

func init() {
	RegisterInitializer(SysBaseMenuOrder, &SysBaseMenu{})
}

func (m *SysBaseMenu) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&m)
}

func (m *SysBaseMenu) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysBaseMenu
	if err := db.WithContext(ctx).Where("path = 'dashboard' ").Find(&out).Error; err != nil {
		return false, nil
	}
	return out.ID != 0, nil
}

func (m *SysBaseMenu) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := m.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	if !ok {
		if err := db.WithContext(ctx).Create(&SysBaseMenuEntities).Error; err != nil {
			menu := SysBaseMenu{}
			return errors.Wrap(err, menu.TableName()+"表数据初始化失败!")
		}
	}
	return nil
}

func (m *SysBaseMenu) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&m)
}

func (*SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
