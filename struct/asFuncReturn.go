/*
@author : Aeishen
@data :  19/07/04, 18:13

@description : 结构体作为函数返回值
*/

package main

import "fmt"

type Dog struct {
	Name string
	Age int
}

//返回结构体值
func CreateValueDog() (dog Dog) {
	dog = Dog{"Tom",12}
	fmt.Printf(" valueDog address is :%p\n",&dog)
	return
}

//返回结构体指针
func CreatePointerDog() (dog *Dog) {
	dog = &Dog{"Amy",21}
	fmt.Printf(" PointerDog address is :%p\n",dog)
	return
}

func test_ValueDog(){
	fmt.Println("======测试结构体作为函数返回值，当返回值为结构体对象时：")
	valueDog1 := CreateValueDog()
	valueDog2 := CreateValueDog()
	fmt.Printf(" valueDog1  address is :%p\n valueDog2  address is :%p\n",&valueDog1,&valueDog2)
	fmt.Printf(" valueDog1 is :%v\n valueDog2 is :%v\n",valueDog1,valueDog2)
	valueDog1.Age = 15
	fmt.Println(" 改valueDog1的age为15后")
	fmt.Printf(" valueDog1 is :%v\n valueDog2 is :%v\n",valueDog1,valueDog2)
}

func test_PointerDog(){
	fmt.Println("======测试结构体作为函数返回值，当返回值为结构体指针时：")
	valueDog1 := CreatePointerDog()
	valueDog2 := CreatePointerDog()
	fmt.Printf(" valueDog1  address is :%p\n valueDog2  address is :%p\n",&valueDog1,&valueDog2)
	fmt.Printf(" valueDog1 is :%v\n valueDog2 is :%v\n",valueDog1,valueDog2)
	valueDog1.Age = 15
	fmt.Println(" 改valueDog1的age为15后")
	fmt.Printf(" valueDog1 is :%v\n valueDog2 is :%v\n",valueDog1,valueDog2)
}

func main() {
	test_ValueDog()
	test_PointerDog()
}