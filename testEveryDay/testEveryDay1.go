/*
   @File : testEveryDay1
   @Author: Aeishen
   @Date: 2019/12/19 21:42
   @Description:
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	//testEveryDay1_1()
	testEveryDay1_2()
}


// sync.Map 没有 Len() 方法
func testEveryDay1_1(){
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(m.Len())
}


// 输出可能不是 2000, 因为 append() 并不是并发安全的
func testEveryDay1_2(){
	var wg sync.WaitGroup
	wg.Add(2)
	var ints = make([]int, 0, 1000)
	go func() {
        for i := 0; i < 1000; i++ {
            ints = append(ints, i)
        }
        wg.Done()
    }()
    go func() {
        for i := 0; i < 1000; i++ {
            ints = append(ints, i)
        }
        wg.Done()
    }()
    wg.Wait()
    fmt.Println(len(ints))
}
