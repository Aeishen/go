/*
   @File : useConst
   @Author: Aeishen
   @Date: 2020/2/20 18:40
   @Description: 常量的使用
*/

package main

import (
	"fmt"
	"unsafe"
)

// 常量必须是编译期可确定的数字，字符串，布尔值
const a,b int = 1,2
const c = "cccc"
const d bool = true

// 常量组1
const (
	e,f = 3,4
	g,h,i = 5,"hhh",false
	)

// 常量组2, 在常量组中，如不提供类型与初始值，那么视作与上一个常量相同, 前提是该常量组没有⼀次定义多个常量，例如： e,f = 3,4
const (
	k = "kkkk"
	l
)

// 常量组3, 常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
const (
	m = "mmm"
	n = len(m)
	o = unsafe.Sizeof(n)
)

// 常量组4, 如果常量类型足以存储初始化值，则不会引发溢出错误
const (
	p byte = 100 // int to byte
	//q int = 1e20 // float64 to int, overflows
)

func main() {
	const j = true  // 未使⽤局部常量不会引发编译错误。
	fmt.Println(a,b,c,d,e,f,g,h,i,k,l,m,n,o,p)
	//fmt.Println(q)
}
