/*
@author : Aeishen
@data :  19/07/05, 16:03

@description : 接口断言x.(T)：这里x表示一个接口的类型，T表示一个类型（也可为接口类型）。一个类型断言就是检查一个接口对象x的动态（运行时的）类型是否和断言的类型T匹配。
*/

package main

import "fmt"

// 定义一个播放器接口，拥有播放音乐的行为
type player1 interface {
	playMusic()
}

// 定义一个相机接口，拥有拍摄照片的行为
type camera1 interface {
	takePhoto()
}

// 定义一个游戏机接口，拥有玩游戏的行为
type playStation interface {
	playGame()
}

// 定义智能手机结构体
type smartPhone1 struct {
	name  string
}


// smartPhone1实现播放音乐的行为，那么它也是一个player1
func (this smartPhone1)playMusic()  {
	fmt.Println(this.name ,"can play music, it is also a player")
}

// smartPhone1实现拍摄照片的行为，那么它也是一个camera1
func (this smartPhone1)takePhoto()  {
	fmt.Println(this.name ,"can take photos, it is also a camera")
}


func main() {
	var p player1 = smartPhone1{"XiaoMiPhone"}


	fmt.Println("1断言的第一种情况：断言的类型T是一个具体类型-------\n")

	// s  := p.(smartPhone1)    //直接断言（如果检查失败，则会抛出panic)，防止panic,用两个变量来接收检查结果,ok == true 代表断言成功
	s ,ok := p.(smartPhone1)
	if ok {
		fmt.Printf("s 运行时的类型：%T, 运行时的值：%v\n",s,s)
		s.takePhoto()
		s.playMusic()
	}

	fmt.Println("\n\n2断言的第二种情况：断言的类型T是一个接口类型-------\n")
	fmt.Println("①断言的接口类型未实现-------\n")
	s1 ,ok1 := p.(playStation)   //smartPhone1 没有实现playGame的行为，那么它不是一个playStation
	if ok1{
		fmt.Printf("s1 运行时的类型：%T, 运行时的值：%v\n",s1,s1)
		s1.playGame()
	}else{
		fmt.Printf("s1 运行时的类型：%T, 运行时的值：%v\n",s1,s1)
		fmt.Println("smartPhone1 不是一个 playStation \n")
	}

	fmt.Println("②断言的接口类型已实现-------\n")
	s2 ,ok2 := p.(camera1)

	if ok2{
		fmt.Printf("s2 运行时的类型：%T, 运行时的值：%v\n",s2,s2)
		s2.takePhoto()
	}


	/*
	断言的第一种情况：
		如果断言的类型T是一个具体类型，则断言 x.(T) 就检查 x 的动态类型是否和 T 的类型相同。
		若断言成功了，则返回一个类型为 T 的对象，该对象的值为接口变量 x 的动态值。换句话说，
		具体类型的类型断言从它的操作对象中获得具体的值。若断言失败了，同样返回一个类型为 T 的
	    对象，但该对象的值为类型 T 的零值

	断言的第二种情况：
		如果断言的类型 T 是一个接口类型，则断言 x.(T) 检查 x 的动态类型是否满足 T 接口（即实现了T接口中的所有方法）
	    若断言成功了,返回一个类型为 T 的接口对象，该对象的值为接口变量 x 的动态值,若断言失败了，
	    同样返回一个类型为 T 的接口对象，但该对象的值为类型 T 的零值 (nil)
	*/
}
