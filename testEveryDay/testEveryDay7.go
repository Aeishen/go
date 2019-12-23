/*
   @File: testEveryDay7
   @Author: Aeishen
   @Date: 2019/12/23 17:05
   @Description:
*/
package main

import "fmt"

func main() {
	testEveryDay7_2()
}


type MyInt1 int
type MyInt2 = int    //创建了 int 的类型别名 MyInt2，注意类型别名的定义时 =

//func testEveryDay7_1(){
//	var i int = 10
//	var i1 MyInt1 = i // 将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译不通过
//	fmt.Println(i1)
//}

func testEveryDay7_2(){
	var i int = 10
	var i2 MyInt2 = i  // MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	fmt.Println(i2)
}
