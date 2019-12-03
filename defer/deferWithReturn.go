/*
源作者：码农桃花源
链接：https://www.imooc.com/article/290861
来源：慕课网

defer 结合 return 的命令拆解
	return xxx
	上面这条语句经过编译之后，变成了三条指令：1. 返回值 = xxx2. 调用defer函数3. 空的return
    1,3步才是Return 语句真正的命令，第2步是defer定义的语句，这里可能会操作返回值。
*/
package main

import "fmt"

func main() {
	fmt.Println(originalFun1())
	fmt.Println(disassembleFunc1())

	fmt.Println(originalFun2())
	fmt.Println(disassembleFunc2())

	fmt.Println(originalFun3())
	fmt.Println(disassembleFunc3())

	fmt.Println(originalFun4())
	fmt.Println(disassembleFunc4())
}

func originalFun1()(r int){
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func disassembleFunc1()(r int){
	t := 5
	r = t          // 1.赋值指令
	defer func() {
		t = t + 5  // 2.defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
	}()
	return         // 3.空的返回指令
}

func originalFun2()(r int){
	defer func() {
		r = r + 5
	}()
	return 1
}

func disassembleFunc2()(r int){
	r = 1          // 1.赋值指令
	defer func() {
		r = r + 5  // 2.defer被插入到赋值与返回之间执行，这个例子中返回值r被改变
	}()
	return         // 3.空的返回指令
}

func originalFun3()(r int){
	defer func(j int) {
		j = j + 5
	}(r)
	return 1
}

func disassembleFunc3()(r int){
	r = 1          // 1.赋值指令
	defer func(r int) {
		r = r + 5  // 2.这里改的r是之前传值传进去的r，不会改变要返回的那个r值
	}(r)
	return         // 3.空的返回指令
}


func originalFun4()(r int){
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}

func disassembleFunc4()(r int){
	r = 1             // 1.赋值指令
	defer func(r *int) {
		*r = *r + 5  // 2.这里改的r是之前传引用传进去的r，会改变要返回的那个r值
	}(&r)
	return           // 3.空的返回指令
}
