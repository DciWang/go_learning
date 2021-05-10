package main

import "fmt"
//多个常量同时声明
const (
	name = "DciWang"
	age = 26
)
//批量定义常量时，如果某一行没有赋值，那么默认跟上一行的值一样。
const (
	n1 =100
	n2
	n3
)

//单个常量声明
const height = 175

func main() {
	sss := "df"
	fmt.Printf("name:%s", name)
	fmt.Println("")
	fmt.Printf("name:%s", sss)

	fmt.Println()
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
