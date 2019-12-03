/*
源作者：码农桃花源
链接：https://www.imooc.com/article/290861
来源：慕课网

defer 配合recover
	Golang被诟病比较多的就是它的error, 经常是各种error满天飞。编程的时候总是会返回一个error, 留给调用者处理。如果是那种致命的错误，
	比如程序执行初始化的时候出问题，直接panic掉，省得上线运行后出更大的问题。但是有些时候，我们需要从异常中恢复。比如服务器程序遇到严重问题，
    产生了panic, 这时我们至少可以在程序崩溃前做一些“扫尾工作”，如关闭客户端的连接，防止客户端一直等待等等。panic会停掉当前正在执行的程序，
    不只是当前协程。在这之前，它会有序地执行完当前协程defer列表里的语句，其它协程里挂的defer语句不作保证。因此，我们经常在defer里挂一个recover语句，
    防止程序直接挂掉，这起到了 try...catch的效果。注意，recover()函数只在defer的上下文中才有效（且只有通过在defer中用匿名函数调用才有效），直接调用的话，只会返回 nil.

    下面例子panic最终会被recover捕获。这样的处理方式在一个http server的主流程常常会被用到。一次偶然的请求可能会触发某个bug, 这时用recover捕获panic, 稳住主流程，不影响其他请求。
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			// 此处不会执行
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(100)
	fmt.Println("end of main function")
}

