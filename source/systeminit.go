package source

import (
	"context"
	"k8s-platform/model"

	"strings"

	"gorm.io/gorm"
)

func init() {
	RegisterInit(&SystemInitTable{})
}

type SystemInitTable struct {
}

func (s *SystemInitTable) InitializerName() string {
	return strings.ToUpper("SystemInitTable")
}

func (s *SystemInitTable) MigrateTable(ctx context.Context, db *gorm.DB) error {
	for _, initializer := range model.InitializerList {
		if err := initializer.MigrateTable(ctx, db); err != nil {
			return err
		}
	}
	return nil
}

func (s *SystemInitTable) InitializeData(ctx context.Context, db *gorm.DB) error {
	for _, initializer := range model.InitializerList {
		if err := initializer.InitData(ctx, db); err != nil {
			return err
		}
	}
	return nil
}

func (s *SystemInitTable) TableCreated(ctx context.Context, db *gorm.DB) bool {
	yes := true
	for _, initializer := range model.InitializerList {
		yes = yes && db.Migrator().HasTable(initializer.TableName())
	}
	return yes
}
