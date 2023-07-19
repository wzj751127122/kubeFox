package service

import (
	"context"
	"errors"

	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Node node

type node struct{}

type NodeResp struct {
	Total int           `json:"total"`
	Items []coreV1.Node `json:"items"`
}

func (n *node) toCells(nodes []coreV1.Node) []DataCell {
	cells := make([]DataCell, len(nodes))
	for i := range nodes {
		cells[i] = nodeCell(nodes[i])
	}
	return cells
}

func (n *node) FromCells(cells []DataCell) []coreV1.Node {
	nodes := make([]coreV1.Node, len(cells))
	for i := range cells {
		nodes[i] = coreV1.Node(cells[i].(nodeCell))
	}
	return nodes
}

func (n *node) GetNodes(filterName string, limit, page int) (nodesResp *NodeResp, err error) {
	nodeList, err := K8s.clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Info("获取node列表失败" + err.Error())
		return nil, errors.New("获取node列表失败" + err.Error())
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDatalist: n.toCells(nodeList.Items),
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
	nodes := n.FromCells(data.GenericDatalist)
	return &NodeResp{
		total,
		nodes,
	}, nil
}

// GetNodeDetail 获取Node详情
func (n *node) GetNodeDetail(Name string) (*coreV1.Node, error) {
	nodeRes, err := K8s.clientSet.CoreV1().Nodes().Get(context.TODO(), Name, metaV1.GetOptions{})
	if err != nil {
		logger.Info("获取node详情失败" + err.Error())
		return nil, errors.New("获取node详情失败" + err.Error())
	}
	return nodeRes, nil
}
