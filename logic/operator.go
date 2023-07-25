package logic

import (
	"k8s-platform/dao"
	"k8s-platform/model"

	"github.com/gin-gonic/gin"
)

func CreateOperationRecord(ctx *gin.Context, record *model.SysOperationRecord) error {
	return dao.OperationSave(ctx, record)
}

func DeleteRecord(ctx *gin.Context, id int) error {
	record := &model.SysOperationRecord{ID: id}
	return dao.OperationDelete(ctx, record)
}

func DeleteRecords(ctx *gin.Context, ids []int) error {
	return dao.OperationDeleteList(ctx, ids)
}

func GetPageList(ctx *gin.Context, in *model.OperationListInput) (*model.OperationListOutPut, error) {
	list, total, err := dao.OperationPageList(ctx, in)
	if err != nil {
		return nil, err
	}
	return &model.OperationListOutPut{OperationList: list, Total: total, PageInfo: in.PageInfo}, nil
}
