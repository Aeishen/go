/*
@author : Aeishen
@data :  19/07/05, 22:34

@description :接口的实际使用（即接口作用的体现）
*/

package main

import "fmt"

// 薪资计算器
type SalaryCalculator interface {
	CalculatingWages() int // 计算总工资的行为
}

//普通员工
type OrdinaryStaff struct {
	id      int
	basePay int
}

//普通员工薪资计算
func (this OrdinaryStaff) CalculatingWages() int {
	allPay := this.basePay
	return allPay
}

func main() {
	ordinary_staff1 := OrdinaryStaff{1, 5000}
	var salaryCalculator SalaryCalculator
	salaryCalculator = ordinary_staff1

	fmt.Println("使用薪资计算器接口调用薪资计算的方法得到普通员工的薪资是：", salaryCalculator.CalculatingWages())
	fmt.Println("使用普通员工结构体调用薪资计算的方法得到普通员工的薪资是：", salaryCalculator.CalculatingWages())

	/*
		由输出可得知接口
	*/

}
