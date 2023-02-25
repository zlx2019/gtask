package router

import "github.com/gin-gonic/gin"

// Router 路由抽象接口,每个业务模块各自实现该接口,进行路由绑定。
type Router interface {
	// Route 注册路由的方法
	Route(e *gin.Engine)
}
