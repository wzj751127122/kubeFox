package opention

import (
	"fmt"
	"k8s-platform/config"
	loggers "k8s-platform/logger"
	"k8s-platform/source"
	localLog "log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultConfigFile = "./setting.yaml"
)

type Options struct {
	GinEngine *gin.Engine
	// The default values.
	DB *gorm.DB
	// Factory    dao.ShareDaoFactory // 数据库接口
	ConfigFile string
}

func NewOptions() (*Options, error) {
	return &Options{
		ConfigFile: defaultConfigFile,
	}, nil
}

func (o *Options) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ConfigFile, "configFile", "", "The location of the kubeFox configuration file")
}

// Complete completes all the required options
func (o *Options) Complete() error {
	// 配置文件优先级: 默认配置，环境变量，命令行
	if len(o.ConfigFile) == 0 {
		// Try to read config file path from env.
		if cfgFile := os.Getenv("KUBEFOX-CONFIG"); cfgFile != "" {
			o.ConfigFile = cfgFile
		} else {
			o.ConfigFile = defaultConfigFile
		}
	}

	if err := config.Init(o.ConfigFile); err != nil {
		return err
	}

	// 初始化默认 api 路由
	o.GinEngine = gin.Default()

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	fmt.Println("Complete success")
	return nil

	
}

func (o *Options) InitDB() error {
	initDbService := source.NewInitDBService(o.DB)
	return initDbService.InitDB()
}

func (o *Options) register() error {
	if err := o.registerLogger(); err != nil {
		return err
	}
	if err := o.registerDatabase(); err != nil {
		return err
	}
	return nil
}

func (o *Options) registerLogger() error {
	//2. 加载日志库
	err := loggers.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		fmt.Printf("logger.Init() err:%v\n", err)
		return nil
	}
	// defer zap.L().Sync()
	return nil
}

func (o *Options) registerDatabase() error {

	newLogger := logger.New(
		localLog.New(os.Stdout, "\r\n", localLog.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	sqlConfig := config.Conf.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Dbname)
	var err error
	if o.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	}); err != nil {
		return err
	}
	// 设置数据库连接池
	sqlDB, err := o.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(config.Conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Conf.MaxOpenConns)
	// sqlDB.SetConnMaxLifetime(time.Duration(config.SysConfig.Mysql.MaxLifetime) * time.Second)
	// o.Factory = dao.NewShareDaoFactory(o.DB)
	return nil
}

// func (o *Options) registerJwt() {
// 	pkg.RegisterJwt(config.SysConfig.Default.JWTSecret)
// }
