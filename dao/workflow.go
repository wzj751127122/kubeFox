package dao

import (
	"errors"
	"k8s-platform/app/opention"
	"k8s-platform/model"

	"github.com/wonderivan/logger"
)

var Workflow workflow

type workflow struct{}

type WorkflowResp struct {
	Items []*model.Workflow `json:"items"`
	Total int               `json:"total"`
}

//
// 获取列表分页查询

func (w *workflow) GetWorkflow(filterName, namespace string, limit, page int) (data *WorkflowResp, err error) {

	// 定义分页数据的起始位置
	startSet := (page - 1) * limit

	//定义数据库查询返回内容
	var workflowList []*model.Workflow

	tx := opention.DB.Where("name like ?", "%"+filterName+"%").Limit(limit).Offset(startSet).Order("id desc").Find(&workflowList)

	if tx.Error != nil && tx.Error.Error() != "record not found" {

		logger.Error("获取workflow失败" + tx.Error.Error())
		return nil, errors.New("获取workflow失败" + tx.Error.Error())
	}

	return &WorkflowResp{
		Items: workflowList,
		Total: len(workflowList),
	}, nil

}

// 获取单条
func (w *workflow) GetWorkflowById(id int) (workflow *model.Workflow, err error) {

	workflow = &model.Workflow{}

	tx := opention.DB.Where("id = ?", id).First(&workflow)
	if tx.Error != nil && tx.Error.Error() != "record not found" {

		logger.Error("获取单条workflow失败" + tx.Error.Error())
		return nil, errors.New("获取单条workflow失败" + tx.Error.Error())
	}

	return

}

// 表数据新增

func (w *workflow) CreateWorkflow(workflow *model.Workflow) (err error) {

	tx := opention.DB.Create(&workflow)
	if tx.Error != nil && tx.Error.Error() != "record not found" {

		logger.Error("创建workflow失败" + tx.Error.Error())
		return errors.New("创建workflow失败" + tx.Error.Error())
	}

	return

}

// 表数据删除

func (w *workflow) DelWorkflow(id int) (err error) {

	tx := opention.DB.Where("id = ?", id).Delete(&model.Workflow{})
	if tx.Error != nil {

		logger.Error("删除workflow失败" + tx.Error.Error())
		return errors.New("删除workflow失败" + tx.Error.Error())
	}

	return

}
