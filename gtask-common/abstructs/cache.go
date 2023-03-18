/**
  @author: Zero
  @date: 2023/3/16 21:15:28
  @desc: 缓存组件抽象层

**/

package abstructs

import (
	"context"
	"time"
)

// Cache 缓存数据接口
type Cache interface {

	// Put 设置缓存
	// key 缓存数据的键
	// value 缓存数据
	// expire 过期时间
	Put(ctx context.Context, key, value string, expire time.Duration) error

	// Get 获取缓存
	// key 缓存数据的键
	Get(ctx context.Context, key string) (string, error)
}
