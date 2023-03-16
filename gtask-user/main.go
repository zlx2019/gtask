package main

import (
	"github.com/gin-gonic/gin"
	. "zero.com/gtask-common/core"
	"zero.com/gtask-common/initialize"
	_ "zero.com/gtask-user/api"
	"zero.com/gtask-user/router"
)

func main() {
	// 创建gin的服务引擎
	engine := gin.Default()
	// 初始化中间件
	initialize.InitMiddleware()
	// 初始化路由
	router.InitRouter(engine)
	// TODO 初始化协程池
	// 启动服务 服务名 端口
	Run(engine, "task-user", ":8080")
}
