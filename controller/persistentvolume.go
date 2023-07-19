package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var PersistentVolume persistentVolume

type persistentVolume struct{}

// DeletePersistentVolume 删除persistentVolume
// ListPage godoc
// @Summary      删除persistentVolume
// @Description  删除persistentVolume
// @Tags         PersistentVolume
// @ID           /api/k8s/spersistentvolume/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "persistentvolume名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/spersistentvolume/del [delete]
func (n *persistentVolume) DeletePersistentVolume(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"命名空间名称" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.PersistentVolume.DeletePersistentVolume(params.Name); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// GetPersistentVolumeList 获取PV列表
// ListPage godoc
// @Summary      获取PV列表
// @Description  获取PV列表
// @Tags         PersistentVolume
// @ID           /api/k8s/persistentvolume/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.PersistentVolumeResp}"
// @Router       /api/k8s/persistentvolume/list [get]
func (n *persistentVolume) GetPersistentVolumeList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
		Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetPersistentVolumeDetail 获取PV的详细信息
// ListPage godoc
// @Summary      获取PV的详细信息
// @Description  获取PV的详细信息
// @Tags         PersistentVolume
// @ID           /api/k8s/persistentvolume/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "persistentVolume名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": *coreV1.PersistentVolume}"
// @Router       /api/k8s/persistentvolume/detail [get]
func (n *persistentVolume) GetPersistentVolumeDetail(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"命名空间名称" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.PersistentVolume.GetPersistentVolumesDetail(params.Name)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
