package main

import (
	"os"

	"k8s-platform/app"

	"github.com/gin-gonic/gin"
	// "github.com/wonderivan/logger"
)

func main() {


	gin.SetMode(gin.ReleaseMode)
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}


}
