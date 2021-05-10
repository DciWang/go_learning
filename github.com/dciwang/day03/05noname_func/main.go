package main

import "fmt"

/*匿名函数*/
var f1 = func(x, y int) int {
	return x + y
}

func main() {
	f1(1, 2)

	//匿名函数一般在方法内部声明
	f2 := func(x, y int) int {
		return x + y
	}

	f2(2, 3)

	func(s1, s2 string) {
		fmt.Println(s1 + s2)
	}("wdc", "q")
}
