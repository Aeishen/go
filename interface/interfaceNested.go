/*
@author : Aeishen
@data :  19/07/08, 11:50

@description : 接口嵌套: 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。
*/

package main

import "fmt"

// 定义一个相机接口，拥有拍摄照片的行为
type camera2 interface {
    takePhoto()
}

// 定义一个播放器接口，拥有播放音乐的行为
type player2 interface {
	playMusic()
}

// 定义一个手机接口，这个接口由 camera2 和 player2 两个接口嵌入。也就是说，phone2 同时拥有了 camera2 和 player2 的所有行为。它还拥有自己的呼叫行为
type phone2 interface {
	camera2
	player2
	call()
}

// 定义一个智能手机结构体
type smartPhone2 struct {
	name string
}

// 实现拍照行为
func (this smartPhone2)takePhoto(){
	fmt.Println(this.name,"can take photoes")
}

// 实现播放行为
func (this smartPhone2)playMusic(){
	fmt.Println(this.name,"can play music")
}

// 实现呼叫行为
func (this smartPhone2)call(){
	fmt.Println(this.name,"can call others")
}


func main() {
	oppoPhone := smartPhone2{"oppo"}

	// 结构体直接调用嵌套的接口的方法
	oppoPhone.takePhoto()
	oppoPhone.playMusic()
	oppoPhone.call()

	// 结构体转换为接口调用,以下所有接口运行时的类型都是 smartPhone2 ，而 smartPhone2 实现了接口所有的方法
	fmt.Println("\n测试oppoPhone的类型-----------------------------")
	var oppoPhone1 camera2 = oppoPhone
	if c,ok := oppoPhone1.(camera2); ok{
		fmt.Printf("oppoPhone can be a camera and it's type is %T now\n",c)
		c.takePhoto()
	}

	var oppoPhone2 player2 = oppoPhone
	if c,ok := oppoPhone2.(player2); ok{
		fmt.Printf("oppoPhone can be a player and it's type is %T now\n",c)
		c.playMusic()
	}

	var oppoPhone3 phone2 = oppoPhone
	if c,ok := oppoPhone3.(phone2); ok{
		fmt.Printf("oppoPhone can be a phone and it's type is %T now\n",c)
		c.playMusic()
		c.takePhoto()
		c.call()
	}

	var oppoPhone4 camera2 = oppoPhone
	if c,ok := oppoPhone4.(phone2); ok{
		fmt.Printf("oppoPhone can be a phone and it's type is %T now\n",c)
		c.playMusic()
		c.takePhoto()
		c.call()
	}


}

