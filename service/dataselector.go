package service

import (
	"sort"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	nwv1 "k8s.io/api/networking/v1"
)

//用于封装排序，过滤，分页的数据类型

type dataSelector struct {
	GenericDatalist []DataCell

	DataSelect *DataSelectQuery
}

// 用于各种资源list的类型转换，转换后可以使用dataselector的排序，分页，过滤方法
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// 定义过滤和分页的结构体
// 过滤：name，分页：limit和page
// limit是单页的数据条数
// page是第几页

type DataSelectQuery struct {
	Filter   *FilterQuery
	Paginate *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

// 获取数组长度
func (d *dataSelector) Len() int {

	return len(d.GenericDatalist)

}

// 用于比较后的大小交换
func (d *dataSelector) Swap(i, j int) {

	d.GenericDatalist[i], d.GenericDatalist[j] = d.GenericDatalist[j], d.GenericDatalist[i]

}

// 定义数组中元素排序的大小比较方式
func (d *dataSelector) Less(i, j int) bool {

	a := d.GenericDatalist[i].GetCreation()
	b := d.GenericDatalist[j].GetCreation()

	return b.Before(a)

}
func (d *dataSelector) Sort() *dataSelector {

	sort.Sort(d)
	return d
}

// Filter方法用于过滤数据，比较数据的name属性，如果包含则返回
func (d *dataSelector) Filter() *dataSelector {

	//判断入参是否为空，如果为空则返回所有数据

	if d.DataSelect.Filter.Name == "" {
		return d
	}

	//如果不为空，则根据name返回数据
	filtered := []DataCell{}

	for _, value := range d.GenericDatalist {
		matches := true
		objName := value.GetName()
		if !strings.Contains(objName, d.DataSelect.Filter.Name) {
			matches = false
			continue
		}
		if matches {
			filtered = append(filtered, value)
		}

	}
	d.GenericDatalist = filtered
	return d

}

// paginate方法用于数组分页，根据limit和page的传参，返回对应的数据
func (d *dataSelector) Paginate() *dataSelector {

	limit := d.DataSelect.Paginate.Limit
	page := d.DataSelect.Paginate.Page

	// 验证参数是否合法，不合法则返回
	if limit <= 0 || page <= 0 {
		return d
	}

	startIndex := limit * (page - 1)
	endIndex := limit*page - 1

	if endIndex > len(d.GenericDatalist) {
		endIndex = len(d.GenericDatalist)
	}

	d.GenericDatalist = d.GenericDatalist[startIndex:endIndex]
	return d

}

// 定义podcell类型，重写getcreate和getname方法，可以进行数据转换
//corev1.pod -> podcell -> Datacell

type podcell corev1.Pod

func (p podcell) GetCreation() time.Time {

	return p.CreationTimestamp.Time
}

func (p podcell) GetName() string {

	return p.Name
}

type deploymentCell appsv1.Deployment

func (d deploymentCell) GetCreation() time.Time {

	return d.CreationTimestamp.Time
}

func (d deploymentCell) GetName() string {

	return d.Name
}

type daemonSetCell appsv1.DaemonSet

func (d daemonSetCell) GetCreation() time.Time {

	return d.CreationTimestamp.Time
}

func (d daemonSetCell) GetName() string {

	return d.Name
}

type statefulSetCell appsv1.StatefulSet

func (d statefulSetCell) GetCreation() time.Time {

	return d.CreationTimestamp.Time
}

func (d statefulSetCell) GetName() string {

	return d.Name
}

type ingressCell nwv1.Ingress

func (i ingressCell) GetCreation() time.Time {

	return i.CreationTimestamp.Time
}

func (i ingressCell) GetName() string {

	return i.Name
}
type serviceCell corev1.Service

func (s serviceCell) GetCreation() time.Time {

	return s.CreationTimestamp.Time
}

func (s serviceCell) GetName() string {

	return s.Name
}

type nodeCell corev1.Node

func (s nodeCell) GetCreation() time.Time {

	return s.CreationTimestamp.Time
}

func (s nodeCell) GetName() string {

	return s.Name
}


type namespaceCell corev1.Namespace

func (n namespaceCell) GetCreation() time.Time {

	return n.CreationTimestamp.Time
}

func (n namespaceCell) GetName() string {

	return n.Name
}


type persistentvolumesCell corev1.PersistentVolume

func (p persistentvolumesCell) GetCreation() time.Time {

	return p.CreationTimestamp.Time
}

func (p persistentvolumesCell) GetName() string {

	return p.Name
}


type configmapCell corev1.ConfigMap

func (c configmapCell) GetCreation() time.Time {

	return c.CreationTimestamp.Time
}

func (c configmapCell) GetName() string {

	return c.Name
}



type persistentVolumeClaimCell corev1.PersistentVolumeClaim

func (p persistentVolumeClaimCell) GetCreation() time.Time {

	return p.CreationTimestamp.Time
}

func (p persistentVolumeClaimCell) GetName() string {

	return p.Name
}


type secretCell corev1.Secret

func (s secretCell) GetCreation() time.Time {

	return s.CreationTimestamp.Time
}

func (s secretCell) GetName() string {

	return s.Name
}


