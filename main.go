package main

import (
	"fmt"
	"k8s-platform/config"
	"k8s-platform/controller"
	"k8s-platform/middle"
	"net/http"

	"k8s-platform/db"

	"k8s-platform/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
)

func main() {

	// 初始化配置文件
	err := config.Init("./setting.yaml")
	if err != nil {
		logger.Error("初始化配置文件读取失败" + err.Error())
		return
	}
	err = db.Init(config.Conf.MysqlConfig)
	if err != nil {
		logger.Error("初始化mysql失败" + err.Error())
		return
	}

	defer db.Close()
	// 初始化k8s
	service.K8s.Init()
	//初始化gin
	r := gin.Default()
	r.Use(middle.Cors())
	r.Use(middle.JWTAuthMiddleware())
	//初始化路由
	controller.Router.InitApiRouter(r)
	
	// d,_:=dao.Workflow.GetWorkflow("","blog",10,1)
	// fmt.Println(d)

	//启动websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":9999",nil)
	}()
	//运行
	err = r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
