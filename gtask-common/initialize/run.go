/**
  @author: Zero
  @date: 2023/3/16 21:37:12
  @desc:

**/

package initialize

// 初始化
func init() {
	InitConfig() //初始化配置
	InitLogger() //初始化日志组件
	InitRedis()  //初始化Redis组件
	InitPool()   //初始化协程池
}
