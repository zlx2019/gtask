package main

import (
	"github.com/gin-gonic/gin"
	. "zero.com/gtask-common/core"
)

func main() {
	// 创建gin的服务引擎
	engine := gin.Default()
	// 启动服务
	Run(engine, "task-user", ":8080")
}
