package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PersistentVolumeClaim persistentVolumeClaim

type persistentVolumeClaim struct {
}

type PersistentVolumeClaimResp struct {
	Total int                            `json:"total"`
	Items []coreV1.PersistentVolumeClaim `json:"items"`
}

type PersistentVolumeClaimNp struct {
	NameSpace                string `json:"namespace"`
	PersistentVolumeClaimNum int    `json:"PersistentVolumeClaim_num"`
}

func (d *persistentVolumeClaim) toCells(PersistentVolumeClaims []coreV1.PersistentVolumeClaim) []DataCell {
	cells := make([]DataCell, len(PersistentVolumeClaims))
	for i := range PersistentVolumeClaims {
		cells[i] = persistentVolumeClaimCell(PersistentVolumeClaims[i])
	}
	return cells
}

func (d *persistentVolumeClaim) FromCells(cells []DataCell) []coreV1.PersistentVolumeClaim {
	PersistentVolumeClaims := make([]coreV1.PersistentVolumeClaim, len(cells))
	for i := range cells {
		PersistentVolumeClaims[i] = coreV1.PersistentVolumeClaim(cells[i].(persistentVolumeClaimCell))
	}
	return PersistentVolumeClaims
}

func (d *persistentVolumeClaim) DeletePersistentVolumeClaim(name, namespace string) error {
	err := K8s.clientSet.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logger.Error("删除PersistentVolumeClaim失败" + err.Error())
		return errors.New("删除PersistentVolumeClaim失败" + err.Error())
	}
	return nil
}

func (d *persistentVolumeClaim) UpdatePersistentVolumeClaim(content, namespace string) error {
	var PersistentVolumeClaim = &coreV1.PersistentVolumeClaim{}
	if err := json.Unmarshal([]byte(content), PersistentVolumeClaim); err != nil {
		logger.Error("反序列化PersistentVolumeClaim失败" + err.Error())
		return errors.New("反序列化PersistentVolumeClaim失败" + err.Error())
	}
	_, err := K8s.clientSet.CoreV1().PersistentVolumeClaims(namespace).Update(context.TODO(), PersistentVolumeClaim, metaV1.UpdateOptions{})
	if err != nil {
		logger.Error("更新PersistentVolumeClaim失败" + err.Error())
		return errors.New("更新PersistentVolumeClaim失败" + err.Error())
	}
	return nil
}

func (d *persistentVolumeClaim) GetPersistentVolumeClaims(filterName, namespace string, limit, page int) (*PersistentVolumeClaimResp, error) {
	PersistentVolumeClaimList, err := K8s.clientSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Error("获取PersistentVolumeClaim失败" + err.Error())
		return nil,errors.New("获取PersistentVolumeClaim失败" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDatalist: d.toCells(PersistentVolumeClaimList.Items),
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
	PersistentVolumeClaims := d.FromCells(data.GenericDatalist)
	return &PersistentVolumeClaimResp{
		Total: total,
		Items: PersistentVolumeClaims,
	}, nil
}

func (d *persistentVolumeClaim) GetPersistentVolumeClaimDetail(name, namespace string) (*coreV1.PersistentVolumeClaim, error) {
	data, err := K8s.clientSet.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		logger.Error("获取PersistentVolumeClaim详情失败" + err.Error())
		return nil,errors.New("获取PersistentVolumeClaim详情失败" + err.Error())
	}
	return data, nil
}
