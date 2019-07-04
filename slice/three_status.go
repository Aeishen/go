/*
@author : Aeishen
@data :  19/07/04, 12:20

@description : 切片的三种状态：空切片，零切片，nil切片
*/

/*
切片的结构（一个包含三个变量的结构体）：
	type slice struct {
		array unsafe.Pointer   //指向底层数组
		length int             //长度
		capcity int            //容量
	}
*/

package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main(){

	fmt.Println("\n三种状态的切片创建：---------------------")

	//零切片:表示底层的数组的二进制内容都是零，如果是一个指针类型的切片则底层数组的内容全是nil
	var zero_slice_1 = make([]int,10)
	fmt.Println("zero_slice_1 :",zero_slice_1,len(zero_slice_1),cap(zero_slice_1))

	var zero_slice_2 = make([]*int,10)
	fmt.Println("zero_slice_2 :",zero_slice_2,len(zero_slice_2),cap(zero_slice_2))

	//空切片
	var empty_slice_1 = []int{}
	fmt.Println("empty_slice_1:",empty_slice_1,len(empty_slice_1),cap(empty_slice_1))

	var empty_slice_2 = make([]int,0)
	fmt.Println("empty_slice_2:",empty_slice_2,len(empty_slice_2),cap(empty_slice_2))

	//nil切片
	var nil_slice_1 []int
	fmt.Println("nil_slice_1:",nil_slice_1,len(nil_slice_1),cap(nil_slice_1))

	var nil_slice_2 = *new([]int)
	fmt.Println("nil_slice_2:",nil_slice_2,len(nil_slice_2),cap(nil_slice_2))


	/*
	   以上空切片与nil切片打印出来结果都是一样的，但其实内部是不同的，我们通过unsafe.Pointer
	   来转换Go语言的任意变量类型。因为切片的内部结构是一个结构体包含三个整型变量，其中第一个变量
	   是一个指针变量，指针变量里面存储的也是一个整型值，只不过这个值是另一个变量的内存地址。我们
	   可以将这个结构体看成长度为3的整型数组 [3]int。然后将切片变量转换成 [3]int
	*/
	fmt.Println("\n获取空切片与nil切片的内部结构：-----------------")

	var real_empty_slice_1 = *(*[3]int)(unsafe.Pointer(&empty_slice_1))
	var real_empty_slice_2 = *(*[3]int)(unsafe.Pointer(&empty_slice_2))
	var real_nil_slice_1 = *(*[3]int)(unsafe.Pointer(&nil_slice_1))
	var real_nil_slice_2 = *(*[3]int)(unsafe.Pointer(&nil_slice_2))

	fmt.Println("real_empty_slice_1:",real_empty_slice_1)
	fmt.Println("real_empty_slice_2:",real_empty_slice_2)
	fmt.Println("real_nil_slice_1:",real_nil_slice_1)
	fmt.Println("real_nil_slice_2:",real_nil_slice_2)


	/*
		从输出可以看出所以空切片的指向底层数组的变量是一个特殊的内存地址，即所有的空切片共享这个内存地址
	    该内存地址在 go 源码中被定义为一个叫 zerobase 的 uintptr（指向任意类型的指针）变量，源码给
	    的注释：“base address for all 0-byte allocations” 即所有0字节分配的基本地址
	*/

	fmt.Println("\n比较空切片与nil切片的区别：-------------------")
	fmt.Println("empty_slice_1 is nil ?",empty_slice_1 == nil)
	fmt.Printf("empty_slice_1 value: %#v\n",empty_slice_1)

	fmt.Println("nil_slice_1 is nil ?",nil_slice_1 == nil)
	fmt.Printf("nil_slice_1 value: %#v\n",nil_slice_1)

    /*
    	从输出结果可以看出空切片并不是nil,理由如上（空切片的指向底层数组的变量是一个特殊的内存地址）
        官方建议最好不要使用空切片，统一使用 nil 切片，同时要避免将切片和 nil 进行比较来执行某些逻辑
    */

    /*
    	空切片和 nil 切片也可能隐藏在结构体中,如下
    */

    type Eg struct {
    	Values []int
	}

    var eg_1 = Eg{}
    var eg_2 = Eg{[]int{}}
		fmt.Println("eg_1 value is nil ?",eg_1.Values == nil)
	fmt.Printf("eg_1 value: %#v\n",eg_1)

	fmt.Println("eg_2 value is nil ?",eg_2.Values == nil)
	fmt.Printf("eg_2 : %#v\n",eg_2)


	/*
		「空切片」和「 nil 切片」还有一个极为不同的地方在于 JSON 序列化，如下
	*/

	json_eg_1, _ := json.Marshal(eg_1)
	json_eg_2, _ := json.Marshal(eg_2)
	fmt.Println("json eg_1:",string(json_eg_1))
	fmt.Println("json eg_2:",string(json_eg_2))

}
