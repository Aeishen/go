package main

import "fmt"

/*
@author : Aeishen
@data :  19/07/10, 12:20

@description : 基于Channel的通信
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	Channel通信是在Goroutine之间进行同步的主要方法。

	无缓存的Channel上的发送操作总在对应的接收操作完成前发生.在无缓存的Channel上的每一次发送操作都
    有与其对应的接收操作相配对，发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine
    上执行2个操作很容易导致死锁）。

	对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是
    Channel的缓存大小（可根据控制Channel的缓存大小来控制并发执行的Goroutine的最大数目）。
*/

var done = make(chan bool)
var msg string

var done_1 = make(chan int, 10) // 带 10 个缓存

func aGoroutine_1() {
	msg = "你好, 世界"
	done <- true
}

func aGoroutine_2() {
	msg = "你好, 世界"
	close(done)
}

func aGoroutine_3() {
	msg = "你好, 世界"
	<-done
}

func aGoroutine_4(i int) {
	fmt.Println("你好, 世界")
	done_1 <- i
}


func main() {

	// 例子1
	//go aGoroutine_1()
	//<-done
	//println(msg)
	/*
		可保证打印出“hello, world”。该程序首先对msg进行写入，然后在done管道上发送同步信号，随
	    后从done接收对应的同步信号，最后执行println函数
	*/

	// 例子2
	//go aGoroutine_2()
	//<-done
	//println(msg)
	/*
		若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。因此在这个例子中,
	    用close(c)关闭管道代替done <- false依然能保证该程序产生相同的行为。
	*/

	// 例子3
	//go aGoroutine_3()
	//done <- true
	//println(msg)
	/*
		基于 “对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前“ 这个规则可
	    知，交换两个Goroutine中的接收和发送操作也是可以的（但是很危险），因为main线程中done <- true
	    发送完成前，后台线程<-done接收已经开始，这保证msg = "hello, world"被执行了，所以之后
	    println(msg)的msg已经被赋值过了。
	    但是，若该Channel为带缓冲的，main线程的done <- true接收操作将不会被后台线程的<-done
	    接收操作阻塞，该程序将无法保证打印出“hello, world”
	*/

	// 例子4

	// 开N个后台打印线程
	for i:= 0;i<cap(done_1);i++{
		go aGoroutine_4(i)
	}

	// 等待N个后台线程完成
	for i:= 0;i<cap(done_1);i++{
		<-done_1
	}
	/*
		对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup
	    来等待一组事件
	*/

}