/**
  @author: Zero
  @date: 2023/3/16 21:34:06
  @desc: 全局会话

**/

package session

import (
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var (
	// Cache redis客户端
	Cache *redis.Client
	// Logger 日志输出工具
	Logger *logrus.Logger
	// TaskPool 协程池
	TaskPool *ants.Pool
)
