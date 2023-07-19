package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var ServiceController serviceController

type serviceController struct{}

// CreateService 创建service
// ListPage godoc
// @Summary      创建service
// @Description  创建service
// @Tags         service
// @ID           /api/k8s/service/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.ServiceCreateInput  true  "body"
// @Success       200  {object}  middle.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/service/create [post]
func (s *serviceController) CreateService(ctx *gin.Context) {
	params := new(service.ServiceCreate)
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.Service.CreateService(params); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "创建成功")
}

// DeleteService 删除service
// ListPage godoc
// @Summary      删除service
// @Description  删除service
// @Tags         service
// @ID           /api/k8s/service/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/service/del [delete]
func (s *serviceController) DeleteService(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"服务名称" validate:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.Service.DeleteService(params.Name, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// UpdateService 更新service
// ListPage godoc
// @Summary      更新service
// @Description  更新service
// @Tags         service
// @ID           /api/k8s/service/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/service/update [put]
func (s *serviceController) UpdateService(ctx *gin.Context) {
	params := new(struct {
		Content   string `json:"content" validate:"required" comment:"更新内容"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	if err := service.Service.UpdateService(params.NameSpace, params.Content); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "修改成功")
}

// GetServiceList 查看service列表
// ListPage godoc
// @Summary      查看service列表
// @Description  查看service列表
// @Tags         service
// @ID           /api/k8s/service/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/service/list [get]
func (s *serviceController) GetServiceList(ctx *gin.Context) {
	params := new(struct {
		FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
		NameSpace  string `json:"namespace" form:"namespace" validate:"" comment:"命名空间"`
		Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.Service.GetServiceList(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetServiceDetail 获取service详情
// ListPage godoc
// @Summary      获取service详情
// @Description  获取service详情
// @Tags         service
// @ID           /api/k8s/service/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/service/detail [get]
func (s *serviceController) GetServiceDetail(ctx *gin.Context) {
	params := new(struct {
		Name      string `json:"name" form:"name" comment:"服务名称" validate:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	data, err := service.Service.GetServiceDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetServicePerNS 根据命名空间获取service数量
// ListPage godoc
// @Summary      根据命名空间获取service数量
// @Description  根据命名空间获取service数量
// @Tags         service
// @ID           /api/k8s/service/numnp
// @Accept       json
// @Produce      json
// @Success       200  {object}  middle.Response"{"code": 200, msg="","data":service.serviceNp }"
// @Router       /api/k8s/service/numnp [get]
func (s *serviceController) GetServicePerNS(ctx *gin.Context) {
	data, err := service.Service.GetServiceNp()

	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
