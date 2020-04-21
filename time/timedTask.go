package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a int, b int)  {
	a = a + b
}

func byTimeSleep()  {
	for  {
		time.Sleep(time.Second)  //Sleep 阻塞当前 go 协程至少 d 时间段。d <= 0时，Sleep 会立刻返回。
		add(1,2)
		fmt.Println("byTimeSleep")
	}
}


func byTimeAfter()  {
	t := time.After(time.Second) // After 会在另一线程经过时间段 d 后向返回值发送当时的时间。等价于NewTimer(d).C。
	for  {
		select {
		case <- t:
			add(1,2)
			fmt.Println("byTimeAfter")
		}
	}
}

func byTimeTick1()  {
	t := time.Tick(time.Second) // Tick 是 NewTicker 的封装，只提供对 Ticker 的通道的访问。如果不需要关闭 Ticker，本函数就很方便。
	for  {
		select {
		case <- t:
			add(1,2)
			fmt.Println("byTimeTick1")
		}
	}
}

func byTimeTick2()  {
	t := time.NewTicker(time.Second) //NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	for  {
		select {
		case <- t.C:
			add(1,2)
			fmt.Println("byTimeTick2")
		}
	}
}


func byTimeTick3()  {
	t := time.NewTicker(time.Second)
	go func() {
		for _ = range t.C{
			add(1,2)
			fmt.Println("byTimeTick3")
		}
	}()

	time.Sleep(3 * time.Second)
	t.Stop()
}

// 不可持续执行
func byTimer()  {
	t := time.NewTimer(time.Second) //NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间

	<- t.C
	add(1,2)
	fmt.Println("byTimer")

}

// 不可持续执行
func byTimeAfterFun()  {
	w := sync.WaitGroup{}
	w.Add(1)
	time.AfterFunc(time.Second,func (){ // AfterFunc 另起一个 go 协程等待时间段 d 过去，然后调用 f。它返回一个 Timer，可以通过调用其 Stop 方法来取消等待和对 f 的调用。
		fmt.Println("byTimeAfterFun")
		w.Done()
	})
	w.Wait()
}


func main() {
	//byTimeSleep()
	//byTimeAfter()
	//byTimeTick1()
	//byTimeTick2()
	byTimeTick3()
	//byTimer()
	//byTimeAfterFun()
}
