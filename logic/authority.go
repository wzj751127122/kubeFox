package logic

import (
	"context"
	"k8s-platform/dao"
	"k8s-platform/model"
)

func SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error {
	return dao.AuthoritySetMenuAuthority(ctx, auth)
}

func GetAuthorityList(ctx context.Context, pageInfo model.PageInfo) (*model.AuthorityList, error) {
	list, total, err := dao.AuthorityPageList(ctx, pageInfo)
	if err != nil {
		return nil, err
	}
	return &model.AuthorityList{
		PageInfo:          pageInfo,
		Total:             total,
		AuthorityListItem: list,
	}, nil
}
