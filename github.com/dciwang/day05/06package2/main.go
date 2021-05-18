package main

import (
	"fmt"
	sum "go_learning/github.com/dciwang/day05/05packeg1"
)

/*每个包导入的时候会自动执行一个init()方法，它没有参数，也没有返回值，也不能手动调用*/
func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(sum.Sum(i, 100-i))
	}
}
