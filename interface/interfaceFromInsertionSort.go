/*
作者：咔叽咔叽
链接：https://juejin.im/post/5de54367e51d457c650df59a
来源：掘金
*/

package main

import "fmt"

func insertionSort(arr []int)  {
	sta,en := 0,len(arr)
	for i := sta + 1; i < en ;i ++{
		for j := i; j > sta && arr[j] < arr[j - 1]; j--{
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
	}
}

/*
问题：都知道 Go 是静态语言，那么就意味着不同的数据类型可能导致上述的插入排序不可用。比如说，某天产品想要支持 uint32 的插入排序。
     嗯，很简单，直接 Ctrl+c + Ctrl+c 稍微修改一下。
	func insertionSortUint32(arr []uint32) {
		sta,en := 0,len(arr)
		for i := sta + 1; i < en ;i ++{
			for j := i; j > sta && arr[j] < arr[j - 1]; j--{
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
    谁知道哪天产品脑子又抽风，他想要支持 float32 类型的插入排序，代码可能又得加几行。
	func insertionSortFloat32(arr []float32) {
		sta,en := 0,len(arr)
		for i := sta + 1; i < en ;i ++{
			for j := i; j > sta && arr[j] < arr[j - 1]; j--{
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}
		}
	}
    好像还看得下去，我们知道 Go 中的类型可不止这 3 种，再这么被搞下去是不是要爆炸了？

解决：
	首先，回到上诉的三个类型的排序中来，我们可以发现这几个排序除了数据类型是基本一致的。如果我们想用一个函数来支持所有的数据类型，
    我们是不是只能使用 interface来实现这个功能？但是 interface 又不支持运算操作，如果断言出来，还是跟以前一样麻烦。我们看看代码
    中需要对数据进行运算操作的地方：
	1.len(arr)
	2.arr[j] < arr[j - 1]
    3.arr[j], arr[j - 1] = arr[j - 1], arr[j]
	发现排序中只有len(arr)、arr[j] < arr[j - 1]、arr[j], arr[j - 1] = arr[j - 1], arr[j]这三种操作 interface 不支持。
    如果我们让 interface 实现这三个方法不就解决了我的问题了吗？接下来我们通过这种思路修改一下我们的插入排序
*/

type sortTyre interface {
	Len()int
	Swap(i,j int)
	Less(i,j int)bool
}

func insertionSortForInterface(arr sortTyre)  {
	sta,en := 0,arr.Len()-1
	for i := sta + 1; i < en ;i ++{
		for j := i; j > sta && arr.Less(j,j - 1); j--{
			arr.Swap(j,j - 1)
		}
	}
}

//我们使用了interface来替代写死的数据类型。如果调用方使用，只要实现 Data 接口就行了。
type Uint32Slice []uint32
func (u Uint32Slice) Len() int {return len(u)}
func (u Uint32Slice) Less(i, j int) bool {return u[i] < u[j]}
func (u Uint32Slice) Swap(i, j int) {u[i], u[j] = u[j], u[i]}

type Float32Slice []float32
func (u Float32Slice) Len() int {return len(u)}
func (u Float32Slice) Less(i, j int) bool {return u[i] < u[j]}
func (u Float32Slice) Swap(i, j int) {u[i], u[j] = u[j], u[i]}


func main() {
	arr := []int{2, 3, 4, 1, 7, 9, 10, 21, 17}
	insertionSort(arr)
	fmt.Println(arr)

	nums := Uint32Slice{2, 3, 4, 1, 7, 9, 10, 21, 17}
	insertionSortForInterface(nums)
	fmt.Println(nums)

	float32Nums := Float32Slice{2, 3, 4, 1, 7, 9, 10, 21, 17}
	insertionSortForInterface(float32Nums)
	fmt.Println(float32Nums)
}

/*
总结：我们通过接口实现了一个支持多种数据类型的插入排序，调用者只需要实现 Data 这个接口就可以使用了，而不用去修改插入排序原有的函数定义。
     这样使得我们的代码抽象度更高也更灵活，当我们面临类似的需求时，接口就是答案。官方sort包针对不同类型就是使用接口实现
*/
