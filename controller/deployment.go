package controller

import (

	"k8s-platform/middle"
	"k8s-platform/service"


	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Deployment deployment

type deployment struct{}

// 获取deployment列表支持分页，过滤
func (d *deployment) GetDeployments(c *gin.Context) {

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

	data, err := service.Deployment.GetDeployment(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {

		logger.Error("获取deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}
	middle.ResponseSuccess(c,data)
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "获取deployment列表成功",
	// 	"data": data,
	// })

}

// 获取deployemnt详情
func (d *deployment) GetDeploymentsDetail(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		DeploymentName string `form:"deployment_name"`
		Namespace      string `form:"namespace"`
	})

	err := c.ShouldBind(params)
	if err != nil {

		logger.Error("bind失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	data, err := service.Deployment.GetDeploymentDetail(params.DeploymentName, params.Namespace)
	if err != nil {


		logger.Error("获取deployment详情失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

middle.ResponseSuccess(c,data)

}

// 创建deployment
func (d *deployment) CreateDeployment(c *gin.Context) {

	var (
		deployCreate = new(service.DeployCreate)
		err          error
	)

	err = c.ShouldBindJSON(deployCreate)
	if err != nil {

		logger.Error("bind失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Deployment.CreateDeployment(deployCreate)
	if err != nil {


		logger.Error("创建deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

middle.ResponseSuccess(c,nil)

}

// 设置deployment副本数

func (d *deployment) ScaleDeployment(c *gin.Context) {

	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
		ScaleNum       int    `json:"scale_num"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {

		logger.Error("bind失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	data, err := service.Deployment.ScaleDeployment(params.DeploymentName, params.Namespace, params.ScaleNum)

	if err != nil {


		logger.Error("扩容deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "设置deployment副本数成功",
	// 	"data": fmt.Sprintf("最新副本数为 %d", data),
	// })
	middle.ResponseSuccess(c,data)

}

//删除deployment

func (d *deployment) DeleteDeployment(c *gin.Context) {

	//处理传入的变量

	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {

		logger.Error("bind失败" + err.Error())
		middle.ResponseError(c, middle.CodeInvalidParam)
		return
	}

	err = service.Deployment.DeleteDeployment(params.DeploymentName, params.Namespace)
	if err != nil {


		logger.Error("删除deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "删除deployment列表成功",
	// 	"data": nil,
	// })

	middle.ResponseSuccess(c,nil)
}

// 重启deployment
func (d *deployment) RestartDeployment(c *gin.Context) {

	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {

		logger.Error("bind失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	err = service.Deployment.RestartDeployment(params.DeploymentName, params.Namespace)
	if err != nil {


		logger.Error("重启deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "重启deployment成功",
	// 	"data": nil,
	// })
middle.ResponseSuccess(c,nil)
}

// 更新deployment
func (d *deployment) UpdateDeployment(c *gin.Context) {
	params := new(struct {
		content   string `json:"content"`
		namespace string `json:"namespace"`
	})

	err := c.ShouldBindJSON(params)
	if err != nil {
		logger.Error("Bind绑定参数失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}

	err = service.Deployment.UpdateDeployment(params.content, params.namespace)
	if err != nil {

		logger.Error("更新deployment失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "更新deployment成功",
	// 	"data": nil,
	// })
	middle.ResponseSuccess(c,nil)
}

func (d *deployment) GetDeploymentNumPreNS(c *gin.Context) {
	data, err := service.Deployment.GetDeploymentNumNp()
	if err != nil {

		logger.Error("获取deployment数量失败" + err.Error())
		middle.ResponseError(c, middle.CodeServerBusy)
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg":  "更新deployment成功",
	// 	"data": data,
	// })
	middle.ResponseSuccess(c,data)

}
