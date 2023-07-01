package model

import "time"

// 与mysql表对齐

type Workflow struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	CreateAt *time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"deleted_at"`

	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Ingress    string `json:"ingress"`
	Service    string `json:"service"`
	Replicas   int32  `json:"replicas"`
	Type       string `json:"type" gorm:"column:type"`
}

// 返回mysql的表名

func (w *Workflow) TableName() string {
	return "workflow"
}
