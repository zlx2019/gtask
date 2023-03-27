/**
  @author: Zero
  @date: 2023/3/27 13:27:05
  @desc: 默认的消息处理

**/

package handler

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

// DefaultMessageHandler 默认的消息处理
type DefaultMessageHandler struct {
}

// HandleMessage 自定义的消息处理函数
func (DefaultMessageHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return nil
	}
	msg := string(message.Body)
	fmt.Printf("消费消息: %s \n", msg)
	return nil
}
