/*
@author : Aeishen
@data :  19/07/10, 10:20

@description : 原子操作
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	所谓的原子操作就是并发编程中“最小的且不可并行化”的操作。通常，如果多个并发体对同一个共享资源进行的操作
	是原子的话，那么同一时刻最多只能有一个并发体对该资源进行操作。从线程角度看，在当前线程修改共享资源期间，
	其它的线程是不能访问该资源的。原子操作对于多线程并发编程模型来说，不会发生有别于单线程的意外情况，共享
	资源的完整性可以得到保证。

    一般情况下，原子操作都是通过“互斥”访问来保证的，通常由特殊的CPU指令提供保护
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 1.借助于sync.Mutex模拟下粗粒度的原子操作
var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func test_1(){
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}

// 2.用互斥锁来保护一个数值型的共享资源，麻烦且效率低下，标准库的sync/atomic包对原子操作提供了丰富的支持。重新实现上面的例子：
var total1 int32
func worker1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i int32
	for i = 0; i <= 10; i++ {
		atomic.AddInt32(&total1, i)
	}
}

func test_2(){
	var wg sync.WaitGroup
	wg.Add(2)
	go worker1(&wg)
	go worker1(&wg)
	wg.Wait()

	fmt.Println(total1)
}

//3.原子操作配合互斥锁可以实现非常高效的单例模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可
//  以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能

type singleton struct {}
var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

//4.将3通用的代码提取出来，就成了标准库中sync.Once的实现
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

//基于sync.Once重新实现单例模式
var (
	instance1 *singleton
	once     sync.Once
)
func Instance1() *singleton {
	once.Do(func() {
		instance1 = &singleton{}
	})
	return instance1
}

//5.sync/atomic包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。atomic.Value原子对象提供
// 了Load和Store两个原子方法，分别用于加载和保存数据，返回值和参数都是interface{}类型，因此可以
// 用于任意的自定义复杂类型。
func test3(){
	var wg sync.WaitGroup

	// 模拟配置加载
	loadConfig := func()int{
		rand.Seed(time.Now().Unix())
		return rand.Intn(100)
	}

	// 模拟请求配置的需求者
	requester := []byte{'a','b','c'}

	var config atomic.Value // 保存当前配置信息
	config.Store(loadConfig())// 初始化配置信息

	// 启动一个后台线程, 加载更新后的配置信息
	go func() {
		for {
			time.Sleep(time.Second / 2)
			config.Store(loadConfig())
		}
	}()

	// 用于处理请求的工作者线程始终采用最新的配置信息
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func() {
			fmt.Println( "---------------获取最新配置信息----------------" )
			for _,v := range requester {
				c := config.Load()
				fmt.Printf( "%s get new config is %v\n", string(v),c)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}



func main() {
	//test_1()
	/*
		11.在worker的循环中，为了保证total.value += i的原子性，我们通过sync.Mutex加锁和解锁来保证
	    该语句在同一时刻只被一个线程访问。对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必
	    须的。如果没有锁的保护，total的最终值将由于多线程之间的竞争而可能会不正确。
	*/

	//test_2()
	/*
        22.atomic.AddInt32函数调用保证了total1的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
	*/

	test3()
	/*
		这是一个简化的生产者消费者模型：后台线程生成最新的配置信息；前台多个工作者线程获取最新的配置信息。所有线程
	    共享配置信息资源。
	*/
}