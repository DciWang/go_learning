package main

import "fmt"

func main() {
	deferDemon()

}


/*def 用于函数结束之前释放资源*/
func deferDemon() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")   //	被defer修饰，总是在return执行之前返回，并且所有的defer语句按栈的特征来执行
	defer fmt.Println("嘻嘻嘻")
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}