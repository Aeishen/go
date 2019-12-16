/*
来源于编程拯救世界，作者江子抑
*/


/*
1.和其他语言不同，在 Go 语言中没有字符类型，字符只是整数的特殊用例。因为在 Go 中，用于表示字符的 byte 和 rune 类型都是整型的别名。
在 Go 的源码中我们可以看到：

	// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	// used, by convention, to distinguish byte values from 8-bit unsigned
	// integer values.
	type byte = uint8
	byte 是 uint8 的别名，长度为 1 个字节，用于表示 ASCII 字符

	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.
	type rune = int32
	rune 是 int32 的别名，长度为 4 个字节，用于表示以 UTF-8 编码的 Unicode 码点(Unicode 从 0 开始，为每个符号指定一个编号，这叫做「码点」（code point）,
    Unicode 和 ASCII 一样，是一种字符集，UTF-8 则是一种编码方式。)

2.在 Go 语言中使用单引号包围来表示字符，例如 'j'。参考例子1-3

3.为什么需要这两种类型？
	我们知道，byte 占用一个字节，因此它可以用于表示 ASCII 字符。而 UTF-8 是一种变长的编码方法，字符长度从 1 个字节到 4 个字节不等。byte 显然不擅长这样的表示，
	就算你想要使用多个 byte 进行表示，你也无从知晓你要处理的 UTF-8 字符究竟占了几个字节。因此，如果你在中文字符串上狂妄地进行截取，一定会输出乱码，参考例子4。
    此时就需要 rune 的帮助了。利用 []rune() 将字符串转为 Unicode 码点再进行截取，这样就无需考虑字符串中含有 UTF-8 字符的情况了，参考例子5。

4.遍历字符串，字符串遍历有两种方式，一种是下标遍历，一种是使用 range。
	下标遍历：
		由于在 Go 语言中，字符串以 UTF-8 编码方式存储，使用 len() 函数获取字符串长度时，获取到的是该 UTF-8 编码字符串的字节长度，通过下标索引字符串将会产生一个字节。
		因此，如果字符串中含有 UTF-8 编码字符，就会出现乱码，参考例子6
	range：
		range 遍历则会得到 rune 类型的字符，参考例子7

5.总结：
	Go 语言中没有字符的概念，一个字符就是一堆字节，它可能是单个字节（ASCII 字符集），也有可能是多个字节（Unicode 字符集）
	byte 是 uint8 的别名，长度为 1 个字节，用于表示 ASCII 字符
	rune 则是 int32 的别名，长度为 4 个字节，用于表示以 UTF-8 编码的 Unicode 码点
	字符串的截取是以字节为单位的
	使用下标索引字符串会产生字节
	想要遍历 rune 类型的字符则使用 range 方法进行遍历
*/


package main

import (
	"fmt"
	"reflect"
)

func main() {

	//例子1： 如果要表示 byte 类型的字符，可以使用 byte 关键字来指明字符变量的类型,
	//又因为 byte 实质上是整型 uint8，所以可以直接转成整型值。在格式化说明符中我们使用 %c 表示字符，%d 表示整型：
	fmt.Println("例子1---------------------------------")
	var byteC byte = 'j'
	fmt.Printf("字符 %c 对应的整型为 %d 类型为 %T\n", byteC, byteC, byteC)


	//例子2：与 byte 相同，想要声明 rune 类型的字符可以使用 rune 关键字指明,
	fmt.Println("例子2---------------------------------")
	var runeC rune = 'J'
	fmt.Printf("字符 %c 对应的整型为 %d 类型为 %T\n", runeC, runeC, runeC)

	//例子3：但如果在声明一个字符变量时没有指明类型，Go 会默认它是 rune 类型：
	fmt.Println("例子3---------------------------------")
	var runeD = 'J'
	fmt.Printf("字符 %c 对应的整型为 %d 类型为 %T\n", runeD, runeD, runeD)

	//例子4
	fmt.Println("例子4---------------------------------")
	testString := "你好，世界"
	fmt.Printf("字符 %c 对应的整型为 %d 类型为 %T\n", testString[0], testString[0], testString[0])
	fmt.Println(testString[:2]) // 输出乱码，因为截取了前两个字节
	fmt.Println(testString[:3]) // 输出「你」，一个中文字符由三个字节表示

	//例子5
	fmt.Println("例子5---------------------------------")
	testString = "你好，世界"
	testString_ := []rune(testString)
	fmt.Printf("字符 %c 对应的整型为 %d 类型为 %T\n", testString_[0], testString_[0], testString_[0])
	fmt.Println(string(testString_[:2])) // 输出：「你好」
	fmt.Println(string(testString_[:3])) // 输出「你好，」

	//例子6
	fmt.Println("例子6---------------------------------")
	testString = "Hello，世界"
	for i := 0; i < len(testString); i++ {
		c := testString[i]
		fmt.Printf("%c 的类型是 %s\n", c, reflect.TypeOf(c))
	}

	//例子7
	fmt.Println("例子7---------------------------------")
	testString = "Hello，世界"
	for _,c := range testString{
		fmt.Printf("%c 的类型是 %s\n", c, reflect.TypeOf(c))
	}

}
