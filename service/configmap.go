package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Configmap configmap

type configmap struct{}

type ConfigmapResp struct {
	Total int                `json:"total"`
	Items []coreV1.ConfigMap `json:"items"`
}

type ConfigmapNp struct {
	NameSpace    string `json:"namespace"`
	ConfigmapNum int    `json:"configmap_num"`
}

func (d *configmap) toCells(Configmaps []coreV1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(Configmaps))
	for i := range Configmaps {
		cells[i] = configmapCell(Configmaps[i])
	}
	return cells
}

func (d *configmap) FromCells(cells []DataCell) []coreV1.ConfigMap {
	Configmaps := make([]coreV1.ConfigMap, len(cells))
	for i := range cells {
		Configmaps[i] = coreV1.ConfigMap(cells[i].(configmapCell))
	}
	return Configmaps
}

func (d *configmap) GetConfigmaps(filterName, namespace string, limit, page int) (*ConfigmapResp, error) {
	ConfigmapList, err := K8s.clientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Error("获取configmap失败" + err.Error())
		return nil,errors.New("获取configmap失败" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDatalist: d.toCells(ConfigmapList.Items),
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
	Configmaps := d.FromCells(data.GenericDatalist)
	return &ConfigmapResp{
		Total: total,
		Items: Configmaps,
	}, nil
}

func (d *configmap) GetConfigmapDetail(name, namespace string) (*coreV1.ConfigMap, error) {
	data, err := K8s.clientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		logger.Error("获取configmap详情失败" + err.Error())
		return nil,errors.New("获取configmap详情失败" + err.Error())
	}
	return data, nil
}

func (d *configmap) DeleteConfigmap(name, namespace string) error {
	err := K8s.clientSet.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logger.Error("获取configmap失败" + err.Error())
		return errors.New("获取configmap失败" + err.Error())
	}
	return err
}

func (d *configmap) UpdateConfigmap(content, namespace string) error {
	var Configmap = &coreV1.ConfigMap{}
	if err := json.Unmarshal([]byte(content), Configmap); err != nil {
		logger.Error("反序列化deployment失败" + err.Error())
		return errors.New("反序列化deployment失败" + err.Error())
	}
	_, err := K8s.clientSet.CoreV1().ConfigMaps(namespace).Update(context.TODO(), Configmap, metaV1.UpdateOptions{})
	if err != nil {
		logger.Error("获取configmap失败" + err.Error())
		return errors.New("获取configmap失败" + err.Error())
	}
	return nil
}
