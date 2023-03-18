/**
  @author: Zero
  @date: 2023/3/16 16:25:07
  @desc: 统一定义的系统error

**/

package errors

var (
	BusinessErr       = NewWithCode(ERROR)                 // 系统未知错误
	MobileNotLegalErr = NewWithCode(MobileNotLegal)        // 手机号码格式非法错误
	PoolTaskCreateErr = NewWithCode(PoolTaskCreateErrCode) //从协程池获取协程失败
)
