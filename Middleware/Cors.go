package Middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {

	//设置访问规则
	config := cors.Config{
		//是否允许所有的站点跨域访问
		AllowOriginFunc: func(origin string) bool {
			return true
		},

		//允许访问的站点列表
		AllowOrigins: []string{"https://foo.com"},
		//允许使用的方法
		AllowMethods:        []string{"PUT", "POST", "GET"},
		AllowPrivateNetwork: false,
		//允许客户端设置的HTTP头部信息
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization", "accept"},
		//允许客户端脚本读到的信息
		ExposeHeaders: []string{"Content-Length"},
		//允许cookie
		AllowCredentials: true,
	}

	return cors.New(config)
}
