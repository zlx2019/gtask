/**
  @author: Zero
  @date: 2023/3/3 11:58:03
  @desc: 系统响应码和响应消息
**/

package errors

// 统一错误码
const (
	ERROR          ErrCode = iota + 1 //ERROR 请求失败。
	MobileNotLegal         = 2001     // 手机号码不合法
)

// ResultMap 错误码与响应消息的映射
var resultMap = map[ErrCode]string{
	ERROR:          "请求失败",
	MobileNotLegal: "手机号码格式不正确",
}

// Message 根据错误码获取响应消息
func Message(code ErrCode) string {
	if message, ok := resultMap[code]; ok {
		return message
	}
	return resultMap[ERROR]
}
