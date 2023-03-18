/**
  @author: Zero
  @date: 2023/3/3 00:43:14
  @desc: 统一结果响应

**/
package mono

import (
	sys_err "errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"zero.com/gtask-common/response/errors"
)

const (
	SUCCESS    = 0      // 请求成功响应码
	SuccessMsg = "操作成功" //请求成功响应消息
)

// Mono 单数据响应对象
type Mono struct {
	// 响应码
	Code int `json:"code"`
	// 响应数据
	Data any `json:"data"`
	// 响应消息
	Message string `json:"message"`
}

// Just 构建响应对象
func Just(code int, data any, message string) *Mono {
	return &Mono{code, data, message}
}

// Ok 响应成功
func Ok(ctx *gin.Context) {
	OkWithMessage(SuccessMsg, ctx)
}

// OkMono 响应成功,并且返回Mono
func OkMono(m *Mono, ctx *gin.Context) {
	Result(m.Code, m.Data, m.Message, ctx)
}

// OkWithMessage 响应成功,返回响应消息
func OkWithMessage(message string, ctx *gin.Context) {
	Result(SUCCESS, nil, message, ctx)
}

// OkWithData 响应成功,返回数据
func OkWithData(data any, ctx *gin.Context) {
	Result(SUCCESS, data, SuccessMsg, ctx)
}

// Fail 响应失败
func Fail(ctx *gin.Context) {
	FailWithCode(errors.ERROR.Code(), ctx)
}

// FailWithErr 响应失败,根据业务error响应结果
func FailWithErr(err error, ctx *gin.Context) {
	var busErr *errors.BusinessError
	// 断言error,是否为自定义的业务error类型
	if sys_err.As(err, &busErr) {
		if busErr.Message == "" {
			busErr.Message = errors.Message(busErr.Code)
		}
		FailWithCodeMessage(busErr.Code.Code(), busErr.Message, ctx)
	} else {
		Fail(ctx)
	}
}

// FailWithMessage 响应失败
func FailWithMessage(message string, ctx *gin.Context) {
	FailWithCodeMessage(errors.ERROR.Code(), message, ctx)
}

// FailWithCode 响应失败,返回响应码
func FailWithCode(code int, ctx *gin.Context) {
	FailWithCodeMessage(code, errors.Message(errors.ErrCode(code)), ctx)
}

// FailMono 响应失败,返回Mono
func FailMono(m *Mono, ctx *gin.Context) {
	FailWithCodeMessage(m.Code, m.Message, ctx)
}

// FailWithCodeMessage 响应失败,返回错误码与错误消息
func FailWithCodeMessage(code int, message string, ctx *gin.Context) {
	Result(code, nil, message, ctx)
}

// Result 将结果响应回客户端
func Result(code int, data any, message string, ctx *gin.Context) {
	if message == "" {
		message = errors.Message(errors.ErrCode(code))
	}
	ctx.JSON(http.StatusOK, Just(code, data, message))
}
