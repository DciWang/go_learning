package main

import "fmt"

/**]
自定义类型以及类型别名
*/

type MyInt int  //自定义类型
type YouInt = int   //类型别名

func main() {
	var m MyInt
	var n YouInt
	m = 26
	n = 26
	fmt.Println("age=", m)
	fmt.Printf("%T\n", m)
	fmt.Println("age=", n)
	fmt.Printf("%T\n", n)


	var x rune
	x ='中'
	fmt.Println("x=", x)
	fmt.Printf("%T\n", x)     //===》int32  rune就是int32的类型别名

	var y byte =1
	fmt.Println("y=", y)
	fmt.Printf("y-type--->%T\n", y)      //------> uint8   byte 就是uint8的类型别名
}
