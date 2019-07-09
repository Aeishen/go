/*
@author : Aeishen
@data :  19/07/05, 14:35

@description : 空接口：一个不包含任何方法的接口，形如：interface{}
*/


package main

import "fmt"

// 定义一个结构体类型
type i_Struct struct {
	value int
}

// 定义一个string类型
type i_String string

// 接受接口类型参数，打印出该接口的动态类型与动态值
func explain(i interface{}) {
	fmt.Printf("type of s is %T\n", i)
	fmt.Printf("value of s is %v\n\n", i)
}

// 接受接口类型切片
func print(i ...interface{}){
	fmt.Println(i...)
}



func main() {

	//将 int 类型的变量当作接口，打印出该接口的动态类型与动态值
	var i_int int
	i_int = 10
	explain(i_int)

	//将 float64 类型的变量当作接口，打印出该接口的动态类型与动态值
	var i_float float64
	i_float = 10.002
	explain(i_float)

	//将结构体类型的变量当作接口，打印出该接口的动态类型与动态值
	var i_struct i_Struct
	i_struct.value = 100
	explain(i_struct)

	//将自定义 string 类型的变量当作接口，打印出该接口的动态类型与动态值
	var i_string i_String
	i_string = "i am string type"
	explain(i_string)

	// 当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果
	var a = []interface{}{123, "abc"}
	print(a...) // 解包，输出123 abc，等价于直接调用Print(123, "abc")
	print(a)    // 未解包，输出[123 abc]，等价于直接调用Print([]interface{}{123, "abc"})

}

/*
	因为空接口不包含任何方法，所以任何类型都默认实现了空接口

    再举个例子：fmt 包中的 Println() 函数
        func Println(a ...interface{}) (n int, err error) {}
		该函数可以接收多种类型的值，比如：int、string、array等。为什么，因为它的形参就是空接口，可以接收任意类型的值。
*/