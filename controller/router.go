package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化router的对象，用于跨包调用，首字母大写
var Router router

type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {

	router.GET("/testapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"data": nil,
		})
	})
	

	router.GET("/api/k8s/pods", Pod.GetPods)

	v1 := router.Group("/api/k8s/pod")

	{

		v1.GET("/detail", Pod.GetPodsDetail)
		v1.DELETE("/del", Pod.DeletePod)
		v1.PUT("/update", Pod.UpdatePod)
		v1.GET("/container", Pod.GetPodContainer)
		v1.GET("/log", Pod.GetPodLog)
		v1.GET("/numnp", Pod.GetPodNumPerNp)

	}

}
