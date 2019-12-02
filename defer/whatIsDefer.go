/*
源作者：码农桃花源
链接：https://www.imooc.com/article/290861
来源：慕课网

什么是defer?
	1.defer是Go语言提供的一种用于注册延迟调用的机制：让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行。
	2.defer语句通常用于一些成对操作的场景：打开连接/关闭连接；加锁/释放锁；打开文件/关闭文件等。
	3.defer在一些需要回收资源的场景非常有用，可以很方便地在函数结束前做一些清理操作


defer的简单使用：
	eg: f,err := os.Open(filename)
		  if err != nil{
			 panic(err)
		  }
		  if f != nil{
			 defer f.Close()
		  }
	ps:在打开文件的语句附近，用defer语句关闭文件。这样，在函数结束之前，会自动执行defer后面的语句来关闭文件。
	   当然，defer会有小小地延迟，对时间要求特别特别特别高的程序，可以避免使用它，其他一般忽略它带来的延迟。


defer底层原理：
	Each time a “defer” statement executes, the function value and parameters to the call are evaluated as usual and saved anew but
    the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns,
    in the reverse order they were deferred. If a deferred function value evaluates to nil, execution panics when the function is invoked,
    not when the “defer” statement is executed.
    每次defer语句执行的时候，会把函数“压栈”，函数参数会被拷贝下来；当外层函数（非代码块，如一个for循环）退出时，defer函数按照定义的逆序执行；
    如果defer执行的函数为nil, 那么会在最终调用函数的产生panic.

    解释：
    1.defer语句并不会马上执行，而是会进入一个栈，函数return前，会按先进后出的顺序执行。也说是说最先被定义的defer语句最后执行。
      先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行；否则，如果前面先执行，那后面函数的依赖就没有了。
    2.在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在defer定义时就把值传递给defer，
      并被cache起来；作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。
    3.defer后面的语句在执行的时候，函数调用的参数会被保存起来，也就是复制了一份。真正执行的时候，实际上用到的是这个复制的变量，
      因此如果此变量是一个“值”，那么就和定义的时候是一致的。如果此变量是一个“引用”，那么就可能和定义的时候不一致。

    例子test1,test2,test3,





*/

package main

import "fmt"

type number int

func main() {
	test1()
	test2()
	test3()
}

//defer后面跟的是一个闭包，i是“引用”类型的变量，最后i的值为2, 因此最后打印了三个2.
func test1(){
	var anyThing [3]int

	for i,_ := range anyThing{
		defer func() {
			fmt.Println("test1:",i)  // 直接用i
		}()
	}
}

//defer后面跟的是一个闭包，i是“引用”类型的变量，j 复制i的值作为函数参数，在defer定义时就把j值传递给defer, 因此最后打印了三个2,1,0.
func test2(){
	var anyThing [3]int

	for i,_ := range anyThing{
		j := i    // 复制i的值
		defer func(int) {
			fmt.Println("test2:",j)
		}(j)
		fmt.Println(&j,&i)
	}
}


func test3(){
	var n number

	//defer后面跟的是函数，参数是非“引用”类型，n 值为0已被缓存，对n直接求值，开始的时候n=0, 所以最后是0;
	defer n.test31()

	//defer后面跟的是函数, 参数是“引用”类型，n的引用被缓存,最后输出n的引用此时的指向的值为3
	defer n.test32()

	//defer后面跟的是一个闭包，调用的时候读取的是n此时的值为3
	defer func() { n.test31() }()

	//defer后面跟的是一个闭包，调用的时候读取的是n的引用此时的指向的值为3
	defer func() { n.test32()}()

	n = 3
}

func (n number) test31(){
	fmt.Println("test3_1:",n)
}

func (n *number) test32(){
	fmt.Println("test3_2:",*n)
}
