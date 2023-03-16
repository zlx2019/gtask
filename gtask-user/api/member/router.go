package member

import (
	"github.com/gin-gonic/gin"
	"zero.com/gtask-user/router"
)

// RouterMember Member路由
type RouterMember struct{}

// 初始化时,注册此路由
func init() {
	router.RegisterRouter(&RouterMember{})
}

// Route Member的路由映射
func (m *RouterMember) Route(e *gin.Engine) {
	// member的处理函数
	hm := New()
	// 获取手机验证码
	e.POST("/project/login/getCaptcha", hm.GetSms)
}
