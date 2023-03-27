/**
  @author: Zero
  @date: 2023/3/26 23:07:01
  @desc:

**/

package main

import (
	"time"
	"zero.com/gtask-common/mq"
)

func main() {
	// 创建一个生产者客户端
	client, err := mq.NewProducerClient("127.0.0.1:4150")
	if err != nil {
		panic(err)
	}
	// 运行客户端
	go client.Run()

	for {
		client.Publish("users", []byte("你好呀哈哈哈"))
		time.Sleep(time.Second)
	}

}
