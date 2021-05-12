package main

import "fmt"

//单个声明
var height int

//批量声明
var (
	userName = "DciWang" //推荐使用小驼峰命名 （声明并赋值）
	age      int
	isSingle bool
)

func main() {

	//	成员变量既然已经声明了，那必须得使用，不然编译不通过。全局变量不然
	var id string
	age = 26
	//简短变量必须
	s3 := "这个是一个人简短变量声明 "
	//匿名变量，用下划线来接受，不会占用空间，分配内存
	age, _ = variable()

	fmt.Printf("姓名：%s", userName)
	fmt.Println()
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(id)
	fmt.Println(s3)
}

func variable() (int, string) {

	return 10, "fff"
}
