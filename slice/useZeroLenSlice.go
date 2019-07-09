/*
@author : Aeishen
@data :  19/07/09, 14:11

@description : 使用0长切片
@from
*/

package main

import "fmt"

// 用于删除[]byte中的空格
func TrimSpace(s []byte) []byte {
	b := s[:0]
	fmt.Println(len(b),cap(b),b[1:len(s)])
	/*
	s 是传入的原切片的一个拷贝，s 其引用的底层数组与原切片相同
	由于b = s[:0]， b 是 s 直接截取来的，与 s、原切片引用相同的底层数组，
	所以 b 或者 s 或许 原切片对底层数组的操作都会影响到彼此
	1.由于截取范围是[:0],所以len(b) = 0
	2.引用相同的底层数组,所以cap(b) = cap(s)，
	3.当对 b 进行截取时候其实也是获取底层数组
	*/

	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

//类似的根据过滤条件原地删除切片元素的算法都可以采用类似的方式处理（因为是删除操作不会出现内存不足的情形）：
func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func main() {
	s := []byte{0,' ',5,3,' ','j','y'}
	fmt.Println(len(s),cap(s),s[1:])

	b1 := TrimSpace(s)
	fmt.Println(len(b1),cap(b1),b1)
	fmt.Println(len(s),cap(s),s)
}

/*
切片高效操作的要点是要降低内存分配的次数，
尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。
*/