package dao

import (
	"context"
	"k8s-platform/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)
var (
	DB  *gorm.DB
)

// type User interface {
// 	Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error)
// 	Save(ctx context.Context, userInfo *model.SysUser) error
// 	Updates(ctx context.Context, userInfo *model.SysUser) error
// 	Delete(ctx context.Context, userInfo *model.SysUser) error
// }

// var _ User = &user{}
// var User user
// type user struct {
// }

func Find(ctx context.Context, userInfo *model.SysUser) (*model.SysUser, error) {
	user := &model.SysUser{}
	if err := DB.WithContext(ctx).Preload("Authorities").Preload("Authority").Where(userInfo).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func Save(ctx context.Context, userInfo *model.SysUser) error {
	return DB.WithContext(ctx).Save(userInfo).Error
}

func Updates(ctx context.Context, userInfo *model.SysUser) error {
	if userInfo.ID == 0 {
		return errors.New("id not set")
	}
	return DB.WithContext(ctx).Updates(userInfo).Error
}

func Delete(ctx context.Context, userInfo *model.SysUser) error {
	return DB.WithContext(ctx).Delete(userInfo).Error
}
