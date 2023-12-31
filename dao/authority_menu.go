package dao

import (
	"context"
	"k8s-platform/app/opention"
	"k8s-platform/model"
)

func AuthorityMenuFindList(ctx context.Context, in *model.SysAuthorityMenu) ([]*model.SysAuthorityMenu, error) {
	var out []*model.SysAuthorityMenu
	return out, opention.DB.WithContext(ctx).Where(&in).Find(&out).Error
}
