package service

import (
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	clientSet *kubernetes.Clientset
	Config    *rest.Config
}

func (k *k8s) Init() error {

	//clientset
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	} else {
		logger.Info("创建 k8s clientSet success")
	}

	k.clientSet = clientset
	k.Config = config
	return nil
}
