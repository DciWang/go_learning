package main

import "fmt"

/*
浮点型

*/
func main() {

	//不同类型的值不能相互赋值
	f1 := 1.2334   //默认Go语言中浮点型位float64
	f2 := float32(1.2334)  //转为float32

	//%T 查看类型
	fmt.Printf("f1: %T\n",f1)
	fmt.Printf("f1: %T\n",f2)


}
