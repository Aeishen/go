/*
@author : Aeishen
@data :  19/07/11, 16:03

@description : 捕获异常
@from :  Go语言高级编程(Advanced Go Programming), 柴树杉 曹春晖/著
*/

/*
	Go语言函数调用的正常流程是函数执行返回语句返回结果，在这个流程中是没有异常的，因此在这个流程中执行recover
	异常捕获函数始终是返回nil。另一种是异常流程: 当函数调用panic抛出异常，函数将停止执行后续的普通语句，但是
	之前注册的defer函数调用仍然保证会被正常执行，然后再返回到调用者。对于当前函数的调用者，因为处理异常状态还
	没有被捕获，和直接调用panic函数的行为类似。在异常发生时，如果在defer中执行recover调用，它可以捕获触发
	panic时的参数，并且恢复到正常的执行流程。
*/

package main

import (
	"log"
)

// recover函数的包装函数
func MyRecover() interface{} {
	log.Println("trace")
	return recover()
}

// （不可行的）例子1：非defer语句中执行recover
func test_1(){
	if p:=recover(); p != nil{      // 捕获结果为nil
		log.Fatal(p)
	}

	panic(111)

	if p:=recover(); p != nil{      // 捕获不到异常
		log.Fatal(p)
	}
}

// （不可行的）例子2：defer中调用的是recover函数的包装函数
func test_2(){
	defer func() {
		if p:= MyRecover(); p != nil{     // 捕获不到异常
			log.Fatal(p)
		}
	}()
	panic(111)
}

// （不可行的）例子3：在嵌套的defer函数中调用recover
func test_3(){
	defer func(){
		defer func() {
			if p:= recover(); p != nil{    // 捕获不到异常
				log.Fatal(p)
			}
		}()
	}()
	panic(111)
}

// （可行的）例子4：defer语句中调用MyRecover函数
func test_4(){
	defer MyRecover()  // 可以正常捕获异常
	panic(111)
}

// （不可行的）例子5：defer语句直接调用recover函数
func test_5(){
	defer recover()  // 捕获不到异常
	panic(111)
}

// （不可行的）例子6：避免recover调用者不能识别捕获到的异常, 应该避免用nil为参数抛出异常
func test_6(){
	defer func() {
		if P := recover(); P != nil {
			log.Fatal(P)     // 可以正常捕获异常，但是打印不出任何信息
		}
	}()
	panic(nil)
}

func main() {
	//test_1()
	/*
		例子1两个recover调用都不能捕获任何异常。在第一个recover调用执行时，函数必然是在正常的非异常执行
		流程中，这时候recover调用将返回nil。发生异常时，第二个recover调用将没有机会被执行到，因为panic调用
		会导致函数马上执行已经注册defer的函数后返回
    */


	//test_2()
	//test_3()
	/*
		例子2和例子3 都是经过了2个函数帧才到达真正的recover函数，这个时候Goroutine的对应上一级栈帧中
		已经没有异常信息。
    */

	//test_4()
	//test_5()
	/*
		必须要和有异常的栈帧只隔一个栈帧，recover函数才能正常捕获异常
	*/

	//test_6()

}


