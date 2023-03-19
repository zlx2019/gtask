/**
  @author: Zero
  @date: 2023/3/16 21:29:26
  @desc: redis客户端初始化

**/

package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"zero.com/gtask-common/session"
)

// InitRedis 初始化Redis客户端
func InitRedis() {
	r := session.Configure.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     r.GetAddr(),
		Password: r.Password,
		DB:       r.Db,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("Conn Redis Success")
			return nil
		},
	})
	session.Cache = client
}
