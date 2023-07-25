package dao

import (
	"context"
	"k8s-platform/app/opention"
	"k8s-platform/model"
)

// type Operation struct {
// 	db *gorm.DB
// }

func OperationFind(ctx context.Context, in *model.SysOperationRecord) (*model.SysOperationRecord, error) {
	out := &model.SysOperationRecord{}
	return out, opention.DB.WithContext(ctx).Where(in).Find(&out).Error
}

func OperationSave(ctx context.Context, in *model.SysOperationRecord) error {
	return opention.DB.WithContext(ctx).Create(in).Error
}

func OperationDelete(ctx context.Context, in *model.SysOperationRecord) error {
	return opention.DB.WithContext(ctx).Delete(in).Error
}

func OperationDeleteList(ctx context.Context, in []int) error {
	return opention.DB.WithContext(ctx).Delete(&[]model.SysOperationRecord{}, "id in (?)", in).Error
}

// func NewOperation(db *gorm.DB) *operation {
// 	return &operation{db: db}
// }

func OperationPageList(ctx context.Context, params *model.OperationListInput) ([]*model.SysOperationRecord, int64, error) {
	var total int64 = 0
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	query := opention.DB.WithContext(ctx)
	var list []*model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.Method != "" {
		query = query.Where("method = ?", params.Method)
	}
	if params.Path != "" {
		query = query.Where("path = ?", params.Path)
	}
	if params.Status != 0 {
		query = query.Where("status = ?", params.Status)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
