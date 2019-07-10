/*
@author : Aeishen
@data :  19/07/09, 14:57

@description : 切片类型强制转换
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/


/*
	为了安全，当两个切片类型[]T和[]Y的底层原始切片类型不同时，Go 语言是无法直接转换类型的。
	不过安全都是有一定代价的，有时候这种转换是有它的价值的——可以简化编码或者是提升代码的性能。
	比如在64位系统上，需要对一个[]float64切片进行高速排序，我们可以将它强制转为[]int整数切片，
	然后以整数的方式进行排序（因为float64遵循IEEE754浮点数标准特性，当浮点数有序时对应的整数
	也必然是有序的）。
*/

package main

import (
	"reflect"
	"sort"
	"unsafe"
)

/*
将[]float64类型的切片转换为[]int类型的切片 方法1:
	第一种强制转换是先将切片数据的开始地址转换为一个较大的数组的指针，然后对数组指针对应的数组重新
	做切片操作。中间需要unsafe.Pointer来连接两个不同类型的指针传递。
*/
func SortFloat64FastV1(a []float64)(b []int){

	// 强制类型转换
	b  = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
	return
}




/*将[]float64类型的切片转换为[]int类型的切片 方法2:
	分别取到两个不同类型的切片头信息指针，任何类型的切片头部信息底层都是对应reflect.SliceHeader
    结构，然后通过更新结构体方式来更新切片信息，从而实现a对应的[]float64切片到c对应的[]int类型
    切片的转换。
*/
func SortFloat64FastV2(a []float64) (b []int){

	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	*bHdr = *aHdr
	// 以int方式给float64排序
	sort.Ints(b)
	return
}

/*
	通过基准测试，我们可以发现用sort.Ints对转换后的[]int排序的性能要比用sort.Float64s排序的性能好一点。
	不过需要注意的是，这个方法可行的前提是要保证[]float64中没有NaN(NaN 即 Not a Number,非数字,用来表
	示无效运算的结果)和Inf(即 Infinite 无穷大)等非规范的浮点数（因为浮点数中NaN不可排序，正0和负0相等，
    但是整数中没有这类情形）
*/