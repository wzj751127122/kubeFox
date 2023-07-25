package controller

import (


	"github.com/gin-gonic/gin"
)

// 初始化router的对象，用于跨包调用，首字母大写


func KubeApiRouter(router *gin.RouterGroup) {

	// router.GET("/testapi", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg":  "ok",
	// 		"data": nil,
	// 	})
	// })
	

	// router.GET("/api/k8s/pods", Pod.GetPods)

	v1 := router.Group("/k8s")

	{
		v1.POST("/deployment/create", Deployment.CreateDeployment)
		v1.DELETE("/deployment/del", Deployment.DeleteDeployment)
		v1.PUT("/deployment/update", Deployment.UpdateDeployment)
		v1.GET("/deployment/list", Deployment.GetDeployments)
		v1.GET("/deployment/detail", Deployment.GetDeploymentsDetail)
		v1.PUT("/deployment/restart", Deployment.RestartDeployment)
		v1.GET("/deployment/scale", Deployment.ScaleDeployment)
		v1.GET("/deployment/numnp", Deployment.GetDeploymentNumPreNS)
	}
	{
		v1.GET("/pod/list", Pod.GetPods)
		v1.GET("/pod/detail", Pod.GetPodsDetail)
		v1.DELETE("/pod/del", Pod.DeletePod)
		v1.PUT("/pod/update", Pod.UpdatePod)
		v1.GET("/pod/container", Pod.GetPodContainer)
		v1.GET("/pod/log", Pod.GetPodLog)
		v1.GET("/pod/numnp", Pod.GetPodNumPerNp)
		v1.GET("/pod/webshell", Pod.WebShell)
	}
	{
		v1.DELETE("/daemonset/del", DaemonSet.DeleteDaemonSet)
		v1.PUT("/daemonset/update", DaemonSet.UpdateDaemonSet)
		v1.GET("/daemonset/list", DaemonSet.GetDaemonSetList)
		v1.GET("/daemonset/detail", DaemonSet.GetDaemonSetDetail)
	}
	{
		v1.DELETE("/statefulset/del", StatefulSet.DeleteStatefulSet)
		v1.PUT("/statefulset/update", StatefulSet.UpdateStatefulSet)
		v1.GET("/statefulset/list", StatefulSet.GetStatefulSetList)
		v1.GET("/statefulset/detail", StatefulSet.GetStatefulSetDetail)
	}
	{
		v1.GET("/node/list", Node.GetNodeList)
		v1.GET("/node/detail", Node.GetNodeDetail)
	}

	{
		v1.PUT("/namespace/create", NameSpace.CreateNameSpace)
		v1.DELETE("/namespace/del", NameSpace.DeleteNameSpace)
		v1.GET("/namespace/list", NameSpace.GetNameSpaceList)
		v1.GET("/namespace/detail", NameSpace.GetNameSpaceDetail)
	}

	{
		v1.DELETE("/persistentvolume/del", PersistentVolume.DeletePersistentVolume)
		v1.GET("/persistentvolume/list", PersistentVolume.GetPersistentVolumeList)
		v1.GET("/persistentvolume/detail", PersistentVolume.GetPersistentVolumeDetail)
	}

	{
		v1.POST("/service/create", ServiceController.CreateService)
		v1.DELETE("/service/del", ServiceController.DeleteService)
		v1.PUT("/service/update", ServiceController.UpdateService)
		v1.GET("/service/list", ServiceController.GetServiceList)
		v1.GET("/service/detail", ServiceController.GetServiceDetail)
		v1.GET("/service/numnp", ServiceController.GetServicePerNS)
	}

	{
		v1.PUT("/ingress/create", IngressController.CreateIngress)
		v1.DELETE("/ingress/del", IngressController.DeleteIngress)
		v1.PUT("/ingress/update", IngressController.UpdateIngress)
		v1.GET("/ingress/list", IngressController.GetIngressList)
		v1.GET("/ingress/detail", IngressController.GetIngressDetail)
		v1.GET("/ingress/numnp", IngressController.GetIngressNumPreNp)
	}

	{
		v1.DELETE("/configmap/del", Configmap.DeleteConfigmap)
		v1.PUT("/configmap/update", Configmap.UpdateConfigmap)
		v1.GET("/configmap/list", Configmap.GetConfigmapList)
		v1.GET("/configmap/detail", Configmap.GetConfigmapDetail)
	}

	{
		v1.DELETE("/persistentvolumeclaim/del", PersistentVolumeClaim.DeletePersistentVolumeClaim)
		v1.PUT("/persistentvolumeclaim/update", PersistentVolumeClaim.UpdatePersistentVolumeClaim)
		v1.GET("/persistentvolumeclaim/list", PersistentVolumeClaim.GetPersistentVolumeClaimList)
		v1.GET("/persistentvolumeclaim/detail", PersistentVolumeClaim.GetPersistentVolumeClaimDetail)
	}

	{
		v1.DELETE("/secret/del", Secret.DeleteSecret)
		v1.PUT("/secret/update", Secret.UpdateSecret)
		v1.GET("/secret/list", Secret.GetSecretList)
		v1.GET("/secret/detail", Secret.GetSecretDetail)
	}

	{
		v1.POST("/workflow/create", WorkFlow.CreateWorkFlow)
		v1.DELETE("/workflow/del", WorkFlow.DeleteWorkflow)
		v1.GET("/workflow/list", WorkFlow.GetWorkflowList)
		v1.GET("/workflow/id", WorkFlow.GetWorkflowByID)
	}


}
