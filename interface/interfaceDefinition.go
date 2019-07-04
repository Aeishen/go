/*
@author : Aeishen
@data :  19/07/04, 23:12

@description :接口定义
*/

package main

import "fmt"

//定义一个手机接口
type Phone interface {
	call() // 拥有呼叫的行为
}

// 华为手机
type HuaWeiPhone struct {
	name string
}

// 小米手机
type XiaoMiPhone struct {
	name string
}

// 华为手机 实现call
func (this HuaWeiPhone) call() {
	fmt.Println(this.name, " is calling")
}

// 小米手机 实现call
func (this XiaoMiPhone) call() {
	fmt.Println(this.name, " is calling")
}

func main() {

	// 定义一个Phone类型变量
	var phone Phone

	hwPhone := HuaWeiPhone{"HuaWei"}
	hwPhone.call()
	phone = hwPhone
	phone.call()

	xmPhone := XiaoMiPhone{"XiaoMi"}
	xmPhone.call()
	phone = xmPhone
	phone.call()
}

/*
	go中的接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口，
    那么这些类型的变量也是这个接口类型的变量，不一定非要显式地声明要去实现哪些接口啦。在上面的例子中，
    XiaoMiPhone和HuaWeiPhone都实现了Phone接口的call()方法，所以它们都是Phone
*/
