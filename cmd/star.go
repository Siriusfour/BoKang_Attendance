package cmd

import (
	"Attendance/Global"
	"fmt"
)

//根据配置文件
//0.初始化配置文件
//1.初始化路由
//2.初始化日志组件
//3.初始化数据库（待定）

import Config "Attendance/config"

func Start() {
	fmt.Println("=======")

	//=======初始化日志组件
	Global.Logger = Config.InitLogger()

	//=======初始化系统配置
	Config.InitConfig()

	//=======初始化grpc客户端
	Config.InitClient()

	//=======初始化redis客户端
	Config.InitRedis()

	//======初始化MySQL数据库
	var err error
	Global.DB, err = Config.InitDB()
	if err != nil {
		panic(err)
	}

	return
}
