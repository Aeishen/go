/*
@author : Aeishen
@data :  19/07/04, 23:12

@description :接口定义
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

func main() {

	// 定义一个Phone类型变量
	var phone phone

	hwPhone := huaWeiPhone{"HuaWei"}
	hwPhone.call()
	phone = hwPhone
	phone.call()

	xmPhone := xiaoMiPhone{"XiaoMi"}
	xmPhone.call()
	phone = xmPhone
	phone.call()
}

/*
	go中的接口是一组方法的集合，但不包含方法的实现、是抽象的，接口中也不能包含变量，任何其他类型只要
    实现了这些方法就是实现了这个接口，那么这些类型的变量也是这个接口类型的变量，不一定非要显式地声明
    要去实现哪些接口啦。
    上面例子phone可以直接使用 . 语法调用 Area() 方法，因为 phone 的具体类型是 hwPhone/xmPhone，
    而 hwPhone/xmPhone 实现了 call() 方法。XiaoMiPhone 和 HuaWeiPhone 都实现了Phone接口的
    call()方法，所以它们都是Phone
*/
