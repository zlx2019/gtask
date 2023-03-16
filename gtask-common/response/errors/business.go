/**
  @author: Zero
  @date: 2023/3/16 16:44:34
  @desc: 自定义业务error

**/

package errors

// ErrCode 自定义错误码类型
type ErrCode int

// Code 错误码转换为int类型
func (c ErrCode) Code() int {
	return int(c)
}

// BusinessError 项目自定义业务错误结构体
type BusinessError struct {
	Code    ErrCode //错误码
	Message string  //错误信息
}

// 重写Error(),获取错误信息
func (be *BusinessError) Error() string {
	return be.Message
}

// New 根据错误信息,构建一个code为`1`的统一业务错误
func New(message string) error {
	return &BusinessError{Code: ERROR, Message: message}
}

// NewWithCode 根据一个错误码构建一个新的error
func NewWithCode(code ErrCode) error {
	return &BusinessError{Code: code}
}

// NewWithCodeMessage 根据错误信息和错误码,构建一个新的error
func NewWithCodeMessage(message string, code ErrCode) error {
	return &BusinessError{code, message}
}
