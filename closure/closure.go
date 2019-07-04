/*
@author : Aeishen
@data :  19/07/04, 15:19

@description : 测试 go 闭包
*/

package main

import "fmt"

func main(){

    // 正常情况下(defer逆序执行)
	for i := 0;i < 5;i ++{
		defer fmt.Println("Not in closure:",i)
	}

    // 闭包情况下(defer逆序执行)
	for i := 0;i < 5;i ++{
		defer func() {
			fmt.Printf("%v : %p\n",i,&i)
			fmt.Println("In closure :",i)
		}()
	}

	/*
		第2个for是外部变量在闭包内通过被改变了的情况，defer函数内的表达式会在它出现的地方就已经被求值，
		然后在退出函数之前，按照defer出现的顺序逆向执行。也就是说当第一条defer求值的时候，i=1,第五条
	    defer求值的时候，i=5，对于两个函数都是如此。区别在于第一个函数的i在每次defer是传值进printf
	    函数的，所以在defer中，i等于有5份拷贝，而第二个函数使用闭包的方式引用了外部变量i其实只有一份！
	*/

	/*
		要在闭包中避免上面的问题，可以有两种方式。
	    方法1: 每次循环构造一个临时变量 i
	    方法2: 通过函数参数传参
	*/

	for i := 0;i < 5;i ++{
		i := i
		defer func() {
			fmt.Println("after 1 In closure:",i)
		}()
	}

	for i := 0;i < 5;i ++{
		defer func(i int) {
			fmt.Println("after 2 In closure:",i)
		}(i)
	}

}
