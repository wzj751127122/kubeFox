package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StatefulSet statefulSet

type statefulSet struct{}

type statefulSetResp struct {
	Total int                  `json:"total"`
	Items []appsV1.StatefulSet `json:"items"`
}

type StatefulSetNp struct {
	NameSpace    string `json:"namespace"`
	DaemonSetNum int    `json:"daemonset_num"`
}

func (d *statefulSet) toCells(statefulSets []appsV1.StatefulSet) []DataCell {
	cells := make([]DataCell, len(statefulSets))
	for i := range statefulSets {
		cells[i] = statefulSetCell(statefulSets[i])
	}
	return cells
}

func (d *statefulSet) FromCells(cells []DataCell) []appsV1.StatefulSet {
	statefulSets := make([]appsV1.StatefulSet, len(cells))
	for i := range cells {
		statefulSets[i] = appsV1.StatefulSet(cells[i].(statefulSetCell))
	}
	return statefulSets
}

func (d *statefulSet) GetStatefulSets(filterName, namespace string, limit, page int) (*statefulSetResp, error) {
	statefulSetList, err := K8s.clientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Info("获取StatefulSet列表失败" + err.Error())
		return nil, errors.New("获取StatefulSet列表失败" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDatalist: d.toCells(statefulSetList.Items),
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
	statefulSets := d.FromCells(data.GenericDatalist)
	return &statefulSetResp{
		Total: total,
		Items: statefulSets,
	}, nil
}

func (d *statefulSet) GetStatefulSetDetail(name, namespace string) (*appsV1.StatefulSet, error) {
	data, err := K8s.clientSet.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		logger.Info("获取StatefulSet详情失败" + err.Error())
		return nil, errors.New("获取StatefulSet详情失败" + err.Error())
	}
	return data, nil
}

func (d *statefulSet) DeleteStatefulSet(name, namespace string) (err error) {

	err = K8s.clientSet.AppsV1().StatefulSets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logger.Error("删除StatefulSet失败" + err.Error())
		return errors.New("删除StatefulSet失败" + err.Error())
	}
	return

}

func (d *statefulSet) UpdateStatefulSet(content, namespace string) (err error) {
	var statefulSet = &appsV1.StatefulSet{}
	err = json.Unmarshal([]byte(content), statefulSet)
	if err != nil {
		if err != nil {
			logger.Error("反序列化daemonset失败" + err.Error())
			return errors.New("反序列化daemonset失败" + err.Error())
		}
	}
	_, err = K8s.clientSet.AppsV1().StatefulSets(namespace).Update(context.TODO(), statefulSet, metaV1.UpdateOptions{})
	if err != nil {
		logger.Error("更新StatefulSet失败" + err.Error())
		return errors.New("更新StatefulSet失败" + err.Error())
	}
	return
}
