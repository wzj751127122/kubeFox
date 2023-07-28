package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Service service

type service struct{}

// 创建service

type serviceResp struct {
	Total int              `json:"total"`
	Items []coreV1.Service `json:"items"`
}

type serviceNp struct {
	NameSpace  string `json:"namespace"`
	ServiceNum int    `json:"service_num"`
}

type ServiceCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`
}

func (s *service) toCells(services []coreV1.Service) []DataCell {
	cells := make([]DataCell, len(services))
	for i := range services {
		cells[i] = serviceCell(services[i])
	}
	return cells
}

func (s *service) FromCells(cells []DataCell) []coreV1.Service {
	services := make([]coreV1.Service, len(cells))
	for i := range cells {
		services[i] = coreV1.Service(cells[i].(serviceCell))
	}
	return services
}

func (s *service) CreateService(data *ServiceCreate) (err error) {

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(data.Type),
			Ports: []corev1.ServicePort{
				{
					Name:     data.Name,
					Port:     data.Port,
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
		Status: coreV1.ServiceStatus{},
	}
	//默认为clusterIP，此处判断NodePort
	if data.NodePort != 0 && data.Type == "NodePort" {
		service.Spec.Ports[0].NodePort = data.NodePort
	}

	_, err = K8s.clientSet.CoreV1().Services(data.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		logger.Error("创建service失败" + err.Error())
		return errors.New("创建service失败" + err.Error())
	}
	return

}

func (s *service) DeleteService(name, namespace string) error {
	err := K8s.clientSet.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除daemonset失败" + err.Error())
		return errors.New("删除daemonset失败" + err.Error())
	}
	return nil
}

func (s *service) UpdateService(namespace, content string) error {
	var Service = &coreV1.Service{}
	if err := json.Unmarshal([]byte(content), Service); err != nil {
		logger.Error("反序列化service失败" + err.Error())
		return errors.New("反序列化service失败" + err.Error())
	}
	_, err := K8s.clientSet.CoreV1().Services(namespace).Update(context.TODO(), Service, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新service失败" + err.Error())
		return errors.New("更新service失败" + err.Error())
	}
	return nil
}

func (s *service) GetServiceList(filterName, namespace string, limit, page int) (*serviceResp, error) {
	ServiceList, err := K8s.clientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("获取service失败" + err.Error())
		return nil, errors.New("获取service失败" + err.Error())
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDatalist: s.toCells(ServiceList.Items),
		DataSelect: &DataSelectQuery{
			Filter:   &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{limit, page},
		},
	}
	//先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDatalist)
	//排序、分页
	data := filtered.Sort().Paginate()
	//将dataCell类型转换为coreV1.Pod
	Services := s.FromCells(data.GenericDatalist)
	return &serviceResp{
		total,
		Services,
	}, nil
}

func (s *service) GetServiceDetail(name, namespace string) (*coreV1.Service, error) {
	data, err := K8s.clientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取service详情失败" + err.Error())
		return nil, errors.New("获取service详情失败" + err.Error())
	}
	return data, nil
}

func (s *service) GetServiceNp() ([]*serviceNp, error) {
	namespaceList, err := K8s.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var services []*serviceNp
	for _, namespace := range namespaceList.Items {
		serviceList, err := K8s.clientSet.CoreV1().Services(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		ServiceNp := &serviceNp{
			NameSpace:  namespace.Name,
			ServiceNum: len(serviceList.Items),
		}
		services = append(services, ServiceNp)
	}
	return services, nil
}
