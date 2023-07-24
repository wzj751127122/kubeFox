package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Node node

type node struct{}

// GetNodeList 获取Node列表
// ListPage godoc
// @Summary      获取Node列表
// @Description  获取Node列表
// @Tags         Node
// @ID           /api/k8s/node/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/node/list [get]
func (n *node) GetNodeList(ctx *gin.Context) {
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

	data, err := service.Node.GetNodes(params.FilterName, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetNodeDetail 获取Node详情
// ListPage godoc
// @Summary      获取Node详情
// @Description  获取Node详情
// @Tags         Node
// @ID           /api/k8s/node/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "node名称"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/node/detail [get]
func (n *node) GetNodeDetail(ctx *gin.Context) {
	params := new(struct {
		Name string `json:"name" form:"name" comment:"Node名称" binding:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.Node.GetNodeDetail(params.Name)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
