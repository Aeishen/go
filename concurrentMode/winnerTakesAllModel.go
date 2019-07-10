package main

import "fmt"

/*@author : Aeishen
@data :  19/07/10, 15:56

@description : 胜者为王模型
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	采用并发编程的动机有很多：并发编程可以简化问题，比如一类问题对应一个处理线程会更简单；并发编程还可
    以提升性能，在一个多核CPU上开2个线程一般会比开1个线程快一些。其实对于提升性能而言，程序并不是简单
    地运行速度快就表示用户体验好的；很多时候程序能快速响应用户请求才是最重要的，当没有用户请求需要处理
    的时候才合适处理一些低优先级的后台任务

    假设我们想快速地搜索“golang”相关的主题，我们可能会同时打开Bing、Google或百度等多个检索引擎。
    当某个搜索最先返回结果后，就可以关闭其它搜索页面了。因为受网络环境和搜索引擎算法的影响，某些搜索
    引擎可能很快返回搜索结果，某些搜索引擎也可能等到他们公司倒闭也没有完成搜索。我们可以采用类似的策
    略来编写这个程序：
*/

//模拟
func searchByBing(s string)string{
	return s + s
}

//模拟
func searchByGoogle(s string)string{
	return s + s
}

//模拟
func searchByBaidu(s string)string{
	return s + s + s
}

func main() {
	ch := make(chan string, 32)


	go func() {
		ch <- searchByBaidu("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBing("golang")
	}()


	fmt.Println(<-ch)
}

/*
	首先，我们创建了一个带缓存的管道，管道的缓存数目要足够大，保证不会因为缓存的容量引起不必要的阻塞。
    然后我们开启了多个后台线程，分别向不同的搜索引擎提交搜索请求。当任意一个搜索引擎最先有结果之后，
    都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞）。但是最终我们只从管道取第一
    个结果，也就是最先返回的结果。

    通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应
    性能
*/