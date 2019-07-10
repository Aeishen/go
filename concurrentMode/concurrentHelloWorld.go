/*
@author : Aeishen
@data :  19/07/10, 15:17

@description : 并发版本的HelloWorld
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

package main

import "fmt"

//例子1: 通过互斥量sync.Mutex来实现同步通信(根据文档，我们不能直接对一个未加锁状态的sync.Mutex进行解锁，
// 这会导致运行时异常。下面这种方式并不能保证正常工作)
//func main() {
//	var mu sync.Mutex
//
//	go func() {
//		fmt.Println("你好, 世界")
//		mu.Lock()
//	}()
//
//	mu.Unlock()
//}

/*
	因为mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型。同时它们
    也没有其它的同步事件可以参考，这两个事件不可排序也就是可以并发的。因为可能是并发的事件，所以main函
	数中的mu.Unlock()很有可能先发生，而这个时刻mu互斥对象还处于未加锁的状态，从而会导致运行时异常
*/

//例子2: 通过斥量sync.Mutex来实现同步通信,改进上面的例子
//func main() {
//	var mu sync.Mutex
//
//	mu.Lock()
//	go func() {
//		fmt.Println("你好, 世界")
//		mu.Unlock()
//	}()
//	mu.Lock()
//}

/*
	在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
	main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unlock()时解锁，此时打印工作
	已经完成了，解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，此时后台线程和主线程再没有其
	它的同步事件参考，它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出
    了，也可能没有退出。虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的
*/

//例子3: 使用sync.Mutex互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步
//func main() {
//	done := make(chan bool)
//
//	go func() {
//		fmt.Println("你好, 世界")
//		<-done
//	}()
//
//	done <- true
//}
/*
	对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前
    该例子虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证main退出之前
    后台线程能正常打印了。更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓
    存大小的影响
*/

//例子4: 使用带缓存的管道来实现同步
func main() {
	done := make(chan bool,1)

	go func() {
		fmt.Println("你好, 世界")
		done <- true
	}()

	<-done
}
