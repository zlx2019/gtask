/**
  @author: Zero
  @date: 2023/3/16 22:02:06
  @desc: Redis 缓存组件

**/

package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"zero.com/gtask-common/response/errors"
	"zero.com/gtask-common/session"
)

// RedisCache 基于Redis实现的缓存组件
type RedisCache struct {
	cli *redis.Client
}

// New 创建一个Redis缓存组件实例
func New() *RedisCache {
	return &RedisCache{cli: session.Cache}
}

// Put 设置缓存数据
func (this *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	cmd := this.cli.Set(ctx, key, value, expire)
	_, err := cmd.Result()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Get 获取缓存数据
func (this *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return this.cli.Get(ctx, key).Result()
}
