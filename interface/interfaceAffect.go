/*
@author : Aeishen
@data :  19/07/05, 22:34

@description :接口的实际使用（即接口作用的体现） 不同的结构体调用同一个接口，指定了这些结构体拥有相同的行为，但行为内在表现可以由结构体自己决定
*/

package main

import "fmt"

//定义一个手机接口
type newphone interface {
	call() // 拥有呼叫的行为
	sales() int //  拥有出售的行为
}

// vivo手机
type vivoPhone struct {
	name string
	price int
}

// 苹果手机
type applePhone struct {
	name string
	price int
	duty int   // 税
}

// 华为手机 实现call
func (this vivoPhone) call() {
	fmt.Println(this.name, " is calling")
}

// 华为手机 实现sales
func (this vivoPhone) sales() int{
	 return this.price
}

// 苹果手机 实现call
func (this applePhone) call() {
	fmt.Println(this.name, " is calling")
}

// 苹果手机 实现sales
func (this applePhone) sales() int{
	return this.price + this.duty
}


// 计算所有手机的价格
func calPrices(allPhone []newphone) {
	allPrice := 0
	for _,onePhone := range allPhone{
		allPrice += onePhone.sales()
	}
	fmt.Println("所有手机的价格是：",allPrice)
}

func main() {

	// 定义一个Phone类型变量
	allNewPhone := []newphone{
		vivoPhone{"VIVO1",5000},
		vivoPhone{"VIVO2",4000},
		vivoPhone{"VIVO3",3000},
		applePhone{"APPLE",5000,200},
		applePhone{"APPLE",6000,200},
	}

	calPrices(allNewPhone)
}

