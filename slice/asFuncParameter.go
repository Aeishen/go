/*
@author : Aeishen
@data :  19/07/08, 18:01

@description : 切片作为函数参数
*/

package main

import "fmt"

// 尝试给切片赋新值
func changeSlice_1(s []int){
	s = []int{1,3,5}
}

// 尝试改变切片第一个元素
func changeSlice_2(s []int){
	s[0] = 1
}


// 尝试交换切片元素
func changeSlice_3(s []int){
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// 尝试给切片追加一个元素后再交换切片元素
func changeSlice_4(s []int){
	s = append(s,999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// 尝试给切片追加多个元素后再交换切片元素
func changeSlice_5(s []int){
	s = append(s,999,888,777,666)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	slice1 := []int{0,1,2}
	changeSlice_1(slice1)
	fmt.Println(slice1)
	/*
	1.输出[0, 1, 2]。 因为是值拷贝传递，changeSlice_1中的 s 是 main 里的 slice1 的一个拷贝，changeSlice_1 将一个新的切片赋予s,
	  从而 s 指向一个新的底层数组
	*/

	slice2 := []int{0,1,2}
	changeSlice_2(slice2)
	fmt.Println(slice2)
	/*
	2.输出[1, 1, 2]。changeSlice_2 中的 s 是 main 里的 slice2 的一个拷贝，但没有给 changeSlice_2 中的 s 新赋值，所以s 与 slice2
	  指向相同的底层数组，在 changeSlice_2 中的修改 s 指向的底层数组的值也会影响到 slice2
	*/

	slice3 := []int{0,1,2}
	changeSlice_3(slice3)
	fmt.Println(slice3)
	/*
	3.输出[2, 1, 0]。changeSlice_3 中的 s 是 main 里的 slice3 的一个拷贝，但没有给 changeSlice_3 中的 s 新赋值，所以s 与 slice3
	  指向相同的底层数组，在 changeSlice_3 中的修改 s 指向的底层数组的值也会影响到 slice3
	*/

	slice4 := []int{0,1,2}
	fmt.Println(len(slice4),cap(slice4)) // 输出 3,3
	changeSlice_4(slice4)
	fmt.Println(slice4)
	/*
	4.输出[0, 1, 2]。changeSlice_4 中的 s 是 main 里的 slice4 的一个拷贝，但没有给 changeSlice_4 中的 s 新赋值，所以s 与 slice4
	  指向相同的底层数组，在 changeSlice_4 中给 s 追加一个元素，此时由于 s 指向的底层元素长度与容量均为3，若追加一个元素，超过了底层数组容量
	  则分配一个新的比原来数组容量大的数组给s，该新数组复制了原来数组中的所有元素，给 s 追加一个元素将添加到该新数组中，同时 s 指向该新数组，现
	  在与 slice4 指向不同地址了，所以不会影响到slice4
	*/

	slice5 := []int{0,1,2}
	fmt.Println(len(slice5),cap(slice5)) // 输出 3,3
	changeSlice_5(slice5)
	fmt.Println(slice5)
	/*
	5.同理于4
	*/

	var slice6 []int  // nil切片
	fmt.Println(len(slice6),cap(slice6)) // 输出 0,0
	changeSlice_4(slice6)
	fmt.Println(slice6)
	/*
	6.同理于4
	*/

	var slice7 = []int{}  // 空切片
	fmt.Println(len(slice7),cap(slice7)) // 输出 0,0
	changeSlice_4(slice7)
	fmt.Println(slice7)
	/*
	7.同理于4
	*/

	var slice8 = make([]int,3) // 零切片
	fmt.Println(len(slice8),cap(slice8)) // 输出 3,3
	changeSlice_4(slice8)
	fmt.Println(slice8)
	/*
	8.同理于4
	*/

	var slice9 []int
	for i := 1; i <= 3; i++ {
		slice9 = append(slice9, i)
	}
	fmt.Println(len(slice9),cap(slice9)) // 输出 3,4 因为对于空切片的每进行一次添加都会扩容（自行百度go切片容量扩充）
	changeSlice_4(slice9)
	fmt.Println(slice9)
	/*
	9.输出[999, 3, 2]。changeSlice_4 中的 s 是 main 里的 slice9 的一个拷贝，但没有给 changeSlice_4 中的 s 新赋值，所以s 与 slice9
	  指向相同的底层数组，在 changeSlice_9 中给 s 追加一个元素，此时由于 s 指向的底层元素长度3，容量均为4，若追加一个元素，不超过了底层数组容量
	  所以底层数组增加一个新的元素999，同时对该底层数组进行元素交换操作后底层为[999,3,2,1]。s 和 slice9 指向同一底层数组，由于 slice9 的len
	  为3，所以输出的 slice9 为[999, 3, 2]
	*/

}
/*
	因为切片中的底层数组部分是通过隐式指针传递(指针本身依然是传值的，但是指针指向的却是同一份的数据)，所以被调用函数是可以通过指针修改掉调用参数切片中
	的数据。除了数据之外，切片结构还包含了切片长度和切片容量信息，这2个信息也是传值的。如果被调用函数中修改了Len或Cap信息的话，就无法反映到调用参数的
	切片中，这时候我们一般会通过返回修改后的切片来更新之前的切片。这也是为何内置的append必须要返回一个切片的原因。
*/