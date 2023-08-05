package operation

import (
	"fmt"
	"k8s-platform/logic"
	"k8s-platform/middle"
	"k8s-platform/model"
	"k8s-platform/utils"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Operation operation

type operation struct{}

// GetOperationRecordList
// @Tags      SysOperationRecord
// @Summary   分页获取SysOperationRecord列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     model.OperationListInput                      true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  middle.ResponseError(ctx, middle.CodeServerBusy)  "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/operation/get_operations [get]
func (o *operation) GetOperationRecordList(ctx *gin.Context) {

	params := &model.OperationListInput{}
	err := ctx.ShouldBind(params)
	if err != nil {
		zap.L().Error("参数绑定失败")
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	// 解码 path 参数
	params.Path, err = url.QueryUnescape(params.Path)
	if err != nil {
		zap.L().Error("解码 path 参数失败")
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	fmt.Println(params.Path)
	data, err := logic.GetPageList(ctx, params)
	if err != nil {
		zap.L().Error("查询失败")
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// DeleteOperationRecord
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     model.Empty
// @Success   200   {object}  middle.ResponseSuccess{msg=string}  "删除SysOperationRecord"
// @Router    /api/operation/{id}/delete_operation [delete]
func (o *operation) DeleteOperationRecord(ctx *gin.Context) {
	recordId, err := utils.ParseInt(ctx.Param("id"))
	if err != nil {
		zap.L().Error("参数绑定失败")
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	err = logic.DeleteRecord(ctx, recordId)
	if err != nil {
		zap.L().Error("删除失败")
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")

}

// DeleteOperationRecords
// @Tags      SysOperationRecord
// @Summary   删除SysOperationRecord
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     model.IdsReq
// @Success   200   {object}  middle.ResponseSuccess{msg=string}  "删除SysOperationRecord"
// @Router    /api/operation/delete_operations [delete]
func (o *operation) DeleteOperationRecords(ctx *gin.Context) {
	params := &model.IdsReq{}
	err := ctx.ShouldBind(params)
	if err != nil {
		zap.L().Error("参数绑定失败")
		middle.ResponseError(ctx, middle.CodeInvalidParam)
		return
	}
	if err := logic.DeleteRecords(ctx, params.Ids); err != nil {
		zap.L().Error("批量删除失败")
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}
