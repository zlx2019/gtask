/**
  @author: Zero
  @date: 2023/3/16 21:34:06
  @desc: 全局会话 中间件客户端

**/

package session

import "github.com/go-redis/redis/v8"

var (
	// Cache redis客户端
	Cache *redis.Client
)
