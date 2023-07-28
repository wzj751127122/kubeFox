package controller

import (
	"k8s-platform/middle"
	"k8s-platform/model"
	// "k8s-platform/model"
	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Pod pod

type pod struct{}

// 获取pod
func (p *pod) GetPods(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// FilterName string `form:"filtername"`
		// Namespace  string `form:"namespace"`
		// Limit      int    `form:"limit"`
		// Page       int    `form:"page"`
		FilterName string `json:"filter_name" form:"filter_name" binding:"" comment:"过滤名"`
		NameSpace  string `json:"namespace" form:"namespace" binding:"" comment:"命名空间"`
		Limit      int    `json:"limit" form:"limit" binding:"" comment:"分页限制"`
		Page       int    `json:"page" form:"page" binding:"" comment:"页码"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetPods(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod列表成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c, data)

}

// 获取pod详情
func (p *pod) GetPodsDetail(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// PodName   string `form:"pod_name"`
		// Namespace string `form:"namespace"`
		PodName   string `json:"pod_name" form:"pod_name" comment:"POD名称" binding:"required"`
		NameSpace string `json:"name_space" form:"namespace" comment:"命名空间" binding:"required"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetDetail(params.PodName, params.NameSpace)
	if err != nil {

		logger.Error("获取pod详情失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod详情成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c, data)

}

// 删除pod

func (p *pod) DeletePod(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// PodName   string `json:"pod_name"`
		// Namespace string `json:"namespace"`
		PodName   string `json:"pod_name" form:"pod_name" comment:"POD名称" binding:"required"`
		NameSpace string `json:"name_space" form:"namespace" comment:"命名空间" binding:"required"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Pod.DeletePod(params.PodName, params.NameSpace)
	if err != nil {

		logger.Error("删除pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "删除pod列表成功",
	// 	"data": nil,
	// })
	middle.ResponseSuccess(c, nil)

}

//获取每个namespace的pod数量

func (p *pod) GetPodNumPerNp(c *gin.Context) {

	data, err := service.Pod.GetPodNum()
	if err != nil {

		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取每个namespace中pod数量成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c, data)

}

// 获取pod中容器的日志
func (p *pod) GetPodLog(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// PodName       string `form:"pod_name"`
		// Namespace     string `form:"namespace"`
		// ContainerName string `form:"container_name"`
		PodName       string `json:"pod_name" form:"pod_name" comment:"POD名称" binding:"required"`
		NameSpace     string `json:"name_space" form:"namespace" comment:"命名空间" binding:"required"`
		ContainerName string `json:"container_name" form:"container_name" comment:"容器名称" binding:"required"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	log, err := service.Pod.GetPodLog(params.ContainerName, params.PodName, params.NameSpace)
	if err != nil {

		logger.Error("获取pod日志失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod中容器日志成功",
	// 	"data": log,
	// })
	middle.ResponseSuccess(c, log)

}

// 获取pod容器
func (p *pod) GetPodContainer(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// PodName   string `form:"pod_name"`
		// Namespace string `form:"namespace"`
		PodName   string `json:"pod_name" form:"pod_name" comment:"POD名称" binding:"required"`
		NameSpace string `json:"name_space" form:"namespace" comment:"命名空间" binding:"required"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetPodContainer(params.PodName, params.NameSpace)
	if err != nil {

		logger.Error("获取pod容器失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod容器成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c, data)

}

// 更新pod      put请求
func (p *pod) UpdatePod(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		// Namespace string `json:"namespace"`
		// Content   string `json:"content"`
		PodName   string `json:"pod_name" form:"pod_name" comment:"POD名称" binding:"required"`
		NameSpace string `json:"name_space" form:"namespace" comment:"命名空间" binding:"required"`
		Content   string `json:"content" form:"content" comment:"内容" binding:"required"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Pod.UpdatePod(params.NameSpace, params.Content)
	if err != nil {

		logger.Error("更新pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "更新pod成功",
	// 	"data": nil,
	// })
	middle.ResponseSuccess(c, nil)

}

func (p *pod) WebShell(ctx *gin.Context) {
	params := new(model.WebShellStr)
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	err = service.Terminal.WsHandler(params, ctx.Writer, ctx.Request)
	if err != nil {
		middle.ResponseError(ctx, middle.CodeServerBusy)
	}
	// if err != nil {
	// 	middle.ResponseError(ctx, middle.CodeServerBusy)
	// }
}
