package dao

// import (
// 	"fmt"
// 	"k8s-platform/config"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/spf13/viper"
// 	"github.com/wonderivan/logger"
// )

// var  (

// 	GORM *gorm.DB
// 	isInit bool
// 	err error

// )

// func Init(cfg *config.MysqlConfig)	(err error)  {

// 	// 判断是否已经初始化了
// 	if isInit {
// 		return
// 	}

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
// 	// 也可以使用MustConnect连接不成功就panic
// 	GORM, err = gorm.Open("mysql", dsn)
// 	if err != nil {
// 		logger.Error("初始化数据库失败" + err.Error())
// 		return 
// 	}
// 	GORM.LogMode(viper.GetBool("app.log_mode"))
// 	GORM.DB().SetMaxIdleConns(viper.GetInt("mysql.Max_IdleConns"))
// 	GORM.DB().SetMaxOpenConns(viper.GetInt("mysql.Max_OpenConns"))
	
// 	isInit = true
// 	logger.Info("链接数据库成功")
	
// 	return
// }

// func Close() {
// 	_ = GORM.Close()
// }
