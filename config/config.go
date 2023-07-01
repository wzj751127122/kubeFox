package config
import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

var Conf = new(AppConfigMap)

type AppConfigMap struct {

	*AppConfig   `mapstructure:"app"`
	// *LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	// *RedisConfig `mapstructure:"redis"`
}

type AppConfig struct{

	// Name      string `mapstructure:"name"`
	// Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	// StartTime string `mapstructure:"start_time"`
	// MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	PodLogTailLine			int  `mapstructure:"pod_logtail"`
	LogMode					bool  `mapstructure:"log_mode"`

}

type MysqlConfig struct{

	Host string		`mapstructure:"host"`
	Port int		`mapstructure:"port"`
	Dbname string	`mapstructure:"dbname"`
	User string		`mapstructure:"user"`
	Password string	`mapstructure:"password"`
	MaxOpenConns int `mapstructure:"Max_OpenConns"`
	MaxIdleConns int `mapstructure:"Max_IdleConns"`

}

// type RedisConfig struct{

// 	Host string		`mapstructure:"host"`
// 	Port int 		`mapstructure:"port"`
// 	DB int			`mapstructure:"db"`
// 	Password string	`mapstructure:"password"`
// 	PoolSize int	`mapstructure:"poolsize"`

// }
// type LogConfig struct{

// 	Level string
//     Filename string
// 	MaxSize int
// 	MaxAge int
// 	MaxBackups int

// }

func Init(filepath string) (err error) {

	viper.SetConfigFile(filepath) // 指定配置文件路径
	// viper.SetConfigName("config")             // 配置文件名称(无扩展名)
	// viper.SetConfigType("yaml")               // 如果配置文件的名称中没有扩展名，则需要配置此项
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() err:%v\n", err)
		return
	}
//把读取到的配置信息反序列化到conf
	err = viper.Unmarshal(Conf)
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.Unmarshal(Conf) err:%v\n", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改！")
		err = viper.Unmarshal(Conf)
		if err != nil {
			fmt.Printf("viper.Unmarshal(Conf) err:%v\n", err)
		}
	})
	return
}
