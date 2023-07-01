package service

import (
	"k8s-platform/dao"
	"k8s-platform/model"
)


var Workflow workflow

type workflow struct{
}



type WorkflowCreate struct{

	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Image         string            `json:"image"`
	Cpu           string            `json:"cpu"`
	Memery        string            `json:"memery"`
	HealthPath    string            `json:"health_path"`
	Replicas      int32             `json:"replicas"`
	Label         map[string]string `json:"label"`
	HealthCheck   bool              `json:"health_check"`
	Hosts		  map[string][]*HttpPath   `json:"hosts"`
}

// 创建workflow

func (w *workflow)CreateWorkflow(data *WorkflowCreate)(err error)  {

	// 若workflow不是ingress类型，传入空字符串即可

	var ingressName string
	if data.Type == "Ingress"{
		ingressName = getIngressName(data.Name)
	}else{
		ingressName = ""
	}

	//组装mysql中的workflow数据
	workflow := &model.Workflow{
		Name: data.Name,
		Namespace: data.Namespace,
		Replicas: data.Replicas,
		Deployment: data.Name,
		Service: getServiceName(data.Name),
		Ingress: ingressName,
		Type: data.Type,
	}
	//调用dao层
	err = dao.Workflow.CreateWorkflow(workflow)
	if err != nil {
		return err
	}
	err = createWorkflowRes(data)
	if err != nil {
		return  err
	}
	
	return
}

func createWorkflowRes(data *WorkflowCreate)(err error)  {

	var serviceType string
	dc := &DeployCreate{
		Name: data.Name,
		Namespace: data.Namespace,
		Replicas: data.Replicas,
		Image: data.Image,
		Label: data.Label,
		Cpu: data.Cpu,
		Memery: data.Memery,
		ContainerPort: data.ContainerPort,
		HealthCheck: data.HealthCheck,
		HealthPath: data.HealthPath,

	}


	err = Deployment.CreateDeployment(dc)
	if err != nil {
		return err
	}
	
	if data.Type != "Ingress" {
		serviceType = data.Type
		
	}else{
		serviceType = "ClusterIP"
	}

	//组装service数据
	sc := &ServiceCreate{
		Name: getServiceName(data.Name),
		Namespace: data.Namespace,
		Type: serviceType,
		ContainerPort: data.ContainerPort,
		NodePort: data.NodePort,
		Port: data.Port,
		Label: data.Label,
	}


	err = Service.CreateService(sc)
	if err != nil {
		return err
	}

	if data.Type == "Ingress" {
		ic := &IngressCreate{
			Name: getIngressName(data.Name),
			Namespace: data.Namespace,
			Label: data.Label,
			Hosts: data.Hosts,
		}

		err = Ingress.CreateIngress(ic)
		if err != nil {
			return err
		}
		
	}

	return
}


// workflow名字转换为service名字，添加-svc后缀
func getServiceName(workflowName string) string {
	return workflowName + "-svc"
}

func getIngressName(workflowName string) string {
	return workflowName + "-ing"
}


// 删除workflow

func (w *workflow)DelById(id int) error {
	workflow,err := dao.Workflow.GetWorkflowById(id)
	if err != nil {
		return err
	}

	err = delworkflowRes(workflow)
	if err != nil {
		return err
	}

	err = dao.Workflow.DelWorkflow(id)
	if err != nil {
		return err
	}
	return nil
	
}

func delworkflowRes(workflow *model.Workflow)(err error)  {

	err = Deployment.DeleteDeployment(workflow.Name,workflow.Namespace)
	if err != nil {
		return err
	}

	err = Service.DeleteService(getServiceName(workflow.Name),workflow.Namespace)
	if err != nil {
		return err
	}
	if workflow.Type == "Ingress" {
		err = Ingress.DeleteIngress(getIngressName(workflow.Name),workflow.Namespace)
		if err != nil {
			return err
		}
		
	}
	return
	
}
func (w *workflow)GetList(filterName,namespace string,limit,page int) (data *dao.WorkflowResp,err error) {
	data,err = dao.Workflow.GetWorkflow(filterName,namespace,limit,page)
	if err != nil {
		return nil, err
	}

	return
}

// 获取单条

func (w *workflow)GetListById(id int) (data *model.Workflow, err error) {
	data,err = dao.Workflow.GetWorkflowById(id)
	if err != nil {
		return nil, err
	}

	return 
}