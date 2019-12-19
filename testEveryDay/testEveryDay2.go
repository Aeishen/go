/*
   @File : testEveryDay2
   @Author: Aeishen
   @Date: 2019/12/19 21:52
   @Description:
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testEveryDay2_1()
}


// struct里的name小写外部无法写入和调用改成Name就可以了
type People struct {
    name string `json:"name"`
}
func testEveryDay2_1(){
	js := `{
         "name":"11"
    }`
    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println("people: ", p)
}