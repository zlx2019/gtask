/**
  @author: Zero
  @date: 2023/3/16 21:53:07
  @desc: 协程池初始化

**/

package initialize

import (
	"github.com/panjf2000/ants/v2"
	"time"
	"zero.com/gtask-common/session"
)

// InitPool ants 协程池初始化
func InitPool() {
	// 初始化一个数量为1000的协程池
	pool, err := ants.NewPool(1000, Option())
	if err != nil {
		session.Logger.Panicf("Pool Init Error: %s", err)
	}
	session.TaskPool = pool
}

// Option 获取协程池的配置参数
func Option() func(options *ants.Options) {
	return func(opt *ants.Options) {
		// 是否关闭回收空闲的work
		opt.DisablePurge = false
		// 回收空闲work的间隔。当DisablePurge为false时才生效
		// 如5 * time.Second 表示空闲5秒后的work会被回收掉
		opt.ExpiryDuration = time.Hour
		// 在初始化池时是否进行内存预分配。
		opt.PreAlloc = true
		//指定是否使用非阻塞模式执行任务。如果设置为true，则在协程池已满的情况下，任务会立即返回一个err，而不是等待空闲协程。
		// false表示不开启,阻塞等待可用的协程。
		opt.Nonblocking = false
		// 阻塞模式下,最多允许阻塞等待的协程数量。
		opt.MaxBlockingTasks = 100
		// 设置日志器
		opt.Logger = session.Logger
		// 指定一个函数用于处理协程中的 panic 异常。
		// TODO 暂时没有好的方案 不处理
		opt.PanicHandler = func(i interface{}) {
			session.Logger.Errorf("Pool Error")
		}
	}
}
