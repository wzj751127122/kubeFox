package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var PersistentVolumeClaim persistentVolumeClaim

type persistentVolumeClaim struct{}

// DeletePersistentVolumeClaim 删除PersistentVolumeClaim
// ListPage godoc
// @Summary      删除PersistentVolumeClaim
// @Description  删除PersistentVolumeClaim
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "PersistentVolumeClaim名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/persistentvolumeclaim/del [delete]
func (s *persistentVolumeClaim) DeletePersistentVolumeClaim(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"配置卷名称" binding:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.PersistentVolumeClaim.DeletePersistentVolumeClaim(params.Name, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// UpdatePersistentVolumeClaim 更新PersistentVolumeClaim
// ListPage godoc
// @Summary      更新PersistentVolumeClaim
// @Description  更新PersistentVolumeClaim
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/persistentvolumeclaim/update [put]
func (s *persistentVolumeClaim) UpdatePersistentVolumeClaim(ctx *gin.Context) {
	params := new(struct {
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
		Content   string `json:"content" form:"content"  binding:"required" comment:"更新内容"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.PersistentVolumeClaim.UpdatePersistentVolumeClaim(params.Content, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetPersistentVolumeClaimList 查看PersistentVolumeClaim列表
// ListPage godoc
// @Summary      查看PersistentVolumeClaim列表
// @Description  查看PersistentVolumeClaim列表
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": }"
// @Router       /api/k8s/persistentvolumeclaim/list [get]
func (s *persistentVolumeClaim) GetPersistentVolumeClaimList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `json:"filter_name" form:"filter_name" binding:"" comment:"过滤名"`
		NameSpace  string `json:"namespace" form:"namespace" binding:"" comment:"命名空间"`
		Limit      int    `json:"limit" form:"limit" binding:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" binding:"" comment:"页码"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.PersistentVolumeClaim.GetPersistentVolumeClaims(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetPersistentVolumeClaimDetail 获取PersistentVolumeClaim详情
// ListPage godoc
// @Summary      获取PersistentVolumeClaim详情
// @Description  获取PersistentVolumeClaim详情
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "PersistentVolumeClaim名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/persistentvolumeclaim/detail [get]
func (s *persistentVolumeClaim) GetPersistentVolumeClaimDetail(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"配置卷名称" binding:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.PersistentVolumeClaim.GetPersistentVolumeClaimDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
