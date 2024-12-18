package main

import (
	"Attendance/Global"
	"Attendance/Router"
	"Attendance/cmd"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//	初始化
	cmd.Start()

	//开启服务
	ctx, cancelCTX := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCTX()

	//gin的引擎实例
	r := gin.New()

	//初始化路由
	Router.InitRouter(r)

	//拿到配置文件的指定端口号
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "9999"
	}

	Server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	//以一个goroutine开启服务
	go func() {
		Global.Logger.Info(fmt.Sprintf("正在开启，监听端口: %s", stPort))
		if err := Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Global.Logger.Error(fmt.Sprintf("listen：%s，启动失败：%s", stPort, err))
			return
		}
		fmt.Printf("开启成功！端口：%s", stPort)
	}()

	//优雅地关闭服务
	<-ctx.Done()
	ctx, cancelShudown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShudown()

	err := Server.Shutdown(ctx)
	if err != nil {
		Global.Logger.Error(fmt.Sprintf("listen：%s，停止服务失败：%s", stPort, err))
		fmt.Printf("停止服务失败%s", err.Error())
	}
	Global.Logger.Info(fmt.Sprintf("服务关闭: %s", stPort))
}
