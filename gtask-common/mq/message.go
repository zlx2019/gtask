/**
  @author: Zero
  @date: 2023/3/26 22:21:22
  @desc: nsq消息体的封装

**/

package mq

// Message nsq消息体
type Message struct {
	// 要发送到的主题
	TopicName string
	// 消息内容
	Body []byte
}

// NewMessage 创建一个新的Message
func NewMessage(topicName string, body []byte) *Message {
	return &Message{
		TopicName: topicName,
		Body:      body,
	}
}
