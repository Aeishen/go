/*
@author : Aeishen
@data :  19/07/08, 10:27

@description : 接口类型判断: 用于将接口的具体类型与各种 case 语句中指定的多种类型进行匹配比较
*/



package main

import "fmt"

// 单接口类型判断，参数只能是接口类型
func switchSingleType(i interface{}){
	switch t := i.(type){
	case int :
		fmt.Printf("Type Int %T with value %v\n", t, t)
	case string :
		fmt.Printf("Type String %T with value %v\n", t, t)
	default:
		fmt.Println("unknowed type\n")
	}
}

// 多接口类型判断，参数只能是接口类型
func switchMulType(item ...interface{}){
	for _,i := range item{
		switch t := i.(type){
		case int :
			fmt.Printf("Type Int %T with value %v\n", t, t)
		case string :
			fmt.Printf("Type String %T with value %v\n", t, t)
		default:
			fmt.Println("unknowed type")
		}
	}
}

func main() {

	fmt.Println("使用单接口类型判断--------------------")
	var i_int int
	switchSingleType(i_int)

	var i_string  = "aaaa"
	switchSingleType(i_string)

	var i_bool  = false
	switchSingleType(i_bool)

	fmt.Println("使用多接口类型判断--------------------")
	var j_int = 1
	var j_string  = "bbbb"
	var j_float  = 0.221
	switchMulType(j_int, j_string, j_float, true)
}
/*
	本例中，i 的类型属于哪个 case ，就会执行 case 下相应的代码。
    但是只有接口类型才可以进行类型判断。其他类型，是不能的。因为
    int , bool , string , float 等等任何类型都默认实现了空接口 interface{}
    所以可以当作参数 i 传入
*/