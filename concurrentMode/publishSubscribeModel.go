/*@author : Aeishen
@data :  19/07/10, 15:56

@description : 发布订阅模型
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	发布订阅（publish-and-subscribe）模型通常被简写为pub/sub模型。在这个模型中，消息生产者
    成为发布者（publisher），而消息消费者则成为订阅者（subscriber），生产者和消费者是M:N的关
    系。在传统生产者和消费者模型中，是将消息发送到一个队列中，而发布订阅模型则是将消息发布给一个
    主题。
*/

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者为一个管道
	topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存大小
	timeout     time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// 构建一个发布者对象, 可以设置发布超时时间和缓存队列的长度
func NewPublisher( publishTimeout  time.Duration, buffer_ int) *Publisher{
	return &Publisher{
		buffer      :  buffer_,
		timeout     :  publishTimeout,
		subscribers :  make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//订阅了全部主题
	all := p.Subscribe()

	//订阅了含有"golang"的主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for  msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}

/*
	在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、也不关心哪一个订阅者正在接收
    主题消息。订阅者和发布者可以在运行时动态添加，是一种松散的耦合关系，这使得系统的复杂性可以随时间
    的推移而增长。在现实生活中，像天气预报之类的应用就可以应用这个并发模式。
*/