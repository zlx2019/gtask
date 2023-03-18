/**
  @author: Zero
  @date: 2023/3/18 17:13:34
  @desc: Gin全局错误处理中间件

**/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"zero.com/gtask-common/response/errors"
	"zero.com/gtask-common/response/mono"
	. "zero.com/gtask-common/session"
)

// Recover Gin请求全局错误处理
func Recover(ctx *gin.Context) {
	defer func() {
		// panic捕获
		if r := recover(); r != nil {
			m := errToString(r)
			//打印错误日志
			Logger.Error(m.Message)
			// 打印详细的错误栈信息
			Logger.Error(string(debug.Stack()))
			// 响应错误信息
			mono.FailMono(m, ctx)
			// 终止请求
			ctx.Abort()
		}
	}()
	// 放行
	ctx.Next()
}

// 获取error的错误消息
func errToString(r any) *mono.Mono {
	m := &mono.Mono{Code: errors.ERROR.Code()}
	switch v := r.(type) {
	case errors.BusinessError:
		m.Code = v.Code.Code()
		m.Message = v.Message
	case error:
		m.Message = v.Error()
	default:
		m.Message = fmt.Sprintf("%v", v)
	}
	return m
}
