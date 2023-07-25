package dao

import (
	"context"
	"k8s-platform/app/opention"
	"k8s-platform/model"
)

func AuthorityFind(ctx context.Context, authInfo *model.SysAuthority) (*model.SysAuthority, error) {
	var out *model.SysAuthority
	return out, opention.DB.WithContext(ctx).Where(authInfo).Find(out).Error
}

func AuthorityFindList(ctx context.Context, authInfo *model.SysAuthority) ([]*model.SysAuthority, error) {
	var out []*model.SysAuthority
	return out, opention.DB.WithContext(ctx).Where(&authInfo).Find(&out).Error
}

func AuthorityPageList(ctx context.Context, params model.PageInfo) ([]model.SysAuthority, int64, error) {
	var total int64 = 0
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	query := opention.DB.WithContext(ctx)
	var list []model.SysAuthority
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.Keyword != "" {
		query = query.Where("authority_name = ?", params.Keyword)
	}
	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("authority_id desc").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func AuthoritySave(ctx context.Context, authInfo *model.SysAuthority) error {
	return opention.DB.WithContext(ctx).Create(authInfo).Error
}

func AuthorityUpdates(ctx context.Context, authInfo *model.SysAuthority) error {
	return opention.DB.WithContext(ctx).Updates(authInfo).Error
}

// SetMenuAuthority 菜单与角色绑定
func AuthoritySetMenuAuthority(ctx context.Context, authInfo *model.SysAuthority) error {
	var s model.SysAuthority
	opention.DB.WithContext(ctx).Preload("SysBaseMenus").First(&s, "authority_id = ?", authInfo.AuthorityId)
	return opention.DB.WithContext(ctx).Model(&s).Association("SysBaseMenus").Replace(&authInfo.SysBaseMenus)
}
