package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DaemonSet daemonSet

type daemonSet struct {
}

type DaemonSetResp struct {
	Total int                `json:"total"`
	Items []appsV1.DaemonSet `json:"items"`
}

type DaemonSetNp struct {
	NameSpace    string `json:"namespace"`
	DaemonSetNum int    `json:"daemonset_num"`
}

func (d *daemonSet) toCells(daemonsets []appsV1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(daemonsets))
	for i := range daemonsets {
		cells[i] = daemonSetCell(daemonsets[i])
	}
	return cells
}

func (d *daemonSet) FromCells(cells []DataCell) []appsV1.DaemonSet {
	daemonSets := make([]appsV1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsV1.DaemonSet(cells[i].(daemonSetCell))
	}
	return daemonSets
}

func (d *daemonSet) GetDaemonSets(filterName, namespace string, limit, page int) (*DaemonSetResp, error) {
	daemonSetList, err := K8s.clientSet.AppsV1().DaemonSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Info("获取daemonset列表失败" + err.Error())
		return nil, errors.New("获取daemonset列表失败" + err.Error())
	}

	selectableData := &dataSelector{
		GenericDatalist: d.toCells(daemonSetList.Items),
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

	daemonSets := d.FromCells(data.GenericDatalist)

	return &DaemonSetResp{
		Total: total,
		Items: daemonSets,
	}, nil
}

func (d *daemonSet) GetDaemonSetDetail(name, namespace string) (*appsV1.DaemonSet, error) {
	data, err := K8s.clientSet.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		logger.Info("获取daemonset详情失败" + err.Error())
		return nil, errors.New("获取daemonset详情失败" + err.Error())
	}
	return data, nil
}

func (d *daemonSet) DeleteDaemonSet(name, namespace string) (err error) {
	err = K8s.clientSet.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logger.Error("删除daemonset失败" + err.Error())
		return errors.New("删除daemonset失败" + err.Error())
	}
	return 
}

func (d *daemonSet) UpdateDaemonSet(content, namespace string) (err error) {
	var daemonset = &appsV1.DaemonSet{}
	err = json.Unmarshal([]byte(content), daemonset)
	if err != nil {
		logger.Error("反序列化daemonset失败" + err.Error())
		return errors.New("反序列化daemonset失败" + err.Error())
	}
	_, err = K8s.clientSet.AppsV1().DaemonSets(namespace).Update(context.TODO(), daemonset, metaV1.UpdateOptions{})
	if err != nil {
		logger.Error("更新daemonset失败" + err.Error())
		return errors.New("更新daemonset失败" + err.Error())
	}
	return 
}
