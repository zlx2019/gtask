package main

import (
	"github.com/gin-gonic/gin"
	. "zero.com/gtask-common/core"
	_ "zero.com/gtask-common/initialize"
	"zero.com/gtask-common/middleware"
	"zero.com/gtask-common/session"
	_ "zero.com/gtask-user/api"
	"zero.com/gtask-user/router"
)

func main() {
	// 创建gin的服务引擎
	engine := gin.Default()
	// 使用自定义错误处理
	engine.Use(middleware.Recover)
	// 初始化路由
	router.InitRouter(engine)
	// 启动服务 服务名 端口
	server := session.Configure.Server
	Run(engine, server.Name, server.Addr)
}
