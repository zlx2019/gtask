/**
  @author: Zero
  @date: 2023/3/26 22:05:36
  @desc: nsq的客户端封装

**/

package mq

import (
	"github.com/nsqio/go-nsq"
	"log"
	"zero.com/gtask-common/session"
)

// ProducerClient Nsq的生产者封装对象
type ProducerClient struct {
	producer *nsq.Producer // nsq的生产者对象
	stop     chan struct{} // 生产者客户端关闭信号
	messages chan *Message // 要发布的消息通道
}

// NewProducerClient 创建一个生产者客户端
// addr nsq地址
func NewProducerClient(addr string) (*ProducerClient, error) {
	// 创建nsq配置
	config := nsq.NewConfig()
	// 创建生产者
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}
	// 设置日志级别
	producer.SetLogger(log.Default(), nsq.LogLevelInfo)
	// 构建客户端
	client := &ProducerClient{
		producer: producer,
		stop:     make(chan struct{}),
		messages: make(chan *Message, 1),
	}
	return client, nil
}

// Publish   发布消息
// topicName 要发送到的主题
// message   要发布的消息
func (client *ProducerClient) Publish(topicName string, message []byte) {
	// 将数据包装成 封装好的消息体,然后推到该客户端的消息通道
	client.messages <- NewMessage(topicName, message)
}

// Pub 发布消息(直接发布,不使用消息通道)
func (client ProducerClient) Pub(topicName string, message []byte) error {
	return client.producer.Publish(topicName, message)
}

// Run 客户端消息通道监听
// 将消息通道中的消息发布统一发布到nsqd中
func (client *ProducerClient) Run() {
	for {
		select {
		case <-client.stop:
			//接收到客户端关闭信号
			client.producer.Stop()
		case msg := <-client.messages:
			// 接收到消息
			err := client.producer.Publish(msg.TopicName, msg.Body)
			if err != nil {
				session.Logger.Errorf("nsq 消息发布错误: " + err.Error())
			}
		}
	}
}
