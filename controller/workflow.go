package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var WorkFlow workflow

type workflow struct{}

// CreateWorkFlow 创建workflow
// ListPage godoc
// @Summary      创建workflow
// @Description  创建workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.WorkFlowCreateInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/workflow/create [post]
func (w *workflow) CreateWorkFlow(ctx *gin.Context) {
	params := new(service.WorkflowCreate)
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	err = service.Workflow.CreateWorkflow(params)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "创建成功")
}

// DeleteWorkflow 删除Workflow
// ListPage godoc
// @Summary      删除Workflow
// @Description  删除Workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/del
// @Accept       json
// @Produce      json
// @Param        ID       query  int  true  "Workflow ID"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/workflow/del [delete]
func (w *workflow) DeleteWorkflow(ctx *gin.Context) {
	params := new(struct {
		ID int `json:"id" form:"id"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.Workflow.DelById(params.ID); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// GetWorkflowList 查看Configmap列表
// ListPage godoc
// @Summary      查看Configmap列表
// @Description  查看Configmap列表
// @Tags         Workflow
// @ID           /api/k8s/workflow/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/workflow/list [get]
func (w *workflow) GetWorkflowList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `json:"filter_name" form:"filter_name" binding:"" comment:"过滤名"`
		Limit      int    `json:"limit" form:"limit" binding:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" binding:"" comment:"页码"`
		Namespace  string `json:"namespace" form:"namespace"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.Workflow.GetList(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetWorkflowByID 根据ID查看workflow
// ListPage godoc
// @Summary      根据ID查看workflow
// @Description  根据ID查看workflow
// @Tags         Workflow
// @ID           /api/k8s/workflow/id
// @Accept       json
// @Produce      json
// @Param        ID       query  int  true  "Workflow ID"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/workflow/id [get]
func (w *workflow) GetWorkflowByID(ctx *gin.Context) {
	params := new(struct {
		ID int `json:"id" form:"id"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.Workflow.GetListById(params.ID)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
