package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wonderivan/logger"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Ingress ingress

type ingress struct{}

type ingressResp struct {
	Total int            `json:"total"`
	Items []nwv1.Ingress `json:"items"`
}

type ingressNp struct {
	NameSpace  string `json:"namespace"`
	IngressNum int    `json:"ingres_num"`
}

type IngressCreate struct {
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace"`
	Label     map[string]string      `json:"label"`
	Hosts     map[string][]*HttpPath `json:"hosts"`
}

type HttpPath struct {
	Path        string
	PathType    nwv1.PathType
	ServiceName string
	ServicePort int32
}

func (i *ingress) toCells(ingress []nwv1.Ingress) []DataCell {
	cells := make([]DataCell, len(ingress))
	for i := range ingress {
		cells[i] = ingressCell(ingress[i])
	}
	return cells
}

func (i *ingress) FromCells(cells []DataCell) []nwv1.Ingress {
	ingress := make([]nwv1.Ingress, len(cells))
	for i := range cells {
		ingress[i] = nwv1.Ingress(cells[i].(ingressCell))
	}
	return ingress
}

func (i *ingress) CreateIngress(data *IngressCreate) (err error) {

	//声明变量用于组装数据
	var ingressRules []nwv1.IngressRule
	var httpIngressPATHs []nwv1.HTTPIngressPath

	//将data中的数据组装成ingress对象

	ingress := &nwv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Status: nwv1.IngressStatus{},
	}

	// 第一层for是将host组装成ingressRule
	for k, v := range data.Hosts {
		ir := nwv1.IngressRule{
			Host: k,
			IngressRuleValue: nwv1.IngressRuleValue{
				HTTP: &nwv1.HTTPIngressRuleValue{Paths: nil}, //这里将path值为空后面在赋值
			},
		}

		//// 第二层for是将path组装成ingresspath
		for _, httpPath := range v {
			hip := nwv1.HTTPIngressPath{
				Path:     httpPath.Path,
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

	_, err = K8s.clientSet.NetworkingV1().Ingresses(data.Namespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		logger.Error("创建ingress失败" + err.Error())
		return errors.New("创建ingress失败" + err.Error())
	}

	return

}

func (i *ingress) DeleteIngress(namespace, name string) error {

	err := K8s.clientSet.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除ingress失败" + err.Error())
		return errors.New("删除ingress失败" + err.Error())
	}

	return nil

}

func (i *ingress) UpdateIngress(namespace, content string) error {
	ingress := &nwv1.Ingress{}
	if err := json.Unmarshal([]byte(content), ingress); err != nil {
		logger.Error("反序列化deployment失败" + err.Error())
		return errors.New("反序列化deployment失败" + err.Error())
	}
	_, err := K8s.clientSet.NetworkingV1().Ingresses(namespace).Update(context.TODO(), ingress, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新ingress失败" + err.Error())
		return errors.New("更新ingress失败" + err.Error())
	}
	return nil
}

func (i *ingress) GetIngressList(filterName, namespace string, limit, page int) (*ingressResp, error) {
	ingressList, err := K8s.clientSet.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	fmt.Println(ingressList)
	if err != nil {
		logger.Error("获取ingress失败" + err.Error())
		return nil, errors.New("获取ingress失败" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDatalist: i.toCells(ingressList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filterd := selectableData.Filter()
	total := len(filterd.GenericDatalist)
	data := filterd.Sort().Paginate()
	ingress := i.FromCells(data.GenericDatalist)
	return &ingressResp{
		Total: total,
		Items: ingress,
	}, nil
}

func (i *ingress) GetIngressDetail(namespace, name string) (*nwv1.Ingress, error) {
	data, err := K8s.clientSet.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取ingress详情失败" + err.Error())
		return nil, errors.New("获取ingress详情失败" + err.Error())
	}
	return data, nil
}

func (i *ingress) GetIngressNp() ([]*ingressNp, error) {
	namespaceList, err := K8s.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var ingressnps []*ingressNp
	for _, namespace := range namespaceList.Items {
		ingress, err := K8s.clientSet.NetworkingV1().Ingresses(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		ingressNp := &ingressNp{
			NameSpace:  namespace.Name,
			IngressNum: len(ingress.Items),
		}
		ingressnps = append(ingressnps, ingressNp)
	}
	return ingressnps, err
}
