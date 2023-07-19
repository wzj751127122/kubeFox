package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var IngressController ingressController

type ingressController struct{}

// CreateIngress 创建ingress
// ListPage godoc
// @Summary      创建ingress
// @Description  创建ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.IngressCreteInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "新增成功}"
// @Router       /api/k8s/ingress/create [post]
func (i *ingressController) CreateIngress(ctx *gin.Context) {
	params := new(service.IngressCreate)
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	if err := service.Ingress.CreateIngress(params); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "新增成功")
}

// DeleteIngress 删除ingress
// ListPage godoc
// @Summary      删除ingress
// @Description  删除ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/ingress/del [delete]
func (i *ingressController) DeleteIngress(ctx *gin.Context) {
	params := new(struct{
		Name      string `json:"name" form:"name" comment:"服务名称" validate:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	if err := service.Ingress.DeleteIngress(params.NameSpace, params.Name); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// UpdateIngress 更新ingress
// ListPage godoc
// @Summary      更新ingress
// @Description  更新ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/ingress/update [put]
func (i *ingressController) UpdateIngress(ctx *gin.Context) {
	params := new(struct{
		Content   string `json:"content" form:"content" validate:"required" comment:"更新内容"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	if err := service.Ingress.UpdateIngress(params.NameSpace, params.Content); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetIngressList 查看ingress列表
// ListPage godoc
// @Summary      查看ingress列表
// @Description  查看ingress列表
// @Tags         ingress
// @ID           /api/k8s/ingress/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":""  }"
// @Router       /api/k8s/ingress/list [get]
func (i *ingressController) GetIngressList(ctx *gin.Context) {
	params := new(struct{
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

	data, err := service.Ingress.GetIngressList(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetIngressDetail 获取ingress详情
// ListPage godoc
// @Summary      获取ingress详情
// @Description  获取ingress详情
// @Tags         ingress
// @ID           /api/k8s/ingress/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":""  }"
// @Router       /api/k8s/ingress/detail [get]
func (i *ingressController) GetIngressDetail(ctx *gin.Context) {
	params := new(struct{
		Name      string `json:"name" form:"name" comment:"服务名称" validate:"required"`
		NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}

	data, err := service.Ingress.GetIngressDetail(params.NameSpace, params.Name)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetIngressNumPreNp 根据命名空间获取ingress数量
// ListPage godoc
// @Summary      根据命名空间获取ingress数量
// @Description  根据命名空间获取ingress数量
// @Tags         ingress
// @ID           /api/k8s/ingress/numnp
// @Accept       json
// @Produce      json
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/ingress/numnp [get]
func (i *ingressController) GetIngressNumPreNp(ctx *gin.Context) {
	data, err := service.Ingress.GetIngressNp()
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
