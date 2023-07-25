package controller

import (
	"k8s-platform/middle"
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
		FilterName string `form:"filtername"`
		Namespace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetPods(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod列表成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c,data)

}

// 获取pod详情
func (p *pod) GetPodsDetail(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
	})

	err := c.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetDetail(params.PodName, params.Namespace)
	if err != nil {

		logger.Error("获取pod详情失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod详情成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c,data)

}

// 删除pod

func (p *pod) DeletePod(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Pod.DeletePod(params.PodName, params.Namespace)
	if err != nil {

		logger.Error("删除pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "删除pod列表成功",
	// 	"data": nil,
	// })
	middle.ResponseSuccess(c,nil)

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
	middle.ResponseSuccess(c,data)

}

// 获取pod中容器的日志
func (p *pod) GetPodLog(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		PodName       string `form:"pod_name"`
		Namespace     string `form:"namespace"`
		ContainerName string `form:"container_name"`
	})

	err := c.Bind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	log, err := service.Pod.GetPodLog(params.ContainerName, params.PodName, params.Namespace)
	if err != nil {

		logger.Error("获取pod日志失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod中容器日志成功",
	// 	"data": log,
	// })
	middle.ResponseSuccess(c,log)

}

// 获取pod容器
func (p *pod) GetPodContainer(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
	})

	err := c.Bind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	data, err := service.Pod.GetPodContainer(params.PodName, params.Namespace)
	if err != nil {

		logger.Error("获取pod容器失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取pod容器成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c,data)

}

// 更新pod      put请求
func (p *pod) UpdatePod(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Pod.UpdatePod(params.Namespace, params.Content)
	if err != nil {

		logger.Error("更新pod失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "更新pod成功",
	// 	"data": nil,
	// })
	middle.ResponseSuccess(c,nil)

}

func (p *pod) WebShell(ctx *gin.Context) {
	params := new(struct {
		Namespace string `form:"namespace"`
		Pod       string `form:"pod_name"`
		Container string `form:"container_name"`
	})
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(ctx, middle.CodeServerBusy)
		return
	}
	service.Terminal.WsHandler(ctx.Writer,ctx.Request)
	// if err != nil {
	// 	middle.ResponseError(ctx, middle.CodeServerBusy)
	// }
}
