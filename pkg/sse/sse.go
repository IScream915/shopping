package sse

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// Event 定义 SSE 事件，包含事件名称与数据
type Event struct {
	Name string
	Data interface{}
}

// Broker 用于管理 SSE 订阅者
type Broker struct {
	// subscribers 存储每个订阅者对应的 channel
	subscribers map[chan Event]bool
	mutex       sync.RWMutex
}

// NewBroker 创建一个新的 Broker 实例
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[chan Event]bool),
	}
}

// Subscribe 为一个新的客户端创建订阅 channel
func (b *Broker) Subscribe() chan Event {
	ch := make(chan Event, 10) // 可设置缓冲，防止阻塞
	b.mutex.Lock()
	b.subscribers[ch] = true
	b.mutex.Unlock()
	return ch
}

// Unsubscribe 取消订阅并关闭 channel
func (b *Broker) Unsubscribe(ch chan Event) {
	b.mutex.Lock()
	delete(b.subscribers, ch)
	b.mutex.Unlock()
	close(ch)
}

// Publish 向所有订阅者广播一条事件消息
func (b *Broker) Publish(event Event) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	for ch := range b.subscribers {
		// 用非阻塞方式发送，避免阻塞
		select {
		case ch <- event:
		default:
		}
	}
}

// Handler 将 Broker 封装成一个 Gin 的 HTTP Handler，用于 SSE 长连接
func (b *Broker) Handler(c *gin.Context) {
	// 设置 SSE 所需的 HTTP 头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// 订阅当前请求连接
	msgChan := b.Subscribe()
	defer b.Unsubscribe(msgChan)

	// 当客户端断开连接时通知
	notify := c.Writer.CloseNotify()

	// 循环等待消息发送给客户端
	for {
		select {
		case event := <-msgChan:
			// 使用 Gin 提供的 SSEvent 方法发送消息
			c.SSEvent(event.Name, event.Data)
			// 刷新缓冲，确保消息实时推送
			c.Writer.Flush()
		case <-notify:
			// 客户端断开连接后退出循环
			return
		}
	}
}
