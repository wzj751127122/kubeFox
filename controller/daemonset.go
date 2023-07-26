package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var DaemonSet daemonSet

type daemonSet struct{}

// DeleteDaemonSet 删除DaemonSet
// ListPage godoc
// @Summary      删除DaemonSet
// @Description  删除DaemonSet
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "DaemonSet名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/daemonset/del [delete]
func (s *daemonSet) DeleteDaemonSet(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"有状态控制器名称" binding:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
	})

	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	err = service.DaemonSet.DeleteDaemonSet(params.Name, params.NameSpace)
	if err != nil {

		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除daemonset成功")
}

// UpdateDaemonSet 更新DaemonSet
// ListPage godoc
// @Summary      更新DaemonSet
// @Description  更新DaemonSet
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/daemonset/update [put]
func (s *daemonSet) UpdateDaemonSet(ctx *gin.Context) {
	params := new(struct {
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
		Content   string `json:"content" binding:"required" comment:"更新内容"`
	})
	err := ctx.ShouldBindJSON(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	err = service.DaemonSet.UpdateDaemonSet(params.Content, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetDaemonSetList 查看DaemonSet列表
// ListPage godoc
// @Summary      查看DaemonSet列表
// @Description  查看DaemonSet列表
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": }"
// @Router       /api/k8s/daemonset/list [get]
func (s *daemonSet) GetDaemonSetList(ctx *gin.Context) {
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

	data, err := service.DaemonSet.GetDaemonSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return

	}
	middle.ResponseSuccess(ctx, data)
}

// GetDaemonSetDetail 获取DaemonSet详情
// ListPage godoc
// @Summary      获取DaemonSet详情
// @Description  获取DaemonSet详情
// @Tags         DaemonSet
// @ID           /api/k8s/DaemonSet/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "DaemonSet名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/daemonset/detail [get]
func (s *daemonSet) GetDaemonSetDetail(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"有状态控制器名称" binding:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	data, err := service.DaemonSet.GetDaemonSetDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
