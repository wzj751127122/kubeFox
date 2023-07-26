package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Secret secret

type secret struct{}

// DeleteSecret 删除Secret
// ListPage godoc
// @Summary      删除Secret
// @Description  删除Secret
// @Tags         Secret
// @ID           /api/k8s/Secret/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Secret名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/Secret/del [delete]
func (s *secret) DeleteSecret(ctx *gin.Context) {
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
	if err := service.Secret.DeleteSecrets(params.Name, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除成功")
}

// UpdateSecret 更新Secret
// ListPage godoc
// @Summary      更新Secret
// @Description  更新Secret
// @Tags         Secret
// @ID           /api/k8s/secret/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/secret/update [put]
func (s *secret) UpdateSecret(ctx *gin.Context) {
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
	if err := service.Secret.UpdateSecrets(params.Content, params.NameSpace); err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetSecretList 查看Secret列表
// ListPage godoc
// @Summary      查看Secret列表
// @Description  查看Secret列表
// @Tags         Secret
// @ID           /api/k8s/Secret/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": }"
// @Router       /api/k8s/Secret/list [get]
func (s *secret) GetSecretList(ctx *gin.Context) {
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
	data, err := service.Secret.GetSecrets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}

// GetSecretDetail 获取Secret详情
// ListPage godoc
// @Summary      获取Secret详情
// @Description  获取Secret详情
// @Tags         Secret
// @ID           /api/k8s/Secret/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Secret名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/Secret/detail [get]
func (s *secret) GetSecretDetail(ctx *gin.Context) {
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
	data, err := service.Secret.GetSecretsDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
