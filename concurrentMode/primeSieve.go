/*@author : Aeishen
@data :  19/07/10, 15:56

@description : 发布订阅模型
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

package main

import (
	"fmt"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int{
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//管道过滤器: 能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int{

	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("---下一个素数: %v\n", i)
				out <- i
			}else{
				fmt.Printf("---删除掉的数：%v\n", i)
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 10; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}