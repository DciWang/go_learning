package main

import "fmt"

const (

/*	iota在const关键字出现时将变量定义为0，const每新增一行常量声明iota计数一次（等同于const语句快中的索引）
	定义枚举很有用
*/
	value1 = iota
	value2
	value3
	value4
	value5
)
/*应用：定义数量级*/
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(value4)
	fmt.Println(value5)
}
