package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Deployment deployment

type deployment struct{}

// 定义列表的返回内容，items是pod的元素列表，total是元素的数量
type DeploymentResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

type DeploymentNp struct {
	Namespace     string `json:"namespace"`
	DeploymentNum int    `json:"deployment_num"`
}

// 定义deploymnetcreat结构体，用于创建deploymen属性

type DeployCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Image         string            `json:"image"`
	Cpu           string            `json:"cpu"`
	Memery        string            `json:"memery"`
	HealthPath    string            `json:"health_path"`
	Replicas      int32             `json:"replicas"`
	ContainerPort int32             `json:"container_port"`
	Label         map[string]string `json:"label"`
	HealthCheck   bool              `json:"health_check"`
}

// 获取deployment列表

func (d *deployment) GetDeployment(filterName, namespace string, limit, page int) (deploymentResp *DeploymentResp, err error) {

	deploymentList, err := K8s.clientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取deployment列表失败" + err.Error())
		return nil, errors.New("获取deployment列表失败" + err.Error())
	}

	// 实例化dataselector结构体，组装数据

	selectableData := &dataSelector{
		GenericDatalist: d.toCells(deploymentList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDatalist)

	data := filtered.Sort().Paginate()

	pods := d.fromCells(data.GenericDatalist)

	return &DeploymentResp{
		Total: total,
		Items: pods,
	}, nil
}

// 获取deployment详情

func (d *deployment) GetDeploymentDetail(deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {

	deployment, err = K8s.clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取deployment详情失败" + err.Error())
		return nil, errors.New("获取deployment详情失败" + err.Error())
	}
	return

}

// 修改deployment副本数

func (d *deployment) ScaleDeployment(deploymentName, namespace string, scaleNum int) (replica int32, err error) {

	scale, err := K8s.clientSet.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取deployment副本数失败" + err.Error())
		return 0, errors.New("获取deployment副本数失败" + err.Error())
	}

	scale.Spec.Replicas = int32(scaleNum)

	newScale, err := K8s.clientSet.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新deployment副本数失败" + err.Error())
		return 0, errors.New("更新deployment副本数失败" + err.Error())
	}

	return newScale.Spec.Replicas, nil

}

// 创建deployment
func (d *deployment) CreateDeployment(data *DeployCreate) (err error) {
	// 初始化一个appsv1的deployment类型的对象，并将data入参传入

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.Label,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          data.Name,
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: data.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
		Status: appsv1.DeploymentStatus{},
	}

	// 判断健康检查功能是否打卡，如果打开则增加健康检查功能

	if data.HealthCheck {
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			InitialDelaySeconds: 5,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}

		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			InitialDelaySeconds: 15,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}
	}

	// 定义容器的limit和request资源
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits = map[corev1.ResourceName]resource.Quantity{
		corev1.ResourceCPU:    resource.MustParse(data.Cpu),
		corev1.ResourceMemory: resource.MustParse(data.Memery),
	}

	deployment.Spec.Template.Spec.Containers[0].Resources.Requests = map[corev1.ResourceName]resource.Quantity{
		corev1.ResourceCPU:    resource.MustParse(data.Cpu),
		corev1.ResourceMemory: resource.MustParse(data.Memery),
	}

	// 调用sdk创建deployment
	_, err = K8s.clientSet.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		logger.Error("创建deployment失败" + err.Error())
		return errors.New("创建deployment失败" + err.Error())
	}
	return
}

// 删除deployment

func (d *deployment) DeleteDeployment(deploymentName, namespace string) (err error) {

	err = K8s.clientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除deployment失败" + err.Error())
		return errors.New("删除deployment失败" + err.Error())
	}
	return

}

// 更新deployment
func (d *deployment) UpdateDeployment(content, namespace string) (err error) {

	var deployment = &appsv1.Deployment{}

	err = json.Unmarshal([]byte(content), deployment)
	if err != nil {
		logger.Error("反序列化deployment失败" + err.Error())
		return errors.New("反序列化deployment失败" + err.Error())
	}

	_, err = K8s.clientSet.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新Deployment失败" + err.Error())
		return errors.New("更新Deployment失败" + err.Error())
	}
	return

}

// 重启deployment
func (d *deployment) RestartDeployment(deploymentName, namespace string) (err error) {

	deploymentsClient := K8s.clientSet.AppsV1().Deployments(namespace)

	data := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))
	// deployment, err := deploymentsClient.Patch(ctx, deployment_name, k8stypes.StrategicMergePatchType, []byte(data), v1.PatchOptions{})

	_, err = deploymentsClient.Patch(context.TODO(), deploymentName, "application/strategic-merge-patch+json", []byte(data), metav1.PatchOptions{})
	if err != nil {
		logger.Error("重启Deployment失败" + err.Error())
		return errors.New("重启Deployment失败" + err.Error())
	}
	return
}

// 获取每个namespace的deployment数量
func (d *deployment) GetDeploymentNumNp() (deploysNps []*DeploymentNp, err error) {
	namespaceList, err := K8s.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("获取namespace列表失败" + err.Error())
		return nil, errors.New("获取namespace列表失败" + err.Error())
	}

	for _, namespace := range namespaceList.Items {
		deployList, err := K8s.clientSet.AppsV1().Deployments(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		deploysNp := &DeploymentNp{
			Namespace:     namespace.Name,
			DeploymentNum: len(deployList.Items),
		}

		deploysNps = append(deploysNps, deploysNp)
	}

	return deploysNps, nil
}

// 定义类型转换的方法

func (d *deployment) toCells(deployment []appsv1.Deployment) []DataCell {

	cells := make([]DataCell, len(deployment))
	for i := range deployment {
		cells[i] = deploymentCell(deployment[i])
	}
	return cells
}
func (d *deployment) fromCells(cells []DataCell) []appsv1.Deployment {

	deployment := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		deployment[i] = appsv1.Deployment(cells[i].(deploymentCell))
	}
	return deployment
}
