package controller

import (
	"k8s-platform/middle"

	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Configmap configmap

type configmap struct{}

// DeleteConfigmap 删除Configmap
// ListPage godoc
// @Summary      删除Configmap
// @Description  删除Configmap
// @Tags         Configmap
// @ID           /api/k8s/configmap/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Configmap名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/configmap/del [delete]
func (s *configmap) DeleteConfigmap(ctx *gin.Context) {
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
	if err := service.Configmap.DeleteConfigmap(params.Name, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// UpdateConfigmap 更新Configmap
// ListPage godoc
// @Summary      更新Configmap
// @Description  更新Configmap
// @Tags         Configmap
// @ID           /api/k8s/configmap/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/configmap/update [put]
func (s *configmap) UpdateConfigmap(ctx *gin.Context) {
	params := new(struct {
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" binding:"required"`
		Content   string `json:"content" form:"content" binding:"required" comment:"更新内容"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.Configmap.UpdateConfigmap(params.Content, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetConfigmapList 查看Configmap列表
// ListPage godoc
// @Summary      查看Configmap列表
// @Description  查看Configmap列表
// @Tags         Configmap
// @ID           /api/k8s/configmap/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": }"
// @Router       /api/k8s/configmap/list [get]
func (s *configmap) GetConfigmapList(ctx *gin.Context) {
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
	data, err := service.Configmap.GetConfigmaps(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetConfigmapDetail 获取Configmap详情
// ListPage godoc
// @Summary      获取Configmap详情
// @Description  获取Configmap详情
// @Tags         Configmap
// @ID           /api/k8s/configmap/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Configmap名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/configmap/detail [get]
func (s *configmap) GetConfigmapDetail(ctx *gin.Context) {
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
	data, err := service.Configmap.GetConfigmapDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
