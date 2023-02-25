package member

import (
	"github.com/gin-gonic/gin"
)

// HandlerMember Member的路由处理函数
type HandlerMember struct {
}

// MemberLogin 登录
func (*HandlerMember) MemberLogin(ctx *gin.Context) {
	ctx.JSON(200, map[string]any{"name": "张三"})
}
