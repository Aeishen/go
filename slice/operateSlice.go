/*
@author : Aeishen
@data :  19/07/09, 15:24

@description : 对切片操作
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

package main



func main() {

// 给切片添加元素------
	// 1.内置的泛型函数append可以在切片的尾部追加N个元素
	var a []int
	a = append(a,1)           // 追加1个元素
	a = append(a,2,3)         // 追加多个元素
	a = append(a,[]int{1,2,3}...)   // 追加一个切片
	/*
		在容量不足的情况下，append的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。
		即使容量足够，依然需要用append函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。
	*/

	// 2.在切片的头部添加元素
	var b = []int{1,2,3}
	b = append([]int{0}, b...)        // 在开头添加1个元素
	b = append([]int{-3,-2,-1}, b...) // 在开头添加1个切片
	/*
		在头部一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，从切片的头部添加
	    元素的性能一般要比从尾部追加元素的性能差很多
	*/

	// 3.在切片的中间添加元素
	var c []int
	pos := 2                   // 表示要插入元素的位置
	val := 1                   // 表示要插入的元素
	val_slice := []int{1,2,3}  // 表示要插入的切片
	c = append(a[:pos], append([]int{val}, a[pos:]...)...)  // 在第pos个位置插入val
	c = append(c[:pos], append(val_slice, a[pos:]...)...)   // 在第pos个位置插入val_slice切片
	/*
		每个添加操作中的第二个append调用都会创建一个临时切片，并将a[pos:]的内容复制到新创建的切
	    片中，然后将临时创建的切片再追加到a[:pos]。
	*/

	// 4.用copy和append组合可以避免创建中间的临时切片，同样是完成添加元素的操作
	var d []int
	d = []int{1,2,3}
	d = append(d,0)      // 切片扩展1个空间
	copy(d[pos + 1:], d[pos:]) // d[pos:]向后移动1个位置
	d[pos] = val               // 设置新添加的元素


	// 5.用copy和append组合也可以实现在中间位置插入多个元素(也就是插入一个切片):
	var e []int
	e = []int{1,2,3}
	e = []int{1,2,3}
	e = append(e,val_slice...)              // 为val_slice切片扩展足够的空间
	copy(e[pos + len(val_slice):], e[pos:]) // e[pos:]向后移动1个位置
	copy(e[pos:],val_slice)                 // 复制新添加的切片
	/*
		稍显不足的是，在第一句扩展切片容量的时候，扩展空间部分的元素复制是没有必要的。没有专门的内置函数用于扩
	    展切片的容量，append本质是用于追加元素而不是扩展容量，扩展切片容量只是append的一个副作用。
	*/


// 给切片删除元素------
	// 1.通过移动数据指针在切片的尾部删除元素
	count := 2
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = a[:len(a)-1]       // 删除尾部1个元素
	a = a[:len(a)-count]   // 删除尾部count个元素

	// 2.通过移动数据指针在切片的头部删除元素
	a = a[1:]             // 删除开头1个元素
	a = a[count:]         // 删除开头count个元素

	// 3.用append将后面的数据向开头移动来删除头部元素，原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	a = append(a[:0], a[count:]...) // 删除开头count个元素

	// 4.用copy来删除头部的元素
	a = a[:copy(a, a[1:])]    // 删除开头1个元素
	a = a[:copy(a, a[count:])] //删除开头count个元素

	// 5.删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用append原地完成
	a = append(a[:pos],a[pos + 1:]...)     // 删除中间1个元素
	a = append(a[:pos],a[pos + count:]...) // 删除中间count个元素

	// 6.删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用copy原地完成
	a = a[:pos+copy(a[pos:], a[pos+1:])]      // 删除中间1个元素
	a = a[:pos+copy(a[pos:], a[pos+count:])]  // 删除中间count个元素

	/*
		删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况。
	*/
}