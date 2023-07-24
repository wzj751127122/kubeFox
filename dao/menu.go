package dao

import (
	"context"
	"k8s-platform/model"
)

func MenuFind(ctx context.Context, in *model.SysBaseMenu) (*model.SysBaseMenu, error) {
	var out *model.SysBaseMenu
	return out, DB.WithContext(ctx).Where(in).Find(&out).Error
}

func MenuFindIn(ctx context.Context, in []string) ([]*model.SysBaseMenu, error) {
	//做一下排序
	var out []*model.SysBaseMenu
	return out, DB.WithContext(ctx).Where("id in (?)", in).Order("sort").Find(&out).Error
}

func MenuFindList(ctx context.Context, in *model.SysBaseMenu) ([]model.SysBaseMenu, error) {
	var out []model.SysBaseMenu
	return out, DB.WithContext(ctx).Order("sort").Where(in).Find(&out).Error
	
}


func MenuSave(ctx context.Context, in *model.SysBaseMenu) error {
	return DB.WithContext(ctx).Create(in).Error
}

func MenuUpdates(ctx context.Context, in *model.SysBaseMenu) error {
	return DB.WithContext(ctx).Updates(in).Error
}
