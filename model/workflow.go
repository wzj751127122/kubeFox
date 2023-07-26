package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// 与mysql表对齐

func init() {
	RegisterInitializer(WorkFlowOrder, &Workflow{})
}

type Workflow struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	CreateAt *time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"deleted_at"`

	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Deployment  string `json:"deployment"`
	Ingress     string `json:"ingress"`
	Service     string `json:"service"`
	Replicas    int32  `json:"replicas"`
	ServiceType string `json:"service_type" gorm:"column:service_type"`
}

// 返回mysql的表名
func (w *Workflow) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&w)
}

func (w *Workflow) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return true, nil
}

func (w *Workflow) InitData(ctx context.Context, db *gorm.DB) error {
	return nil
}

func (w *Workflow) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(w)
}

func (w *Workflow) TableName() string {
	return "t_workflow"
}

func GetWorkflowTableName() string {
	temp := &Workflow{}
	return temp.TableName()
}
