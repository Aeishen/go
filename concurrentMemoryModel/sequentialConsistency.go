/*
@author : Aeishen
@data :  19/07/10, 11:44

@description : 顺序一致性内存模型
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	在Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不
	满足顺序一致性内存模型，需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两
	个事件是并发的。为了最大化并行，Go语言的编译器和处理器在不影响上述规定的前提下可能会对执行语句重新排序
	（CPU也会对一些指令进行乱序执行）

	因此，如果在一个Goroutine中顺序执行a = 1; b = 2;两个语句，虽然在当前的Goroutine中可以认为a = 1;
	语句先于b = 2;语句执行，但是在另一个Goroutine中b = 2;语句可能会先于a = 1;语句执行，甚至在另一个
	Goroutine中无法看到它们的变化（可能始终在寄存器中）。也就是说在另一个Goroutine看来, a = 1; b = 2;
	两个语句的执行顺序是不确定的。如果一个并发程序无法确定事件的顺序关系，那么程序的运行结果往往会有不确定
	的结果。
*/

package main

import "sync"

// 例子1
func test_4(){
	go println("你好, 世界")
}

/*
	例子1根据Go语言规范，main函数退出时程序结束，不会等待任何后台线程。因为Goroutine的执行和main函数
    的返回事件是并发的，谁都有可能先发生，所以什么时候打印，能否打印都是未知的。

	用前面的原子操作并不能解决问题，因为我们无法确定两个原子操作之间的顺序
*/

//例子2 通过同步原语来给两个事件明确排序
func test_5() {
	done := make(chan int)

	go func(){
		println("你好, 世界")
		done <- 1
	}()

	<-done
}
/*
	当<-done执行时，必然要求done <- 1也已经执行。根据同一个Gorouine依然满足顺序一致性规则，我们可以
	判断当done <- 1执行时，println("你好, 世界")语句必然已经执行完成了。因此，现在的程序确保可以正常
    打印结果。
*/

//例子3：通过sync.Mutex互斥量实现同步
func test_6() {
	var mu sync.Mutex

	mu.Lock()
	go func(){
		println("你好, 世界")
		mu.Unlock()
	}()
	mu.Lock()
}
/*
	可以确定后台线程的mu.Unlock()必然在println("你好, 世界")完成后发生（同一个线程满足顺序一致性），
	main函数的第二个mu.Lock()必然在后台线程的mu.Unlock()之后发生（sync.Mutex保证），此时后台线
    程的打印工作已经顺利完成了
*/

func main() {
	test_6()
}