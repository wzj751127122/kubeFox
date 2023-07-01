package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Service service

type service struct{}


// 创建service

type ServiceCreate struct{
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`


}


func (s *service)CreateService(data *ServiceCreate)(err error)  {

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: data.Name,
			Namespace: data.Namespace,
			Labels: data.Label,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(data.Type),
			Ports: []corev1.ServicePort{
				{
					Name: data.Name,
					Port: data.Port,
					Protocol: "tcp",
					TargetPort: intstr.IntOrString{
						Type: 0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
	}
	//默认为clusterIP，此处判断NodePort
	if data.NodePort != 0 &&data.Type == "NodePort"{
		service.Spec.Ports[0].NodePort = data.NodePort
	}

	_ ,err =K8s.clientSet.CoreV1().Services(data.Namespace).Create(context.TODO(),service,metav1.CreateOptions{})
	if err != nil {
		logger.Error("创建service失败"+err.Error())
		return errors.New("创建service失败"+err.Error())
	}
	return


	
}

func (s *service)DeleteService(serviceName, namespace string) error {
	fmt.Println("service")
	return nil
}