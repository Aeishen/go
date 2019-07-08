/*
@author : Aeishen
@data :  19/07/04, 23:12

@description :接口定义/实现
*/

/*
    go中的接口是一组方法的集合，但不包含方法的实现、是抽象的，接口中也不能包含变量，任何其他类型只要
    实现了这些方法就是实现了这个接口，那么这些类型的变量也是这个接口类型的变量，不一定非要显式地声明
    要去实现哪些接口啦。
*/

package main

import "fmt"

//定义一个手机接口
type phone interface {
	call() // 拥有呼叫的行为
}

// 华为手机
type huaWeiPhone struct {
	name string
}

// 小米手机
type xiaoMiPhone struct {
	name string
}

// 华为手机 实现call
func (this huaWeiPhone) call() {
	fmt.Println(this.name, " is calling")
}

// 小米手机 实现call
func (this xiaoMiPhone) call() {
	fmt.Println(this.name, " is calling")
}


// 除了使用值接收者实现接口，也可以用指针接收者
type xxPhone struct {
	name string
}

// xxPhone的指针 实现call
func (this *xxPhone) call() {
	fmt.Println(this.name, " is calling")
}



func main() {

	// 定义一个Phone类型变量
	var phone phone

	// 值实现接口
	hwPhone := huaWeiPhone{"HuaWei"}
	hwPhone.call()
	phone = hwPhone   // 值调用（在go中，值接收者的方法可以使用值或者指针调用）
	phone.call()

	xmPhone := xiaoMiPhone{"XiaoMi"}
	xmPhone.call()
	phone = &xmPhone  // 指针调用（在go中，值接收者的方法可以使用值或者指针调用）
	phone.call()

	// 指针实现接口
	xx := xxPhone{"xx"}
	xx.call()
	phone = &xx      // 此处只能用指针调用（在go中，指针接收者的方法可以使用指针调用）
	phone.call()

}

/*
    上面例子phone可以直接使用 . 语法调用 call() 方法，因为 phone 的具体类型是 hwPhone/xmPhone，
    而 hwPhone/xmPhone 实现了 call() Phone接口的方法，所以 hwPhone/xmPhone 都是Phone

    本例也是利用接口实现“多态”
*/
