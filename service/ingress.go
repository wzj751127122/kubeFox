package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/wonderivan/logger"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Ingress ingress


type ingress struct{}

type IngressCreate struct{
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace"`
	Label     map[string]string      `json:"label"`
	Hosts     map[string][]*HttpPath `json:"hosts"`
}

type  HttpPath struct{

	Path string
	PathType    nwv1.PathType
	ServiceName string
	ServicePort int32
}

func (i *ingress)CreateIngress(data *IngressCreate) (err error) {

	//声明变量用于组装数据
	var ingressRules []nwv1.IngressRule
	var httpIngressPATHs  []nwv1.HTTPIngressPath


	//将data中的数据组装成ingress对象

		ingress := &nwv1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name: data.Name,
				Namespace: data.Namespace,
				Labels: data.Label,
			},
			Status: nwv1.IngressStatus{},
		}

		// 第一层for是将host组装成ingressRule
		for k,v := range data.Hosts{
			ir := nwv1.IngressRule{
				Host: k,
				IngressRuleValue: nwv1.IngressRuleValue{
					HTTP: &nwv1.HTTPIngressRuleValue{Paths: nil},    //这里将path值为空后面在赋值
				},
			}

			//// 第二层for是将path组装成ingresspath
			for _,httpPath := range v {
				hip := nwv1.HTTPIngressPath{
					Path: httpPath.Path,
					PathType: &httpPath.PathType,
					Backend: nwv1.IngressBackend{
						Service: &nwv1.IngressServiceBackend{
							Name: httpPath.ServiceName,
							Port: nwv1.ServiceBackendPort{
								Number: httpPath.ServicePort,
							},
						},
					},
				}

				httpIngressPATHs = append(httpIngressPATHs, hip)
			}

			// 此处赋值，前面值为空了

			ir.IngressRuleValue.HTTP.Paths = httpIngressPATHs

			// 将每个ir对象组装成数组，这个ir就是ingressrule每个元素是一个host和多个path
			ingressRules = append(ingressRules, ir)
		}
		ingress.Spec.Rules = ingressRules

		_,err = K8s.clientSet.NetworkingV1().Ingresses(data.Namespace).Create(context.TODO(),ingress,metav1.CreateOptions{})
		if err != nil {
			logger.Error("创建ingress失败"+err.Error())
			return errors.New("创建ingress失败"+err.Error())
		}

		return


	
}

func (i *ingress)DeleteIngress(ingressName,namespace string) error {
	
	fmt.Println("ingress")
	return nil

}