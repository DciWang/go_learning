package main

import "fmt"

func main() {


	a()
	b()
	c()
}

func a() {
	fmt.Println("func a is running")
	fmt.Println("数据库连接")
}
func b() {
	defer func() {
		/*
		注：
		recover必须搭配defer使用
		defer 必须定义在panic语句之前定义
		*/
		err := recover()     //尝试去修复，然后本方法后面的方法会继续尝试执行（很少去使用）
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("程序发生了严重的错误") //程序崩溃退出
	fmt.Println("func b is running")
}
func c() {
	fmt.Println("func c is running")
}
