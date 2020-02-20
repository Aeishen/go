/*
   @File : useIota
   @Author: Aeishen
   @Date: 2020/2/20 20:17
   @Description:枚举
*/

package main

import "fmt"

// 关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
const (
	Sunday = iota // 0
	Monday        // 1，通常省略后续⾏表达式。
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

const (
	_ = iota                    // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

// 同一常量组中，可以提供多个itoa，它们各自增长
const (
	A, B = iota, iota << 10 // 0, 0 << 10
	C, D                    // 1, 1 << 10
)

// 如果 iota ⾃增被打断，须显式恢复。
const (
	E = iota               // 0
	F                      // 1
	G = "G"                // G
	H                      // G，与上⼀⾏相同。
	I = iota               // 4，显式恢复。注意计数包含了 G、H 两⾏。
	J                      // 5
)

// 可通过⾃定义类型来实现枚举类型限制。
type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color)  {
	fmt.Println(c)
}

func main() {
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)
	fmt.Println(KB,MB,GB,TB)
	fmt.Println(A,B,C,D)
	fmt.Println(E,F,G,H,I,J)

	K := Black
	test(K)
	L := Red
	test(L)
	M := Blue
	test(M)

	//N := 1
	//test(N)    // Error: cannot use N (type int) as type Color in function argument

	test(1) // 常量会被编译器⾃动转换。
}
