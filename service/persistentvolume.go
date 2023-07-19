package service

import (
	"context"
	"errors"

	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PersistentVolume persistentVolume

type persistentVolume struct{}

type PersistentVolumeResp struct {
	Total int                       `json:"total"`
	Items []coreV1.PersistentVolume `json:"items"`
}

func (n *persistentVolume) toCells(pvs []coreV1.PersistentVolume) []DataCell {
	cells := make([]DataCell, len(pvs))
	for i := range pvs {
		cells[i] = persistentvolumesCell(pvs[i])
	}
	return cells
}

func (n *persistentVolume) FromCells(cells []DataCell) []coreV1.PersistentVolume {
	nodes := make([]coreV1.PersistentVolume, len(cells))
	for i := range cells {
		nodes[i] = coreV1.PersistentVolume(cells[i].(persistentvolumesCell))
	}
	return nodes
}

func (n *persistentVolume) GetPersistentVolumes(filterName string, limit, page int) (*PersistentVolumeResp, error) {
	PersistentVolumeList, err := K8s.clientSet.CoreV1().PersistentVolumes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Info("获取PersistentVolumes失败" + err.Error())
		return nil, errors.New("获取PersistentVolumes失败" + err.Error())
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDatalist: n.toCells(PersistentVolumeList.Items),
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
	PersistentVolumes := n.FromCells(data.GenericDatalist)
	return &PersistentVolumeResp{
		total,
		PersistentVolumes,
	}, nil
}

// GetPersistentVolumesDetail 获取PersistentVolume详情
func (n *persistentVolume) GetPersistentVolumesDetail(Name string) (*coreV1.PersistentVolume, error) {
	PersistentVolumesRes, err := K8s.clientSet.CoreV1().PersistentVolumes().Get(context.TODO(), Name, metaV1.GetOptions{})
	if err != nil {
		logger.Info("获取PersistentVolumes详情失败" + err.Error())
		return nil, errors.New("获取PersistentVolumes详情失败" + err.Error())
	}
	return PersistentVolumesRes, nil
}

func (n *persistentVolume) DeletePersistentVolume(name string) error {
	err := K8s.clientSet.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logger.Info("删除PersistentVolumes失败" + err.Error())
		return errors.New("删除PersistentVolumes失败" + err.Error())
	}
	return nil
}
