/*
@author : Aeishen
@data :  19/07/09, 14:40

@description : 避免切片内存泄漏
@from
*/

/*
切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。
但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这
会延迟自动内存回收器对底层数组的回收。
*/

package main

func main() {
	a1,a2,a3,a4 := 1,2,3,4
	var a []*int
	a = []*int{&a1,&a2,&a3,&a4}

	// 被删除的最后一个元素依然被引用, 可能导致GC操作被阻碍
	a = a[:len(a)-1]

	//改进：先将需要自动内存回收的元素设置为nil，保证自动回收器可以发现需要回收的对象，然后再进行切片的删除操作
	a[len(a)-1] = nil  // GC回收最后一个元素内存
	a = a[:len(a)-1]   // 从切片删除最后一个元素

	/*
	当然，如果切片存在的周期很短的话，可以不用刻意处理这个问题。因为如果切片本身已经可以被GC回收的话，切片对应
	的每个元素自然也就是可以被回收的了。
	*/
}

/*
再例如：
    FindPhoneNumber函数加载整个文件到内存，然后搜索第一个出现的电话号码，最后结果以切片方式返回。

	func FindPhoneNumber(filename string) []byte {
		b, _ := ioutil.ReadFile(filename)
        return regexp.MustCompile("[0-9]+").Find(b)
    }
    这段代码返回的[]byte指向保存整个文件的数组。因为切片引用了整个原始数组，导致自动垃圾回收器不能
    及时释放底层数组的空间。一个小的需求可能导致需要长时间保存整个文件数据。这虽然这并不是传统意义上
    的内存泄漏，但是可能会拖慢系统的整体性能。

    要修复这个问题，可以将感兴趣的数据复制到一个新的切片中（数据的传值是Go语言编程的一个哲学，虽然传
    值有一定的代价，但是换取的好处是切断了对原始数据的依赖）：

	func FindPhoneNumber(filename string) []byte {
		b, _ := ioutil.ReadFile(filename)
		b = regexp.MustCompile("[0-9]+").Find(b)
		return append([]byte{}, b...)
	}
}
*/