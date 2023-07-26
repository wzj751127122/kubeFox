package main

import (

	"os"

	"k8s-platform/app"


	"github.com/gin-gonic/gin"
	// "github.com/wonderivan/logger"
)

func main() {

	// // 初始化配置文件
	// err := config.Init("./setting.yaml")
	// if err != nil {
	// 	fmt.Printf("settings.Init() err:%v\n", err)
	// 	return
	// }

	// fmt.Println("mode=", config.Conf.Mode)

	// err = dao.Init(config.Conf.MysqlConfig)
	// if err != nil {
	// 	fmt.Printf("初始化mysql err:%v\n", err)
	// 	return
	// }

	// defer dao.Close()
	gin.SetMode(gin.ReleaseMode)
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	// d,_:=dao.Workflow.GetWorkflow("","blog",10,1)
	// fmt.Println(d)

	//启动websocket
	// go func() {
	// 	http.HandleFunc("/ws", service.Terminal.WsHandler)
	// 	http.ListenAndServe(":9999", nil)
	// }()
	//运行
	// err = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	// if err != nil {
	// 	fmt.Printf("run server failed, err:%v\n", err)
	// 	return
	// }
}
