/**
  @author: Zero
  @date: 2023/3/26 22:40:41
  @desc: nsq消费者监听器封装

**/

package mq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

// ConsumerListener 消费者监听器
type ConsumerListener struct {
	consumer    *nsq.Consumer //nsq消费者
	topicName   string        //要消费的主题
	channelName string        //要消费的channel
	stop        chan struct{} //关闭监听器的信号通道
}

// NewConsumerListener 创建一个消费者监听器
// TopicName   要监听的主题
// channelName 要消费的频道
// handler     消息的处理器
func NewConsumerListener(TopicName, channelName string, handler nsq.Handler) (*ConsumerListener, error) {
	// 创建配置
	config := nsq.NewConfig()
	// 创建消费者
	consumer, err := nsq.NewConsumer(TopicName, channelName, config)
	if err != nil {
		return nil, err
	}
	// 添加消息处理函数
	consumer.AddHandler(handler)
	// 构建监听器
	listener := &ConsumerListener{
		consumer:    consumer,
		topicName:   TopicName,
		channelName: channelName,
		stop:        make(chan struct{}),
	}
	return listener, nil
}

// Run 		启动消费者客户端,并且注册消息处理器
// addr 	nsq的地址
func (listener *ConsumerListener) Run(addr string) error {
	// 让消费者连接nsq中的主题
	err := listener.consumer.ConnectToNSQD(addr)
	if err != nil {
		return err
	}
	fmt.Println("启动消费者监听器成功~")
	// 阻塞,不停的消费来自于该主题的消息
	select {
	// 监听器收到,关闭信号 关闭消费者,退出程序.
	case <-listener.stop:
		listener.consumer.Stop()
		fmt.Println("消费者监听器已关闭~")
	}
	return nil
}

// Stop 监听器关闭
func (listener *ConsumerListener) Stop() {
	listener.stop <- struct{}{}
}
