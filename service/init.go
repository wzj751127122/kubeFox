package service

import (
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	clientSet *kubernetes.Clientset
}

func (k *k8s) Init() {

	//clientset
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic("获取 k8s 配置失败"+err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("创建 k8s clientSet失败" + err.Error())
	} else {
		logger.Info("创建 k8s clientSet success")
	}

	k.clientSet = clientset
}
