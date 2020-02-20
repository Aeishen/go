/*
   @File : useRange
   @Author: Aeishen
   @Date: 2020/2/20 22:11
   @Description:Range使用
*/

package main

import "fmt"

func main() {
	// range 会复制对象
	a := [3]int{0, 1, 2}

	for i, v := range a {             // i、v 都是从复制品中取出，对 v 改变不会改变原数组
		if i == 0 {                   // 在修改前，我们先修改原数组。
			v = 999
			fmt.Println(a,v)          // 确认修改有效，输出 [999, 999, 999],999。
		}
	}

	fmt.Println(a)                    // 输出 [100, 101, 102]。

	// 改用引用类型，底层数据不会被复制
	//for i,n := 0,len(a); i < n; i++{
	//
	//}
}
