/**
  @author: Zero
  @date: 2023/3/26 23:06:27
  @desc:

**/

package main

import (
	"zero.com/gtask-common/mq"
	"zero.com/gtask-common/mq/handler"
)

func main() {
	// 创建一个消费者监听器
	listener, err := mq.NewConsumerListener("users", "v1", &handler.DefaultMessageHandler{})
	if err != nil {
		// 监听器创建错误
		panic(err)
	}
	// 启动监听器
	if err = listener.Run("127.0.0.1:4150"); err != nil {
		// 启动错误
		panic(err)
	}
}
