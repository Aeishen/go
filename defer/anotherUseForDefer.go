/*
来源于微信公众号网管叨bi叨 ，作者KevinYan11
*/
package main

import (
	"fmt"
	"log"
	"time"
)

/*
可以使用 defer 在任何函数开始后和结束前执行配对的代码。这个隐藏的功能在网上的教程和书籍中很少提到。要使用此功能，需要创建一个函数并
使它本身返回另一个函数，返回的函数将作为真正的延迟函数。在 defer 语句调用父函数后在其上添加额外的括号来延迟执行返回的子函数如下所示：
父函数返回的函数将是实际的延迟函数。父函数中的其他代码将在函数开始时（由 defer 语句放置的位置决定）立即执行。
*/
func main() {
	defer greet()()
	fmt.Println("Some code here...")
}

func greet() func() {
	fmt.Println("Hello!")
	return func() {
		fmt.Println("Bye!")
	} // this will be deferred
}

/*
这为开发者提供了什么能力？因为在函数内定义的匿名函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以
引用该函数的变量。在下一个示例中看到的，参数变量在 measure 函数第一次执行和其延迟执行的子函数内都能访问到：
*/

//func main() {
//	example()
//	otherExample()
//}

func example(){
	defer measure("example")()
	fmt.Println("Some code here")
}

func otherExample(){
	defer measure("otherExample")()
	fmt.Println("Some other code here")
}

func measure(name string) func() {
	start := time.Now()
	fmt.Printf("Starting function %s\n", name)
	return func(){
		fmt.Printf("Exiting function %s after %s\n", name, time.Since(start))
	}
}

/*
此外函数命名的返回值也是函数内的局部变量，所以上面例子中的 measure 函数如果接收命名返回值作为参数的话，那么命名返回值在延迟执行
的函数中也能访问到，这样就能将 measure 函数改造成记录入参和返回值的工具函数。

下面的示例是引用《go 语言程序设计》中的代码段:
可以想象，将代码延迟在函数的入口和出口使用是非常有用的功能，尤其是在调试代码的时候。
*/

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses   // ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg,time.Since(start))
	}
}