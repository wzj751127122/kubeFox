package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

var Terminal terminal

type terminal struct{}

func (t *terminal) WsHandler(w http.ResponseWriter, r *http.Request){

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		logger.Error("加载配置文件失败" + err.Error())
		return
	}
	//解析前端入参，获取namespace，podname，container
	err = r.ParseForm()
	if err != nil {
		logger.Error("解析失败" + err.Error())
		return
	}

	namespace := r.Form.Get("namespace")
	podName := r.Form.Get("pod_name")
	containerName := r.Form.Get("container_name")
	logger.Info("exec pod: %s,container_name: %s,namespace: %s", podName, containerName, namespace)

	pty, err := NewTerminalSession(w, r, nil)
	if err != nil {
		logger.Error("实例化terminalSession失败" + err.Error())
		return
	}

	defer func() {
		logger.Info("关闭terminalSession")
		pty.Close()
	}()

	req := K8s.clientSet.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(namespace).SubResource("exec").VersionedParams(
		&v1.PodExecOptions{
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
			Container: containerName,
			Command:   []string{"/bin/sh"},
		}, scheme.ParameterCodec)

	logger.Info("exec post request url:", req)

	//升级SPDY协议
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		logger.Error("建立SPDY失败" + err.Error())
		return
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             pty,
		Stdout:            pty,
		Stderr:            pty,
		Tty:               true,
		TerminalSizeQueue: pty,
	})
	if err != nil {
		logger.Error("执行 pod 命令失败" + err.Error())
		//将报错返回web
		pty.Write([]byte("执行 pod 命令失败" + err.Error()))
		pty.Done()
	}

}

//消息内容

type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

// 交互的结构体，接管输入和输出
type TerminalSession struct {
	ws       *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

//初始化一个websocket upgrade类型的对象，用于http升级成为ws协议

var upgrader = func() websocket.Upgrader {

	upgrader := websocket.Upgrader{}
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}()

// 创建terminalSession类型的对象并返回
func NewTerminalSession(w http.ResponseWriter, r *http.Request, responseHander http.Header) (*TerminalSession, error) {

	conn, err := upgrader.Upgrade(w, r, responseHander)
	if err != nil {
		return nil, errors.New("升级websocket失败" + err.Error())
	}
	session := &TerminalSession{
		ws:       conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

// 返回值int是读成功多少数据
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.ws.ReadMessage()
	if err != nil {
		log.Panicf("read message err ,%v", err)
		return 0, err
	}
	var msg TerminalMessage
	err = json.Unmarshal(message, &msg)
	if err != nil {
		log.Panicf("read parse message err ,%v", err)
		return 0, err
	}

	switch msg.Operation {
	case "stdin":
		return copy(p, []byte(msg.Data)), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	case "ping":
		return 0, nil
	default:
		log.Printf("unknow message type %s", msg.Operation)
		return 0, fmt.Errorf("unknow message type %s", msg.Operation)
	}

}

// 写数据的方法，拿到k8s的返回内容向终端输出
func (t *TerminalSession) Write(p []byte) (int, error) {

	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		log.Panicf("write parse message err ,%v", err)
		return 0, err
	}
	err = t.ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Panicf("write message err ,%v", err)
		return 0, err
	}
	return len(p), nil

}

func (t *TerminalSession) Done() {
	close(t.doneChan)
}

func (t *TerminalSession) Close() {
	t.ws.Close()
}

// resize方法，是否退出终端
func (t *TerminalSession) Next() *remotecommand.TerminalSize {

	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}
