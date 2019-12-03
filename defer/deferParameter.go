/*
源作者：码农桃花源
链接：https://www.imooc.com/article/290861
来源：慕课网

defer 语句的参数
	1.作为函数参数，定义的时候就会求值，定义的时候变量的值是多少，最后该值是多少
	2.作为闭包引用的变量，在真正调用闭包时，才获取该变量的值。
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	paramFunc1()
	paramFunc2()
	paramFunc3()
}


func paramFunc1() {
	var err error

	fmt.Println(err)  //err作为函数参数

	err = errors.New("defer err1")
	return
}

func paramFunc2() {
	var err error

	defer func(err error) {
		fmt.Println(err)
	}(err)          //err作为函数参数

	err = errors.New("defer err3")
	return
}

func paramFunc3() {
	var err error

	defer func() {
		fmt.Println(err)  //闭包
	}()

	err = errors.New("defer err2")
	return
}

