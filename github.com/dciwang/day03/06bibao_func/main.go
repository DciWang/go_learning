package main

import "fmt"

/*
1.闭包是什么
	闭包是一个函数，这个函数包含了他外部作用域的一个变量
*/
func main() {
	ret := f1(100)
	rett :=ret(200)
	fmt.Println("rett=",rett)




}

func f1(x int) func(y int) int {

	return func(y int) int {
		x +=y    //内部函数使用外部函数的变量
		return x
	}
}



