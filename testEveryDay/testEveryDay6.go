/*
   @File: testEveryDay6
   @Author: Aeishen
   @Date: 2019/12/23 16:45
   @Description:
*/
package main

import "fmt"

func main() {

}

/*
	for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的
    所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
*/
func testEveryDay6_1(){
	s := []int{1,2,3,4}
	m := make(map[int]*int)

	for k,v := range s{
		m[k] = &v
	}
	for k,v := range m {
        fmt.Println(k,"->",*v)
    }
}