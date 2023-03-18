package router

import (
	"github.com/gin-gonic/gin"
	"zero.com/gtask-common/session"
)

// 所有要注册的路由接口实现
var routers []Router

// RegisterRouter 注册路由函数
func RegisterRouter(child ...Router) {
	routers = append(routers, child...)
}

// InitRouter 初始化所有的路由
func InitRouter(engine *gin.Engine) {
	// 注册所有的路由
	for _, router := range routers {
		router.Route(engine)
	}
	session.Logger.Info("Gin Router Load Success")
}
