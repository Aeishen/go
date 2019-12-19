/*
   @File : testEveryDay3
   @Author: Aeishen
   @Date: 2019/12/19 21:59
   @Description:
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//testEveryDay3_1()
	testEveryDay3_2()
}

// ch 未被初始化，关闭时会报错
func testEveryDay3_1()  {
	var ch chan int
    var count int
    go func() {
        ch <- 1
    }()
    go func() {
        count++
        close(ch)
	}()
	<-ch
	fmt.Println(count)
}

// 一段时间后总是输出 #goroutines:2, 程序执行到第二个 groutine 时，ch 还未初始化，导致第二个 goroutine 阻塞。需要注意的是第一个 goroutine 不会阻塞。
func testEveryDay3_2(){
	var ch chan int
	go func() {
        ch = make(chan int, 1)
        ch <- 1
    }()
    go func(ch chan int) {
        time.Sleep(time.Second)
        <-ch
    }(ch)
    c := time.Tick(1 * time.Second)
    for range c {
        fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
    }
}