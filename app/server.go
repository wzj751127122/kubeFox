package app

import (
	"context"
	"fmt"
	"k8s-platform/app/opention"
	"k8s-platform/controller"
	"k8s-platform/middle"
	"k8s-platform/service"
	"k8s-platform/utils"
	"net/http"
	"time"

	"os"


	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewServerCommand() *cobra.Command {
	
	opts, err := opention.NewOptions()
	if err != nil {
		zap.L().Fatal("unable to initialize command options: %v", zap.Any("err", err))
	}
	cmd := &cobra.Command{
		Use:  "kubeFox-run",
		Long: "The kubeFox run controller is a daemon that embeds the core control loops.",
		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = opts.InitDB(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = Run(opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 绑定命令行参数
	opts.BindFlags(cmd)
	return cmd
}

func Run(opt *opention.Options) error {
	// 打印logo
	utils.PrintLogo()
	// 设置核心应用接口
	// v1.Setup(opt)
	//初始化K8s client  TODO 未来移除
	InitLocalK8s()
	// 初始化 APIs 路由
	// router.InstallRouters(opt)
	//初始化gin
	r := gin.Default()
	r.Use(middle.Cors())
	r.Use(middle.JWTAuthMiddleware())
	//初始化路由
	controller.Router.InitApiRouter(r)
	// 启动优雅服务
	runServer(opt)
	return nil
}

func InitLocalK8s() {
	// 初始化k8s
	err := service.K8s.Init()
	if err != nil {
		utils.Must(err)
	}
}

// 优雅启动貔貅服务
func runServer(opt *opention.Options) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: opt.GinEngine,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		zap.L().Info("Success", zap.String("starting kubeFox server running on", "127.0.0.1"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("failed to listen kubeFox server: ", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 5 seconds.
	quit := utils.SetupSignalHandler()
	<-quit
	zap.L().Info("shutting kubeFox server down ...")

	// The context is used to inform the server it has 5 seconds to finish the request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("kubeFox server forced to shutdown: ", zap.Error(err))
		os.Exit(1)
	}
	zap.L().Info("kubeFox server exit successful")
}
