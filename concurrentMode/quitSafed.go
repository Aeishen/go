/*@author : Aeishen
@data :  19/07/11, 9:42

@description : 并发的安全退出
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/


/*
	Go语言中不同Goroutine之间主要依靠管道进行通信和同步。要同时处理多个管道的发送或接收操作，我们
    需要使用select关键字。当select有多个分支时，会随机选择一个可用的管道分支，如果没有可用的管道
    分支则选择default分支，否则会一直保存阻塞状态。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

//例子1：基于select实现的管道的超时判断
func test_timeOut(c chan int,done *bool){
	for{
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second):  //超时操作
			*done = true
		}
	}
}

//例子2：select的default分支实现非阻塞的管道发送或接收操作
func test_default(c chan int){
	for{
		select {
		case v := <-c:
			fmt.Println(v)
		default:   // 没有数据时做的操作
			fmt.Println("没有数据")
		}
	}
}

//例子3：当有多个管道均可操作时，select会随机选择一个管道，基于该特性我们可以用select实现一个生成随机数序列的程序
func test_createSequence(){
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 1:
			case ch <- 2:
			}
		}
	}()


	for v := range ch {
		fmt.Println(v)
	}
}

//例子4：通过select和default分支可以很容易实现一个Goroutine的退出控制
func test_quit(cannel chan bool){
	for{
		select {
		case <-cannel:
			// 接收到退出信号退出操作
		default:
			fmt.Println("hello")
			// 正常工作操作
		}
	}
}


//例子5：管道的发送操作和接收操作是一一对应的，如果要停止多个Goroutine那么可能需要创建同样数量
// 的管道，这个代价太大了。其实我们可以通过close关闭一个管道来实现广播的效果，所有从关闭管道接收的
// 操作均会收到一个零值和一个可选的失败标志
func test_close(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cannel:
			fmt.Println("工作结束")
			// 退出
		}
	}
}

//我们通过close来关闭cancel管道向多个Goroutine广播退出的指令。不过这个程序依然不够稳健：当每个
// Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，因为
// main线程并没有等待各个工作Goroutine退出工作完成的机制。我们可以结合sync.WaitGroup来改进：
func test_close_1(cannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cannel:
			fmt.Println("工作结束")
			return
			// 退出
		}
	}
}


func main() {
    // 测试例子1
	//c := make(chan int)
	//done := false
	//go test_timeOut(c,&done)
	//for i := 0;i < 5; i++{
	//	if done{
	//		fmt.Println("超时")
	//		break
	//	}else{
	//		c <- i
	//		time.Sleep(time.Second)
	//	}
	//}


	// 测试例子2
	//c := make(chan int)
	//go test_default(c)
	//for i := 0;i < 5; i++{
	//	c <- i
	//}
	//time.Sleep(time.Second / 10)


	// 测试例子3
	// test_createSequence()


	// 测试例子4
	//cannel := make(chan bool)
	//go test_quit(cannel)
	//
	//time.Sleep(time.Second)
	//cannel <- true


	// 测试例子5
	//cancel := make(chan bool)
	//for i := 0; i < 5; i++ {
	//	go test_close(cancel)
	//}
	//time.Sleep(time.Second)
	//close(cancel)


	// 测试例子6
	//var wg sync.WaitGroup
	//cancel := make(chan bool)
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go test_close_1(cancel,&wg)
	//}
	//time.Sleep(time.Second)
	//close(cancel)
	//wg.Wait()
}

/*
	尽量保证每个工作者并发体的创建、运行、暂停和退出都是在main函数的安全控制之下了。
*/