```
并发不是并行。并发更关注的是程序的设计层面，并发的程序完全是可以顺序执行的，只有在真正的多核CPU上才可能真正地同时运行。并行更关注的是程序的运行层面，并行一般是简单的大量重复，例如GPU中对图像处理都会有大量的并行运算。为更好的编写并发程序，从设计之初Go语言就注重如何在编程语言层级上设计一个简洁安全高效的抽象模型，让程序员专注于分解问题和组合方案，而且不用被线程管理和信号互斥这些繁琐的操作分散精力。

并发编程的核心概念是同步通信，在并发编程中，对共享资源的正确访问需要精确的控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而Go语言却另辟蹊径，它将共享的值通过Channel传递(实际上多个独立执行的线程很少主动共享资源)。在任意给定的时刻，最好只有一个Goroutine能够拥有该资源。数据竞争从设计层面上就被杜绝了。为了提倡这种思考方式，Go语言将其并发编程哲学化为一句口号：Do not communicate by sharing memory; instead, share memory by communicating.（不要通过共享内存来通信，而应通过通信来共享内存）。这是更高层次的并发编程哲学(通过管道来传值是Go语言推荐的做法)。虽然像引用计数这类简单的并发问题通过原子操作或互斥锁就能很好地实现，但是通过Channel来控制访问能够让你写出更简洁正确的程序
```



@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著