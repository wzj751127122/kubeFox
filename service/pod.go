package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"io"



	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Pod pod

type pod struct{}

// 定义列表的返回内容，items是pod的元素列表，total是元素的数量
type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}


type PodsNp struct{
	Namespace string  	`json:"namespace"`

	PodNum int			`json:"pod_num"`
}

// 获取pod列表，支持排序过滤分页
func (p *pod) GetPods(filterName, namespace string, limit, page int) (podResp *PodsResp, err error) {

	podList, err := K8s.clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取pod列表失败" + err.Error())
		return nil, errors.New("获取pod列表失败" + err.Error())
	}

	// 实例化dataselector结构体，组装数据

	selectableData := &dataSelector{
		GenericDatalist: p.toCells(podList.Items),
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

	pods := p.fromCells(data.GenericDatalist)

	return &PodsResp{
		Total: total,
		Items: pods,
	}, nil
}

// 获取pod详情
func (p *pod) GetDetail(podName, namespace string) (pod *corev1.Pod, err error) {

	pod, err = K8s.clientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取Pod详情失败" + err.Error())
		return nil, errors.New("获取Pod详情失败" + err.Error())
	}
	return

}

// 删除pod

func (p *pod) DeletePod(podName, namespace string) (err error) {

	err = K8s.clientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除Pod失败" + err.Error())
		return errors.New("删除Pod失败" + err.Error())
	}
	return

}

// 更新Pod

func (p *pod) UpdatePod( namespace, content string) (err error) {

	var pod = &corev1.Pod{}

	err = json.Unmarshal([]byte(content), pod)
	if err != nil {
		logger.Error("反序列化失败" + err.Error())
		return errors.New("反序列化失败" + err.Error())
	}

	_, err = K8s.clientSet.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新Pod失败" + err.Error())
		return errors.New("更新Pod失败" + err.Error())
	}
	return
}

// 获取Pod容器名

func (p *pod) GetPodContainer(podName, namespace string) (containers []string, err error) {

	pod, err := p.GetDetail(podName, namespace)
	if err != nil {
		return nil, err
	}

	for _, container := range pod.Spec.Containers {

		containers = append(containers, container.Name)

	}
	return
}

// 获取pod内容器的日志

func (p *pod) GetPodLog(containerName, podName, namespace string) (log string, err error) {


	limit := int64(200)

	// 创建日志请求
	req := K8s.clientSet.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		Container: containerName,
		Follow:    false,
		TailLines: &limit,
	})

	// 发送日志请求
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logger.Error("获取PodLog失败" + err.Error())
		return "", errors.New("获取PodLog失败" + err.Error())
	}
	defer podLogs.Close()

	// 读取日志内容
	buf := new(bytes.Buffer)
	_,err = io.Copy(buf,podLogs)
	if err != nil {
		logger.Error("拷贝PodLog失败" + err.Error())
		return "", errors.New("拷贝PodLog失败" + err.Error())
	}

	return buf.String(),nil



}


// 获取每个namespace的pod数量
func (p *pod)GetPodNum()(podsNps []*PodsNp,err error)  {

	namespaceList,err := K8s.clientSet.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		logger.Error("获取namespace列表失败" + err.Error())
		return nil, errors.New("获取namespace列表失败" + err.Error())
	}

	for _,namespace :=range namespaceList.Items{
		podList,err := K8s.clientSet.CoreV1().Pods(namespace.Name).List(context.TODO(),metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
	

	podsNp := &PodsNp{
		Namespace : namespace.Name,
		PodNum : len(podList.Items),
	}

	podsNps = append(podsNps,podsNp)
	}

return podsNps,nil
}

// 定义类型转换的方法

func (p *pod) toCells(pods []corev1.Pod) []DataCell {

	cells := make([]DataCell, len(pods))
	for i := range pods {
		cells[i] = podcell(pods[i])
	}
	return cells
}
func (p *pod) fromCells(cells []DataCell) []corev1.Pod {

	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(podcell))
	}
	return pods
}
