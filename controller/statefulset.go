package controller

import (
	"k8s-platform/middle"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var StatefulSet statefulSet

type statefulSet struct{}

// DeleteStatefulSet 删除statefulSet
// ListPage godoc
// @Summary      删除statefulSet
// @Description  删除statefulSet
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "statefulSet名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middle.ResponseData "{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/statefulset/del [delete]
func (s *statefulSet) DeleteStatefulSet(ctx *gin.Context) {
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

	err = service.StatefulSet.DeleteStatefulSet(params.Name, params.NameSpace)
	if err != nil {

		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "删除daemonset成功")
}

// UpdateStatefulSet 更新statefulSet
// ListPage godoc
// @Summary      更新statefulSet
// @Description  更新statefulSet
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/statefulset/update [put]
func (s *statefulSet) UpdateStatefulSet(ctx *gin.Context) {
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

	err = service.StatefulSet.UpdateStatefulSet(params.Content, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, "更新成功")
}

// GetStatefulSetList 查看statefulSet列表
// ListPage godoc
// @Summary      查看statefulSet列表
// @Description  查看statefulSet列表
// @Tags         statefulSet
// @ID           /api/k8s/statefulset/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middle.ResponseData"{"code": 200, msg="","data": }"
// @Router       /api/k8s/statefulset/list [get]
func (s *statefulSet) GetStatefulSetList(ctx *gin.Context) {
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

	data, err := service.StatefulSet.GetStatefulSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return

	}
	middle.ResponseSuccess(ctx, data)
}

// GetStatefulSetDetail 获取statefulSet详情
// ListPage godoc
// @Summary      获取statefulSet详情
// @Description  获取statefulSet详情
// @Tags         statefulSet
// @ID           /api/k8s/statefulSet/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "statefulSet名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middle.ResponseData"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/statefulset/detail [get]
func (s *statefulSet) GetStatefulSetDetail(ctx *gin.Context) {
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

	data, err := service.StatefulSet.GetStatefulSetDetail(params.Name, params.NameSpace)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(ctx, data)
}
