package model


type WebShellStr struct {

	Namespace string `form:"namespace"`
	Pod       string `form:"pod_name"`
	Container string `form:"container_name"`
}