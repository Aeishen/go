/*@author : Aeishen
@data :  19/07/11, 14:09

@description : context包
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
context包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作
*/

package main

import (
	"context"
	"fmt"
	"sync"
)

// 例子1：用context包来重新实现前面的线程安全退出或超时的控制
func test_1(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// 例子2：Go语言是带内存自动回收特性的，因此内存一般不会泄漏。在前面素数筛的例子中，GenerateNatural
// 和PrimeFilter函数内部都启动了新的Goroutine，当main函数不再使用管道时后台Goroutine有泄漏的风险。
// 我们可以通过context包来避免这个问题，下面是改进的素数筛实现：

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural_1(ctx context.Context) chan int{
	c := make(chan int)
	go func(){
		for i:=2;;i++{
			select {
			case <- ctx.Done():
				return
			case c <- i:
			}
		}
	}()
	return c
}
// 管道过滤器: 删除能被素数整除的数
func PrimeFilter_1(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
        for{
			if i := <-in; i % prime != 0{
				select {
				case <- ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {

	// 测试例子1
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //当并发体超时，每个工作者都可以安全退出
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go test_1(ctx, &wg)
	//}
	//
	//time.Sleep(time.Second)
	//cancel()                //当main主动停止工作者Goroutine时，每个工作者都可以安全退出
	//
	//wg.Wait()


	// 测试例子2
	ctx, cancel := context.WithCancel(context.Background())// 通过 Context 控制后台Goroutine状态
	ch := GenerateNatural_1(ctx)  // 通过 Context 控制后台Goroutine状态

	for i:=0;i<10;i++{
		prime := <-ch      // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter_1(ctx, ch, prime) // 基于新素数构造的过滤器
	}
	cancel() //当main函数完成工作前，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。
}
