package main

import "fmt"

//变量的作用域

var x =100
func main() {

f1()
}

func f1(){
	/*函数中找变量的顺序
	1.先从函数内部查找
	2.找不到就往函数外面找，一直找到全局
	*/
	x :=10    //局部变量
	fmt.Println(x)

	for i := 0; i < 10; i++ {    //代码块变量

	}
}