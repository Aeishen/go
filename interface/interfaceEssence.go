/*
@author : Aeishen
@data :  19/07/05, 09:14

@description :接口的本质（即接口的内在表现）一个接口可以被认为是由一个元组（类型，值）在内部表示的。type是接口的基础具体类型（动态类型/运行时类型），value是具体类型的值（动态值/运行时的值）。
*/

package main

import "fmt"

type testInterface interface{
	printSelfValue()
}

// 定义一个结构体
type intStruct struct {
	value int
}

// 定义一个类型
type floatType float64

func (this intStruct)printSelfValue(){
	fmt.Println(this.value)
}


func (this floatType)printSelfValue(){
	fmt.Println(this)
}


//获取接口的类型与值
func getEssenxe(this testInterface){
	fmt.Printf("该接口的动态类型是 %T, 动态值是 %v\n",this,this)
}


//接口判空
func isEmtry(t testInterface,j *intStruct){
	if j == nil {
		fmt.Println("结构体 j 是nil")
	}else{
		fmt.Println("结构体 j 不是nil")
	}
	if  t == nil {
		fmt.Println("接口 t 是nil")
	}else{
		fmt.Println("接口 t 不是nil")
		fmt.Printf("接口 t 的 type 是 %T,接口 t 的 value 是 %v\n",t,t)
	}
}


func main() {

	fmt.Println("\n1.将结构体赋值给接口，再获取接口类型与值------------------------\n")
	var t testInterface
	i := intStruct{1}
	t = i
	getEssenxe(t)

	f := floatType(3.33)
	t = f
	getEssenxe(t)

	/*
		变量的类型在声明时指定、且不能改变，称为静态类型。接口类型的静态类型就是接口本身。接口没有静态值，
		它指向的是动态值。接口类型的变量存的是实现接口的类型的值。该值就是接口的动态值，实现接口的类型就
		是接口的动态类型。

		在上面例子中接口变量 t 的静态类型是 testInterface ，是不能改变的。动态类型却是不固定的，第一
		次将 i 赋值给 t 后，t 的动态类型是 intStruct，再将 f 赋值给 t 后，t 的动态类型是 floatType

		有时候，接口的动态类型又称为具体类型，当我们访问接口类型的时候，返回的是底层动态值的类型。
    */

	fmt.Println("\n2.将一个空结构体赋值给接口再对接口判空------------------------\n")
	var j *intStruct
	//将 *intStruct 赋值给接口，接口动态类型为 *intStruct，动态值为 nil
	t = j
	isEmtry(t ,j)


	fmt.Println("\n3.直接对一个没有赋值的接口判空------------------------\n")
	var t1 testInterface
	//未将任何类型赋值给接口，接口动态值和动态类型都为 nil
	isEmtry(t1 ,j)

	fmt.Println("\n4.将一个没有赋值的接口 t3 赋值给相同的一个没有赋值的接口 t2 ，再对 t2 判空------------------------\n")
	var t2 testInterface
	var t3 testInterface
	t2 = t3
	//运行时 t2 的类型没有改变（即没有动态类型），还是nil，所以结果为true
	fmt.Println(t2 == nil)

	/*
		从输出可知：
	    当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil，此时的接口成为 nil 接口
	*/
}


/*
     这是 go 语言的规范
	 var x interface{}  // x is nil and has static type interface{}  ：x是 nil 并且具有静态类型接口{}
	 var v *T           // v has value nil, static type *T           ：v的值为 nil，静态类型 *T
	 x = 42             // x has value 42 and dynamic type int       ：x的值为 42，动态类型为int
	 x = v              // x has value (*T)(nil) and dynamic type *T ：x具有值 (*T)(nil) 和动态类型*T
*/



