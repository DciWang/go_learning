package main

import "fmt"

/*
can be any type of implementation (like the Generic of java)
*/
func main() {
	info := make(map[int]interface{})
	info[1] = "wangdc"
	info[2] = 24
	info[3] = true

	for i := 1; i <= len(info); i++ {
		//  assertion（断言）
		switch v := info[i].(type) {
		case string:
			fmt.Printf("x is a string，value is %v\n", v)
		case int:
			fmt.Printf("x is a int is %v\n", v)
		case bool:
			fmt.Printf("x is a bool is %v\n", v)
		default:
			fmt.Println("unsupport type！")
		}
	}
}
