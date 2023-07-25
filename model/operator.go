package model
import (
	"context"
	"time"

	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(OperatorationOrder, &SysOperationRecord{})
}

type OperationListInput struct {
	PageInfo
	Method string `json:"method" form:"method" ` // 请求方法
	Path   string `json:"path" form:"path" `     // 请求路径
	Status int    `json:"status" form:"status" ` // 请求状态
}

type OperationListOutPut struct {
	Total         int64                       `json:"total"`
	OperationList []*SysOperationRecord `json:"list"`
	PageInfo
}


type SysOperationRecord struct {
	ID           int           `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`  // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                    // 用户id
	User         SysUser       `json:"user"`
	CreatedAt     time.Time                                  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time                                  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt                             `gorm:"index" json:"-"`
}

func (s *SysOperationRecord) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(s)
}

func (s *SysOperationRecord) InitData(ctx context.Context, db *gorm.DB) error {
	// 审计表，不需要初始化数据
	return nil
}

func (s *SysOperationRecord) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	return true, nil
}

func (s *SysOperationRecord) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&s)
}

func (s *SysOperationRecord) TableName() string {
	return "sys_operation_record"
}
