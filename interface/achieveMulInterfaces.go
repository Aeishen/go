/*
@author : Aeishen
@data :  19/07/05, 15:15

@description : 实现多接口
*/


package main

import "fmt"

// 定义一个播放器接口，拥有播放音乐的行为
type player interface {
	playMusic()
}

// 定义一个相机接口，拥有拍摄照片的行为
type camera interface {
	takePhoto()
}

// 定义智能手机结构体
type smartPhone struct {
	name  string
}


// 智能手机结构体实现播放音乐的行为，那么它也是一个播放器
func (this smartPhone)playMusic()  {
	fmt.Println(this.name ,"can play music, it is also a player")
}

// 智能手机结构体实现拍摄照片的行为，那么它也是一个相机
func (this smartPhone)takePhoto()  {
	fmt.Println(this.name ,"can take photos, it is also a camera")
}

func main() {

	//创建一个smartPhone对象，赋值给camera接口与player接口，因为smartPhone结构体实现了camera接口与player接口，所以smartPhone相当于camera或者player
	phone := smartPhone{"HuaWeiPhone"}
	var p player = phone
	var c camera = phone

    //此时 p 和 c 具有相同的动态类型和动态值，分别调用各自实现的方法 playMusic() 和 takePhoto()
	p.playMusic()
	c.takePhoto()


	/*
	    使用如下将会报错 c.playMusic undefined (type camera has no field or method playMusic) p.takePhoto undefined (type player has no field or method takePhoto)
		c.playMusic()
		p.takePhoto()
	    因为c的静态类型是camera，其不包含playMusic行为，p的静态类型是player，其不包含takePhoto行为
	    但是使用断言可以在某些程度上解决这个问题，移步interfaceAssert.go
	*/
}