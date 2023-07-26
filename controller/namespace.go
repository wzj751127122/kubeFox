package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var NameSpace namespace

type namespace struct{}

// CreateNameSpace 创建namespace
// ListPage godoc
// @Summary      创建namespace
// @Description  创建namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/create
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/namespace/create [put]
func (n *namespace) CreateNameSpace(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"命名空间名称" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	if err := service.NameSpace.CreateNameSpace(params.Name); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "创建成功")
}

// DeleteNameSpace 删除namespace
// ListPage godoc
// @Summary      删除namespace
// @Description  删除namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/namespace/del [delete]
func (n *namespace) DeleteNameSpace(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"命名空间名称" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.NameSpace.DeleteNameSpace(params.Name); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// GetNameSpaceList 获取NameSpace列表
// ListPage godoc
// @Summary      获取NameSpace列表
// @Description  获取NameSpace列表
// @Tags         NameSpace
// @ID           /api/k8s/namespace/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/namespace/list [get]
func (n *namespace) GetNameSpaceList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `json:"filter_name" form:"filter_name" binding:"" comment:"过滤名"`
		Limit      int    `json:"limit" form:"limit" binding:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" binding:"" comment:"页码"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.NameSpace.GetNameSpaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetNameSpaceDetail 获取NameSpace详情
// ListPage godoc
// @Summary      获取NameSpace详情
// @Description  获取NameSpace详情
// @Tags         NameSpace
// @ID           /api/k8s/namespace/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/namespace/detail [get]
func (n *namespace) GetNameSpaceDetail(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"命名空间名称" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.NameSpace.GetNameSpacesDetail(params.Name)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
